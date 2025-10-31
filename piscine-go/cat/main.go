package main

import (
	"bufio"
	"io"
	"os"

	"github.com/01-edu/z01"
)

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		reader := bufio.NewReader(os.Stdin)
		for {
			b, err := reader.ReadByte()
			if err == io.EOF {
				break
			}
			z01.PrintRune(rune(b))
		}
		return
	}

	for _, filename := range args {
		file, err := os.Open(filename)
		if err != nil {
			printStr("ERROR: open " + filename + ": no such file or directory\n")
			os.Exit(1)
		}

		reader := bufio.NewReader(file)
		for {
			b, err := reader.ReadByte()
			if err == io.EOF {
				break
			}
			z01.PrintRune(rune(b))
		}

		file.Close()
	}
}
