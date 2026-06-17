package main

import (
	"bytes"
	"io"
	"os"

	//"reflect"
	"strings"
	"testing"
)

// Helper to create mock banner content
// Each char will be 4 columns wide for testing simplicity
func getMockContent() []string {
	content := make([]string, 1000)
	for i := range content {
		// Create a predictable pattern: "####" for each char line
		content[i] = "####"
	}
	return content
}

func TestCharWidth(t *testing.T) {
	content := getMockContent()
	// Since our mock content lines are "####", width should be 4
	got := charWidth('A', content)
	want := 4

	if got != want {
		t.Errorf("charWidth() = %d; want %d", got, want)
	}
}

func TestJustifyRow(t *testing.T) {
	content := getMockContent() // spaceWidth is 4
	termWidth := 20             // Total width to fill

	// Test Case: Two words
	// "Hi Hi" -> Words are "Hi" (8px) and "Hi" (8px) = 16px total.
	// Remaining = 20 - 16 = 4px.
	// spaceWidth is 4. gaps = 1.
	// spacesPerGap = 4 / (1 * 4) = 1.
	row := "Hi Hi"
	got := justifyRow(row, content, termWidth)
	want := "Hi Hi" // Result should have 1 space between

	if got != want {
		t.Errorf("justifyRow() = %q; want %q", got, want)
	}

	// Test Case: Extra Space
	// "A B" -> A(4) + B(4) = 8px. Remaining = 20 - 8 = 12px.
	// gaps = 1. spacesPerGap = 12 / (1 * 4) = 3 spaces.
	row2 := "A B"
	got2 := justifyRow(row2, content, termWidth)
	want2 := "A   B" // 3 spaces between

	if got2 != want2 {
		t.Errorf("justifyRow() = %q; want %q", got2, want2)
	}
}

func TestRenderText_Basic(t *testing.T) {
	// Note: This requires ApplyAlignment to be defined in your main package.
	// We capture Stdout to verify printing.
	content := getMockContent()

	// Mock implementation of ApplyAlignment if not already in your test scope
	// (If it's in another file in the same package, you don't need this)

	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	RenderText("A", content, "left", 10)

	w.Close()
	var buf bytes.Buffer
	io.Copy(&buf, r)
	os.Stdout = old

	lines := strings.Split(strings.TrimSpace(buf.String()), "\n")
	if len(lines) != 8 {
		t.Errorf("RenderText produced %d lines; want 8", len(lines))
	}
}

func TestJustifyRow_SingleWord(t *testing.T) {
	content := getMockContent()
	row := "Hello"
	got := justifyRow(row, content, 100)
	if got != "Hello" {
		t.Errorf("justifyRow with single word should return word as-is, got %q", got)
	}
}
