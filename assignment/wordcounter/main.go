package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter text line by line. Type 'exit' to finish.")

	wordCount := 0
	lineCount := 0

	countLines := len(os.Args) > 1 && os.Args[1] == "-l"

	for {
		fmt.Print("> ")
		scanner.Scan()
		line := scanner.Text()

		if line == "exit" {
			break
		}

		if countLines {
			lineCount++
		} else {
			words := strings.Fields(line)
			for _, word := range words {
				if word != "exit" {
					wordCount++
				}
			}
		}

	}

	if countLines {
		fmt.Println("Number of lines entered:", lineCount)
	} else {
		fmt.Println("Number of words entered:", wordCount)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading entry:", err)
	}

}
