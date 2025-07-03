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
	mdPath := filepath.Join("testdata", "model.md")
	htmlPath := filepath.Join("testdata", "model.html")

	md, err := os.ReadFile(mdPath)
	if err != nil {
		t.Fatalf("reading Markdown: %v", err)
	}

	got := parseContent(md)

	want, err := os.ReadFile(htmlPath)
	if err != nil {
		t.Fatalf("reading golden HTML: %v", err)
	}

	if !bytes.Equal(got, want) {
		t.Errorf("mismatch:\n got=\n%s\n\nwant=\n%s", got, want)
	}
}
