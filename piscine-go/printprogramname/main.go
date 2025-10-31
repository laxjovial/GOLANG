package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	name := os.Args[0]
	lastSlash := -1
	for i, j := range name {
		if j == '/' {
			lastSlash = i
		}
	}
	if lastSlash != -1 {
		name = name[lastSlash+1:]
	}
	for _, j := range name {
		z01.PrintRune(j)
	}
	z01.PrintRune('\n')
}
