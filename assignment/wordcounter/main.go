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

func wordCounter(text string) int {
	return len(strings.Fields(text))
}

func lineCounter(text string) int {
	if text == "" {
		return 0
	}
	return len(strings.Split(text, "\n"))
}

func main() {
	mod := "words"

	if len(os.Args) > 1 && os.Args[1] == "-l" {
		mod = "lines"
	}

	text := getInput()

	switch mod {
	case "lines":
		fmt.Println("Number of lines: ", lineCounter(text))
	default:
		fmt.Println("Number of words: ", wordCounter(text))
	}
}
