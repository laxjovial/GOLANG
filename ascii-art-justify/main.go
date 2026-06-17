package main

import (
	"fmt"
	"os"
)

func main() {
	// Parse command-line arguments into input text, banner style, alignment, and error (if any)
	input, banner, align, err := ParseArgs(os.Args)
	if err != nil {
		fmt.Println(err) 
		return 
	}
	
	// Exit early if no input string is provided
	if input == "" {
		return
	}

	// Handle explicit newline input ("\\n") by printing a blank line
	if input == "\\n" {
		fmt.Println()
		return
	}

	// Validate that the provided banner name is supported
	if !IsValidBanner(banner) {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]")
		fmt.Println("Example: go run . --align=right something standard")
		return
	}

	// Load the banner file (e.g., "standard.txt") into memory
	content, err := LoadBanner(banner + ".txt")
	if err != nil {
		fmt.Println("error:", err) // Display file loading error
		return     
	}

	// Get the current terminal width and render the ASCII art with alignment
	termWidth := GetTerminalWidth()
	RenderText(input, content, align, termWidth)
}
