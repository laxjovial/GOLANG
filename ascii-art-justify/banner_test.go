package main

import (
	"fmt"
	"os"
	"reflect"
	"testing"
)

//tests for the function - LoadBanner

// helper to create temp file
func createTempFile(t *testing.T, content string) string {
	t.Helper()

	tmpFile, err := os.CreateTemp("", "banner_test_*.txt")
	if err != nil {
		t.Fatalf("failed to create temp file: %v", err)
	}

	if _, err := tmpFile.WriteString(content); err != nil {
		t.Fatalf("failed to write temp file: %v", err)
	}

	tmpFile.Close()
	return tmpFile.Name()
}

func TestLoadBanner_SuccessMultipleLines(t *testing.T) {
	filename := createTempFile(t, "line1\nline2\nline3")
	defer os.Remove(filename)

	result, err := LoadBanner(filename)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := []string{"line1", "line2", "line3"}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func TestLoadBanner_SingleLine(t *testing.T) {
	filename := createTempFile(t, "onlyline")
	defer os.Remove(filename)

	result, err := LoadBanner(filename)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := []string{"onlyline"}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func TestLoadBanner_EmptyFile(t *testing.T) {
	filename := createTempFile(t, "")
	defer os.Remove(filename)

	result, err := LoadBanner(filename)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// IMPORTANT: strings.Split("", "\n") returns []string{""}
	expected := []string{""}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func TestLoadBanner_TrailingNewline(t *testing.T) {
	filename := createTempFile(t, "line1\nline2\n")
	defer os.Remove(filename)

	result, err := LoadBanner(filename)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// trailing newline produces empty string at the end
	expected := []string{"line1", "line2", ""}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func TestLoadBanner_FileDoesNotExist(t *testing.T) {
	_, err := LoadBanner("non_existent_file.txt")

	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

func TestLoadBanner_PermissionDenied(t *testing.T) {
	tmpFile, err := os.CreateTemp("", "banner_test_perm_*.txt")
	if err != nil {
		t.Fatalf("failed to create temp file: %v", err)
	}
	filename := tmpFile.Name()
	tmpFile.Close()

	// remove read permission
	err = os.Chmod(filename, 0000)
	if err != nil {
		t.Fatalf("failed to chmod: %v", err)
	}
	defer os.Remove(filename)

	_, err = LoadBanner(filename)
	if err == nil {
		t.Fatal("expected permission error, got nil")
	}
}

//tests for the function - GetChar

func TestGetChar_Valid(t *testing.T) {
	lines := make([]string, 400) // enough space

	for i := range lines {
		lines[i] = fmt.Sprintf("line_%d", i)
	}

	ch := 'A'
	start := (int(ch) - 32) * 9

	expected := lines[start+1 : start+9]
	result := GetChar(ch, lines)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func TestGetChar_SpaceChar(t *testing.T) {
	lines := make([]string, 20)
	for i := range lines {
		lines[i] = string(rune(i))
	}

	ch := ' ' // ASCII 32 → start = 0
	expected := lines[1:9]

	result := GetChar(ch, lines)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func TestGetChar_OutOfBounds(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("expected panic but did not get one")
		}
	}()

	lines := make([]string, 10) // too small
	ch := 'A'

	GetChar(ch, lines) // should panic due to slice bounds
}

// tests for the function - IsValidBanner

func TestIsValidBanner(t *testing.T) {
	tests := []struct {
		name   string
		banner string
		want   bool
	}{
		{"valid standard", "standard", true},
		{"valid shadow", "shadow", true},
		{"valid thinkertoy", "thinkertoy", true},

		{"invalid empty", "", false},
		{"invalid random", "random", false},
		{"invalid uppercase", "STANDARD", false},
		{"invalid typo", "standerd", false},
		{"invalid extra space", " standard", false},
		{"invalid suffix", "shadow1", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsValidBanner(tt.banner)
			if got != tt.want {
				t.Errorf("IsValidBanner(%q) = %v; want %v", tt.banner, got, tt.want)
			}
		})
	}
}
