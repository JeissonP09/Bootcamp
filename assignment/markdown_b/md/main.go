package main

import (
	"flag"
	"fmt"
	"os"
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
	out := flag.String("out", "", "name output (without HTML)")
	flag.Parse()

	if err := run(*out); err != nil {
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

func run(out string) error {
	if out == "" {
		return fmt.Errorf("flag -out is mandatory")
	}

	filename := out + ".html"
	data := []byte(header + footer)

	if err := saveHTML(filename, data); err != nil {
		return fmt.Errorf("could not be saved HTML: %w", err)
	}
	return nil
}
