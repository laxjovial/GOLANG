package main

import (
	"os"
	"strings"
)

// Reads banner files like standard.txt
// Splits on new line then saves as slice of string
func LoadBanner(filename string) ([]string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return strings.Split(string(data), "\n"), nil
}

// Uses characters ascii code to get their ascii art from a banner file
func GetChar(ch rune, lines []string) []string {
	start := (int(ch) - 32) * 9
	return lines[start+1 : start+9]
}

// Checks for valid banner files - standard, shadow, thinkertoy
func IsValidBanner(banner string) bool {
	return banner == "standard" || banner == "shadow" || banner == "thinkertoy"
}
