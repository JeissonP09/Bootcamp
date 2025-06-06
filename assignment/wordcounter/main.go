package main

import (
	"bufio"
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

func counter(text string, lines bool) int {
	if text == "" {
		return 0
	}

	if lines {
		return len(strings.Split(text, "\n"))
	}
	return len(strings.Fields(text))
}

func main() {
	text := getInput()
	countLines := false

	if len(os.Args) > 1 && os.Args[1] == "-l" {
		countLines = true
	}

	result := counter(text, countLines)

	if countLines {
		fmt.Println("Number of lines: ", result)
	} else {
		fmt.Println("Number of words: ", result)
	}
}
