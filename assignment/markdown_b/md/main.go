package main

import (
	"flag"
	"fmt"
	"io"
	"os"

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

	if err := run(*in, *out, os.Stdout); err != nil {
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

func run(in, out string, writer io.Writer) error {
	if in == "" {
		return fmt.Errorf("the flag -in is obligatory")
	}

	mdData, err := os.ReadFile(in)
	if err != nil {
		return fmt.Errorf("reading %q: %w", in, err)
	}

	html := parseContent(mdData)

	var filename string
	if out == "" {
		tmpFile, err := os.CreateTemp(".", "md*.html")
		if err != nil {
			return fmt.Errorf("created temp file: %w", err)
		}
		tmpFile.Close()
		filename = tmpFile.Name()
	} else {
		filename = out + ".html"
	}

	if err := saveHTML(filename, html); err != nil {
		return fmt.Errorf("could not be saved %q: %w", filename, err)
	}

	fmt.Fprintln(writer, filename)

	return nil
}

func parseContent(input []byte) []byte {
	unsafe := blackfriday.Run(input)
	safe := bluemonday.UGCPolicy().SanitizeBytes(unsafe)

	var html []byte
	html = append(html, []byte(header)...)
	html = append(html, safe...)
	html = append(html, []byte(footer)...)

	return html
}
