package main

import (
	"fmt"
	"master/master"
	"os"
)

func main() {

	// Ensure the right amount of arguments for the standard input
	if len(os.Args) != 2 {

		// Error handling for wrong amount of arguments for the standard input
		fmt.Println("Error: ensure to provide more than one argument")
		return
	}

	// Picking the input/word to do the effect on
	input := os.Args[1]

	// Convert input type to string
	strInput := string(input)

	// Application of our function
	result := master.Processor(strInput, "standard.txt")

	// Print out the result
	fmt.Println(result)
}
