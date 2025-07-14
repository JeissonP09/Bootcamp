package main

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

var (
	dir = "testdata"
	mdPath = filepath.Join(dir, "model.md")
	outBase = filepath.Join(dir, "departure")
	outPath = filepath.Join(dir, "departure.html")
	goldenPath = filepath.Join(dir, "model.html")
)

func TestRunUsingOutFlag(t *testing.T) {
	_ = os.Remove(outPath)
	
	var output bytes.Buffer
	err := run(mdPath, outBase, &output)
	if err != nil {
		t.Fatalf("error in run: %v", err)
	}

	if strings.TrimSpace(output.String()) != outPath {
		t.Errorf("output path = %q, expected %q", output.String(), outPath)
	}

	got, err := os.ReadFile(outPath)
	if err != nil {
		t.Fatalf("reading generated output: %v", err)
	}

	want, err := os.ReadFile(goldenPath)
	if err != nil {
		t.Fatalf("reading golden HTML: %v", err)
	}

	if !bytes.Equal(got, want) {
		t.Errorf("HTML mismatch\n---- GOT ----\n%s\n---- WANT ----\n%s\n", got, want)
	}
}

func TestRunWithoutOutFlag(t * testing.T) {
	var output bytes.Buffer
	err := run(mdPath, "", &output)
	if err != nil {
		t.Fatalf("error in run: %v", err)
	}

	generatedPath := strings.TrimSpace(output.String())
	got, err := os.ReadFile(generatedPath)
	if err != nil {
		t.Fatalf("reading generated file: %v", err)
	}

	want, err := os.ReadFile(goldenPath)
	if err != nil {
		t.Fatalf("reading golden HTML: %v", err)
	}
	
	if !bytes.Equal(got, want) {
		t.Errorf("HTML mismatch\n---- GOT ----\n%s\n---- WANT ----\n%s\n", got, want)
	}

	_ = os.Remove(generatedPath)
}

func TestParseContent(t *testing.T) {
	md, err := os.ReadFile(mdPath)
	if err != nil {
		t.Fatalf("reading Markdown: %v", err)
	}

	got := parseContent(md)

	want, err := os.ReadFile(goldenPath)
	if err != nil {
		t.Fatalf("reading golden HTML: %v", err)
	}

	if !bytes.Equal(got, want) {
		t.Errorf("mismatch:\n ----GOT----\n%s\n\n----WANT----\n%s", got, want)
	}
}
