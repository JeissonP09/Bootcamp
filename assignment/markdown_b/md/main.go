package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday/v2"
)

const header = `<!DOCTYPE html>
  <html>
    <head>
      <meta http-equiv="content-type" content="text/html; charset=utf-8" />
      <title>Markdown Preview Tool</title>
    </head>
    <body>
`

const footer = `
    </body>
  </html>
`

func main() {
	in := flag.String("in", "", "path to the markdown file")
	out := flag.String("out", "", "name output (without HTML)")
	flag.Parse()

	if err := run(*in, *out); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func saveHTML(path string, data []byte) error {

	if err := os.WriteFile(path, data, 0644); err != nil {
		return fmt.Errorf("writing file %q: %w", path, err)
	}
	return nil
}

func run(in, out string) error {
	if in == "" {
		return fmt.Errorf("the flag -in is obligatory")
	}

	mdData, err := os.ReadFile(in)
	if err != nil {
		return fmt.Errorf("reading %q: %w", in, err)
	}

	body := parseContent(mdData)

	var filename string
	if out != "" {
		filename = out + ".html"
	} else {
		base := filepath.Base(in)
		filename = base + ".html"
	}

	data := []byte(header)
	data = append(data, body...)
	data = append(data, []byte(footer)...)

	if err := saveHTML(filename, data); err != nil {
		return fmt.Errorf("could not be saved %q: %w", filename, err)
	}
	return nil
}

func parseContent(input []byte) []byte {
	unsafe := blackfriday.Run(input)
	safe := bluemonday.UGCPolicy().SanitizeBytes(unsafe)
	return safe
}
