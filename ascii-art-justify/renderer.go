package main

import (
	"fmt"
	"strings"
)

// charWidth returns the rendered width of a single character in the banner
func charWidth(ch rune, content []string) int {
	lines := GetChar(ch, content)
	return len(lines[0])
}

// justifyRow takes a row string and distributes space between words so the
// rendered ASCII art spans the full terminal width.
func justifyRow(row string, content []string, termWidth int) string {
	words := strings.Fields(row)
	if len(words) <= 1 {
		return row // single word — nothing to distribute
	}

	// Calculate the rendered width of each word
	wordWidths := make([]int, len(words))
	for i, word := range words {
		for _, ch := range word {
			wordWidths[i] += charWidth(ch, content)
		}
	}

	// Total rendered width of all words with no spacing
	totalWordWidth := 0
	for _, w := range wordWidths {
		totalWordWidth += w
	}

	// Width of a single space character in the banner
	spaceWidth := charWidth(' ', content)

	// How many banner-space characters fit in the remaining gap between words
	gaps := len(words) - 1
	remaining := termWidth - totalWordWidth
	spacesPerGap := remaining / (gaps * spaceWidth)
	if spacesPerGap < 1 {
		spacesPerGap = 1
	}

	return strings.Join(words, strings.Repeat(" ", spacesPerGap))
}

func RenderText(input string, content []string, align string, termWidth int) {
	rows := strings.Split(input, "\\n")

	for _, row := range rows {
		if row == "" {
			fmt.Println()
			continue
		}

		// For justify, expand spaces between words before rendering
		if align == "justify" {
			row = justifyRow(row, content, termWidth)
		}

		// Build ascii lines for the row
		asciiRows := make([]string, 8)
		for i := 0; i < 8; i++ {
			line := ""
			for _, ch := range row {
				charLines := GetChar(ch, content)
				line += charLines[i]
			}
			asciiRows[i] = line
		}

		// Align each line
		for i := 0; i < 8; i++ {
			aligned := ApplyAlignment(asciiRows[i], align, termWidth)
			fmt.Println(aligned)
		}
	}
}
