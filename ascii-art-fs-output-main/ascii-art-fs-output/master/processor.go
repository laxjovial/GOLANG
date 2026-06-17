package master

import (
	"fmt"
	"os"
	"strings"
)

// Processor reads a banner file and returns the full ASCII art for input.
// If ansiCode is non-empty, characters whose index appears in colorSet are
// wrapped in that ANSI code + Reset. Pass nil / "" for plain (no color) output.
func Processor(input, filename, ansiCode string, colorSet map[int]bool) string {
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error: unable to read the .txt file")
		return ""
	}

	splitLines := strings.Split(
		strings.ReplaceAll(string(data), "\r", ""),
		"\n",
	)

	// Split by literal \n
    parts := strings.Split(input, "\\n")
    
    // Check if the input is ONLY newlines
    onlyNewlines := true
    for _, p := range parts {
        if p != "" {
            onlyNewlines = false
            break
        }
    }

    if onlyNewlines && len(parts) > 1 {
        return strings.Repeat("\n", len(parts)-1)
    }

	var sb strings.Builder
	for _, word := range strings.Split(input, "\\n") {
		if word == "" {
			sb.WriteByte('\n')
			continue
		}
		sb.WriteString(renderWord(word, splitLines, ansiCode, colorSet))
	}
	return sb.String()
}

// renderWord builds the 8-row ASCII art block for a single word,
// applying color to the characters whose index is in colorSet.
func renderWord(word string, splitLines []string, ansiCode string, colorSet map[int]bool) string {
	runes := []rune(word)
	var sb strings.Builder

	for row := 0; row < 8; row++ {
		for idx, char := range runes {
			artLine := splitLines[row+(int(char-' ')*9)+1]
			sb.WriteString(colorize(artLine, idx, ansiCode, colorSet))
		}
		sb.WriteByte('\n')
	}
	return sb.String()
}

// colorize wraps a single art segment in ANSI codes when the character index
// is present in colorSet, otherwise returns the segment unchanged.
func colorize(segment string, charIdx int, ansiCode string, colorSet map[int]bool) string {
	if ansiCode == "" || !colorSet[charIdx] {
		return segment
	}
	return ansiCode + segment + Reset
}
