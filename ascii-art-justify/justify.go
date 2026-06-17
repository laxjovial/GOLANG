package main

import (
	"os"
	"os/exec"
	"strconv"
	"strings"
)

// Text on terminal alignment; left, center, right
func ApplyAlignment(text string, align string, width int) string {
	switch align {
	case "left": // align to left
		return text
	case "center": // align to center
		padding := (width - len(text)) / 2
		if padding < 0 {
			padding = 0
		}
		return strings.Repeat(" ", padding) + text
	case "right": // align to right
		padding := width - len(text)
		if padding < 0 {
			padding = 0
		}
		return strings.Repeat(" ", padding) + text
	case "justify":
		// Word spacing is already expanded by justifyRow() in renderer.go
		// before the ASCII art is rendered, so no post-render padding needed.
		return text
	default:
		return text
	}
}

// GetTerminalWidth returns the width of the current terminal in columns.
// Falls back to 80 if the terminal cannot be accessed or detected.
func GetTerminalWidth() int {

	// Open the controlling terminal device
	tty, err := os.Open("/dev/tty")
	if err != nil {
		return 80 // Default width if no TTY is available
	}
	defer tty.Close() // Ensure the terminal is closed after use
// Ensure the terminal is closed after use

	// Execute "tput cols" to get the terminal width (number of columns)
	cmd := exec.Command("tput", "cols")
	cmd.Stdin = tty //  Attach terminal input so tput can detect the environment
	
	// Run the command and capture its output
	output, err := cmd.Output()
	if err != nil {
		return 80
	}

	// Convert output from bytes to string and remove extra whitespace
	widthStr := strings.TrimSpace(string(output))
	// Convert the string value to an integer
	width, err := strconv.Atoi(widthStr)
	if err != nil {
		return 80 // fallback
	}

	return width
}
