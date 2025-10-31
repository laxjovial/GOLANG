package main

import (
	"os"

	"github.com/01-edu/z01"
)

func printRune(r rune) {
	z01.PrintRune(r)
}

func printString(s string) {
	for _, r := range s {
		printRune(r)
	}
	printRune('\n')
}

func printHelp() {
	printString("--insert")
	printString("  -i")
	printRune('\t')
	printRune(' ')
	printString("This flag inserts the string into the string passed as argument.")
	printString("--order")
	printString("  -o")
	printRune('\t')
	printRune(' ')
	printString("This flag will behave like a boolean, if it is called it will order the argument.")
}

func startsWith(s, prefix string) bool {
	if len(s) < len(prefix) {
		return false
	}
	for i := 0; i < len(prefix); i++ {
		if s[i] != prefix[i] {
			return false
		}
	}
	return true
}

func sliceAfter(s, prefix string) string {
	return s[len(prefix):]
}

func sortStringASCII(s string) string {
	runes := []rune(s)
	n := len(runes)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1-i; j++ {
			if runes[j] > runes[j+1] {
				runes[j], runes[j+1] = runes[j+1], runes[j]
			}
		}
	}
	return string(runes)
}

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		printHelp()
		return
	}

	insertStr := ""
	orderFlag := false
	mainStrFound := false
	mainStr := ""

	for _, arg := range args {
		if arg == "--help" || arg == "-h" {
			printHelp()
			return
		}

		if startsWith(arg, "--insert=") {
			insertStr = sliceAfter(arg, "--insert=")
			continue
		}
		if startsWith(arg, "-i=") {
			insertStr = sliceAfter(arg, "-i=")
			continue
		}

		if arg == "--order" || arg == "-o" {
			orderFlag = true
			continue
		}

		if !mainStrFound {
			mainStr = arg
			mainStrFound = true
		}
	}

	if !mainStrFound {
		printHelp()
		return
	}

	result := mainStr + insertStr

	if orderFlag {
		result = sortStringASCII(result)
	}

	printString(result)
}
