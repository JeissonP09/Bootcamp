package main

import (
	"testing"
)

func TestWordCounter(t *testing.T) {
	cases := []struct {
		name     string
		text     string
		expected int
	}{
		{"1. When the user types a single sentence with some words.", "Soy el primer test", 4},
		{"2. When the user types multiple sentences with some words per sentence.", "first line\nsecond line", 4},
		{"3. When the user types a single word.", "word", 1},
		{"4. When the user types a single composed word. Example: read-only.", "read-only", 1},
		{"5. When the user types multiple break lines (Example: \n\n).", "example \n\n word", 2},
		{"6. When the user types Exit, exit, or EXIT", "Exit\n exit\n EXIT", 3},
	}

	for _, c := range cases {
		result := counter(c.text, false)
		if result != c.expected {
			t.Errorf("Fail: %s: Expected [%d], but result obtained: [%d]", c.name, c.expected, result)
		}
	}
}

func TestLineCounter(t *testing.T) {
	cases := []struct {
		name     string
		text     string
		expected int
	}{
		{"1. When the user types a single line.", "one line", 1},
		{
			name: "2. When the user types multiple lines without a break line in between lines.",
			text: `one line
			second line
			three line`,
			expected: 3,
		},
		{"3. When the user types multiple lines with one or more break lines in between lines.", "one line\n\n test\n\n line", 5},
		{"Exit at the beginning of the line", "Exit line of test", 1},
		{"exit in the middle of the line", "line exit test", 1},
		{"EXIT at the end of the line", "line test EXIT", 1},
	}

	for _, c := range cases {
		result := counter(c.text, true)
		if result != c.expected {
			t.Errorf("Fail: %s: Expected [%d], but result obtained: [%d]", c.name, c.expected, result)
		}
	}
}
