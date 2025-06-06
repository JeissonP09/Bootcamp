package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func getInput() string {
	fmt.Println("Enter text line by line. Type 'exit' to finish.")
	fmt.Print(">")
	scanner := bufio.NewScanner(os.Stdin)
	var lines []string

	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) == "exit" {
			break
		}

		lines = append(lines, line)
	}

	return strings.Join(lines, "\n")
}

func counter(text string, countLines, countBytes bool) int {
	if text == "" {
		return 0
	}

	switch {
	case countLines:
		return len(strings.Split(text, "\n"))
	case countBytes:
		return len([]byte(text))
	default:
		return len(strings.Fields(text))
	}
}

func main() {
	countLines := flag.Bool("l", false, "count lines")
	countBytes := flag.Bool("b", false, "count bytes")
	flag.Parse()

	text := getInput()

	result := counter(text, *countLines, *countBytes)

	if *countLines {
		fmt.Println("Number of lines:", result)
	} else if *countBytes {
		fmt.Println("Number of bytes:", result)
	} else {
		fmt.Println("Number of words:", result)
	}
}
