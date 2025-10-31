package main

import (
	"os"

	"github.com/01-edu/z01"
)

func isVowel(r rune) bool {
	switch r {
	case 'a', 'e', 'i', 'o', 'u',
		'A', 'E', 'I', 'O', 'U':
		return true
	}
	return false
}

func printRune(r rune) {
	z01.PrintRune(r)
}

type vowelPos struct {
	argIndex  int
	runeIndex int
}

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		printRune('\n')
		return
	}

	vowelPositions := []vowelPos{}
	vowels := []rune{}

	for ai, arg := range args {
		for ri, r := range arg {
			if isVowel(r) {
				vowelPositions = append(vowelPositions, vowelPos{ai, ri})
				vowels = append(vowels, r)
			}
		}
	}

	if len(vowels) == 0 {
		for i, arg := range args {
			for _, r := range arg {
				printRune(r)
			}
			if i != len(args)-1 {
				printRune(' ')
			}
		}
		printRune('\n')
		return
	}

	n := len(vowels)
	for i := 0; i < n/2; i++ {
		vowels[i], vowels[n-1-i] = vowels[n-1-i], vowels[i]
	}

	argsRunes := make([][]rune, len(args))
	for i, arg := range args {
		argsRunes[i] = []rune(arg)
	}

	for i, pos := range vowelPositions {
		argsRunes[pos.argIndex][pos.runeIndex] = vowels[i]
	}

	for i, argRunes := range argsRunes {
		for _, r := range argRunes {
			printRune(r)
		}
		if i != len(argsRunes)-1 {
			printRune(' ')
		}
	}
	printRune('\n')
}
