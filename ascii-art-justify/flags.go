package main

import (
	"errors"
	"strings"
)

func ParseArgs(args []string) (input string, banner string, align string, err error) {
	// Default positions for string and banner arguments
	stringIndex := 1
	bannerIndex := 2

	// Default alignment (empty means no alignment applied)
	align = ""

	// Check if an alignment option is provided as the second argument
	if len(args) > 1 && strings.HasPrefix(args[1], "--align=") {
		// Split the alignment argument into key and value
		// "--align=center" ["--align", "center"]
		parts := strings.SplitN(args[1], "=", 2)

		// Validate that the split resulted in exactly 2 parts
		if len(parts) != 2 {
			err = errors.New("Usage: go run . [OPTION] [STRING] [BANNER]\nExample: go run . --align=right something standard")
			return
		}

		// Extract the alignment value
		align = parts[1]

		// Ensure the alignment value is one of the allowed options
		if align != "left" && align != "right" && align != "center" && align != "justify" {
			err = errors.New("Usage: go run . [OPTION] [STRING] [BANNER]\nExample: go run . --align=right something standard")
			return
		}

		// Adjust indexes since alignment option takes position args[1]
		stringIndex = 2
		bannerIndex = 3
	}

	// Check if the input string exists
	if len(args) <= stringIndex {
		err = errors.New("Missing input string")
		return
	}

	// Assign the input string
	input = args[stringIndex]

	// Default banner style
	banner = "standard"

	// If a banner argument is provided, override the default
	if len(args) > bannerIndex {
		banner = args[bannerIndex]
	}

	return
}
