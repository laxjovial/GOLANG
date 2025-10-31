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
	// os.Args[0] is the program name, so start from len(os.Args)-1 down to 1
	for i := len(os.Args) - 1; i > 0; i-- {
		printLine(os.Args[i])
	}
}
