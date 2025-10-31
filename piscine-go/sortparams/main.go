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

func compareASCII(s1, s2 string) bool {
	i := 0
	for i < len(s1) && i < len(s2) {
		if s1[i] > s2[i] {
			return true
		} else if s1[i] < s2[i] {
			return false
		}
		i++
	}
	return len(s1) > len(s2)
}

func main() {
	args := os.Args[1:]
	n := len(args)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1-i; j++ {
			if compareASCII(args[j], args[j+1]) {
				args[j], args[j+1] = args[j+1], args[j]
			}
		}
	}

	for _, arg := range args {
		printLine(arg)
	}
}
