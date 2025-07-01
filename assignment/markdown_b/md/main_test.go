package main

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"
)

func TestRun(t *testing.T) {
	tmp := t.TempDir()

	oldWd, _ := os.Getwd()
	t.Cleanup(func() { os.Chdir(oldWd) })

	os.Chdir(tmp)

	err := os.WriteFile("testfile.md", []byte("#Test\n"), 0644)
	if err != nil {
		t.Fatal(err)
	}

	if err := run("testfile.md", "testfile"); err != nil {
		t.Fatalf("run return error: %v", err)
	}

	path := filepath.Join(tmp, "testfile.html")
	content, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.HasPrefix(content, []byte(header)) {
		t.Error("lack header")
	}
	if !bytes.HasSuffix(content, []byte(footer)) {
		t.Error("lack footer")
	}
}

func TestParseContent(t *testing.T) {
	md := []byte(`# Test Markdown File

Just a test

## Bullets

- Links [Link1](https://example.com)

## Quotes

> Quotes in **bold** and _italic_ text
`)

	got := parseContent(md)
	if len(got) == 0 {
		t.Fatal("parseContent return empty slice")
	}

	checks := []struct{ name, want string }{
		{"h1", `<h1>Test Markdown File</h1>`},
		{"paragraph", `<p>Just a test</p>`},
		{"h2", `<h2>Bullets</h2>`},
		{"ul", `<ul>`},
		{"li", `<li>Links <a href="https://example.com" rel="nofollow">Link1</a></li>`},
		{"blockquote", `<blockquote>`},
		{"strong", `<strong>bold</strong>`},
		{"em", `<em>italic</em>`},
	}

	for _, c := range checks {
		if !bytes.Contains(got, []byte(c.want)) {
			t.Errorf("fragment %q no found in output:\nexpected to include: %s", c.name, c.want)
		}
	}
}
