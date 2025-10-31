package main

import (
	"os"

	"github.com/01-edu/z01"
)

func printLine(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func main() {
	// os.Args[0] is the program name, so start from 1
	for i := 1; i < len(os.Args); i++ {
		printLine(os.Args[i])
	}
}
