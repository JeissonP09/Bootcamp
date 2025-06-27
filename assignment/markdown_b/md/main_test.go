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

	err := run("testfile")
	if err != nil {
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
