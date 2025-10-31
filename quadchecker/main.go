package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if len(lines) == 0 {
		fmt.Println("Not a quad function")
		return
	}

	height := len(lines)
	width := len(lines[0])

	for _, line := range lines {
		if len(line) != width {
			fmt.Println("Not a quad function")
			return
		}
	}

	matches := []string{}

	if checkQuadA(lines, width, height) {
		matches = append(matches, fmt.Sprintf("[quadA] [%d] [%d]", width, height))
	}
	if checkQuadB(lines, width, height) {
		matches = append(matches, fmt.Sprintf("[quadB] [%d] [%d]", width, height))
	}
	if checkQuadC(lines, width, height) {
		matches = append(matches, fmt.Sprintf("[quadC] [%d] [%d]", width, height))
	}
	if checkQuadD(lines, width, height) {
		matches = append(matches, fmt.Sprintf("[quadD] [%d] [%d]", width, height))
	}
	if checkQuadE(lines, width, height) {
		matches = append(matches, fmt.Sprintf("[quadE] [%d] [%d]", width, height))
	}

	if len(matches) == 0 {
		fmt.Println("Not a quad function")
	} else {
		fmt.Println(strings.Join(matches, " || "))
	}
}

func checkQuadA(lines []string, w, h int) bool {
	for row := 0; row < h; row++ {
		for col := 0; col < w; col++ {
			expected := getQuadAChar(row, col, w, h)
			if lines[row][col] != expected {
				return false
			}
		}
	}
	return true
}

func getQuadAChar(row, col, w, h int) byte {
	if (row == 0 || row == h-1) && (col == 0 || col == w-1) {
		return 'o'
	}

	if row == 0 || row == h-1 {
		return '-'
	}

	if col == 0 || col == w-1 {
		return '|'
	}

	return ' '
}

func checkQuadB(lines []string, w, h int) bool {
	for row := 0; row < h; row++ {
		for col := 0; col < w; col++ {
			expected := getQuadBChar(row, col, w, h)
			if lines[row][col] != expected {
				return false
			}
		}
	}
	return true
}

func getQuadBChar(row, col, w, h int) byte {
	if row == 0 && col == 0 {
		return '/'
	}

	if row == 0 && col == w-1 {
		return '\\'
	}

	if row == h-1 && col == 0 {
		return '\\'
	}

	if row == h-1 && col == w-1 {
		return '/'
	}

	if row == 0 || row == h-1 || col == 0 || col == w-1 {
		return '*'
	}

	return ' '
}

func checkQuadC(lines []string, w, h int) bool {
	for row := 0; row < h; row++ {
		for col := 0; col < w; col++ {
			expected := getQuadCChar(row, col, w, h)
			if lines[row][col] != expected {
				return false
			}
		}
	}
	return true
}

func getQuadCChar(row, col, w, h int) byte {
	if row == 0 && (col == 0 || col == w-1) {
		return 'A'
	}

	if row == h-1 && (col == 0 || col == w-1) {
		return 'C'
	}

	if row == 0 || row == h-1 || col == 0 || col == w-1 {
		return 'B'
	}

	return ' '
}

func checkQuadD(lines []string, w, h int) bool {
	for row := 0; row < h; row++ {
		for col := 0; col < w; col++ {
			expected := getQuadDChar(row, col, w, h)
			if lines[row][col] != expected {
				return false
			}
		}
	}
	return true
}

func getQuadDChar(row, col, w, h int) byte {
	if col == 0 && (row == 0 || row == h-1) {
		return 'A'
	}

	if col == w-1 && (row == 0 || row == h-1) {
		return 'C'
	}

	if row == 0 || row == h-1 || col == 0 || col == w-1 {
		return 'B'
	}

	return ' '
}

func checkQuadE(lines []string, w, h int) bool {
	for row := 0; row < h; row++ {
		for col := 0; col < w; col++ {
			expected := getQuadEChar(row, col, w, h)
			if lines[row][col] != expected {
				return false
			}
		}
	}
	return true
}

func getQuadEChar(row, col, w, h int) byte {
	if row == 0 && col == 0 {
		return 'A'
	}

	if row == 0 && col == w-1 {
		return 'C'
	}

	if row == h-1 && col == 0 {
		return 'C'
	}

	if row == h-1 && col == w-1 {
		return 'A'
	}

	if row == 0 || row == h-1 || col == 0 || col == w-1 {
		return 'B'
	}

	return ' '
}
