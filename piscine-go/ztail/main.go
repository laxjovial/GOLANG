package main

import (
	"fmt"
	"os"
)

func atoi(s string) int {
	n := 0
	for _, r := range s {
		n = n*10 + int(r-'0')
	}
	return n
}

func main() {
	args := os.Args
	if len(args) < 4 || args[1] != "-c" {
		fmt.Printf("Usage: go run . -c <number> <files...>\n")
		return
	}

	byteCount := atoi(args[2])
	files := args[3:]
	exitCode := 0

	for i, fileName := range files {
		data, err := os.ReadFile(fileName)
		if err != nil {
			fmt.Printf("open %s: no such file or directory\n", fileName)
			exitCode = 1
			continue
		}

		if len(files) > 1 {
			if i > 0 {
				fmt.Printf("\n")
			}
			fmt.Printf("==> %s <==\n", fileName)
		}

		if byteCount > len(data) {
			fmt.Printf("%s", string(data))
		} else {
			fmt.Printf("%s", string(data[len(data)-byteCount:]))
		}
	}

	if exitCode != 0 {
		os.Exit(1)
	}
}
