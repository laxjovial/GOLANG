package master

import (
	"fmt"
	"os"
	"strings"
)

func Processor(input, filename string) string {
	result := ""

	// This reads the fonts/style file.
	inputFile, err := os.ReadFile(filename)

	// Error handling for if the file cannnot be read
	if err != nil {
		fmt.Println("Error: unable to read the .txt file")
		return ""
	}

	// Converting  each line in the fonts file from bytes to strings and saving in a variable
	lines := string(inputFile)

	// To handle line endings
	lines = strings.ReplaceAll(lines, "\r", "")

	// Split on new lines in the fonts file
	SplitLines := strings.Split(lines, "\n")

	// Treats '\n' as new line rather than actual text, splitting words on that new line
	words := strings.Split(input, "\\n")

	// Iterating through the words in our standard input
	for _, word := range words {

		// If input is empty, \n, do not try to draw an empty word,
		// print blank line and move to next word
		if word == "" {
			fmt.Println()
			continue
		}

		// Initializing the 8 sets of lines for each word
		for i := 0; i < 8; i++ {

			// Picking each character in the word to get their ascii art
			for _, char := range word {
				result += SplitLines[i+(int(char-' ')*9)+1]
			}

			// To move to the next line when printing out the result
			result += "\n"
		}
	}

	// To give the entire output
	return result
}
