package main

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestRunUsingOutFlag(t *testing.T) {
	tmp := t.TempDir()

	oldWd, _ := os.Getwd()
	t.Cleanup(func() { os.Chdir(oldWd) })
	os.Chdir(tmp)
	
	err := os.WriteFile("testfile.md", []byte("# Test\n"), 0644)
	if err != nil {
		t.Fatal(err)
	}
	
	var output bytes.Buffer
	err = run("testfile.md", "departure", &output)
	if err != nil {
		t.Fatalf("error in run: %v", err)
	}

	filename := filepath.Join(tmp, "departure.html")
	content, err :=os.ReadFile(filename)
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

func TestRunWithoutOutFlag(t * testing.T) {
	tmp := t.TempDir()

	oldWd, _ := os.Getwd()
	t.Cleanup(func() { os.Chdir(oldWd) })
	os.Chdir(tmp)

	err := os.WriteFile("temp.md", []byte("# Test\n"), 0644)
	if err != nil {
		t.Fatal(err)
	}

	var output bytes.Buffer
	err = run("temp.md", "", &output)
	if err != nil {
		t.Fatalf("error in run: %v", err)
	}

	filename := strings.TrimSpace(output.String())
	content, err := os.ReadFile(filename)
	if err != nil {
		t.Fatalf("The generated file could not be read: %v", err)
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
