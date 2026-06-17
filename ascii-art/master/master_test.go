package master

import (
	"testing"
)

func TestProcessor(t *testing.T) {
	input := "k"
	expected := "       \n _     \n| | _  \n| |/ / \n|   <  \n|_|\\_\\ \n       \n       \n"

	
	result := Processor(input, "standard.txt")
	if result != expected {
		t.Errorf("Processor(%q) = %q; want %q", input, result, expected)
	}
}
