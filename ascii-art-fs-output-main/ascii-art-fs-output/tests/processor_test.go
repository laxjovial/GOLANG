package tests

import (
	"os"
	"strings"
	"testing"

	"acad.learn2earn.ng/git/oelaho/ascii-art-color/master"
)

func writeTempBanner(t *testing.T) string {
    t.Helper()
    var sb strings.Builder
    for ch := 32; ch <= 65; ch++ {
        // Every character block starts with a newline separator
        sb.WriteString("\n") 
        for row := 0; row < 8; row++ {
            if ch == int('A') {
                sb.WriteString("AAAAAAAA\n")
            } else {
                sb.WriteString("        \n")
            }
        }
    }
    content := sb.String() // DO NOT TRIM. The first line should be the separator for ' '.

    f, err := os.CreateTemp(t.TempDir(), "banner*.txt")
    if err != nil {
        t.Fatalf("failed to create temp banner: %v", err)
    }
    defer f.Close()
    
    if _, err := f.WriteString(content); err != nil {
        t.Fatalf("failed to write temp banner: %v", err)
    }
    return f.Name()
}

func TestProcessor_MissingFile(t *testing.T) {
	result := master.Processor("hello", "/nonexistent/path/banner.txt", "", nil)
	if result != "" {
		t.Errorf("expected empty string for missing file, got %q", result)
	}
}

func TestProcessor_NewlineOnly(t *testing.T) {
	bannerPath := writeTempBanner(t)
	result := master.Processor("\\n", bannerPath, "", nil)
	if result != "\n" {
		t.Errorf("Processor(\\n) = %q, want newline", result)
	}
}

func TestProcessor_WithColor(t *testing.T) {
	bannerPath := writeTempBanner(t)
	ansi := "\033[31m"
	colorSet := master.AllIndexes(1)
	result := master.Processor("A", bannerPath, ansi, colorSet)

	if !strings.Contains(result, ansi) {
		t.Errorf("Processor output missing ANSI code %q", ansi)
	}
	if !strings.Contains(result, master.Reset) {
		t.Errorf("Processor output missing Reset code")
	}
}

func TestProcessor_NoColor(t *testing.T) {
	bannerPath := writeTempBanner(t)
	result := master.Processor("A", bannerPath, "", nil)

	if strings.Contains(result, "\033[") {
		t.Errorf("Processor output unexpectedly contains ANSI escape codes")
	}
	if result == "" {
		t.Error("Processor returned empty string for valid input")
	}
}

func TestProcessor_NoColorSubstring(t *testing.T) {
	bannerPath := writeTempBanner(t)
	// Index 1 is not in colorSet, so no ANSI codes should appear
	colorSet := map[int]bool{1: true}
	result := master.Processor("A", bannerPath, "\033[32m", colorSet)

	// 'A' is at index 0, which is NOT in colorSet, so output should be plain
	if strings.Contains(result, "\033[") {
		t.Errorf("Processor colored a character not in colorSet")
	}
}

func TestProcessor_ColorSubstringMatch(t *testing.T) {
	bannerPath := writeTempBanner(t)
	ansi := "\033[33m"
	// Color index 0 only — 'A' is at index 0
	colorSet := map[int]bool{0: true}
	result := master.Processor("A", bannerPath, ansi, colorSet)

	if !strings.Contains(result, ansi) {
		t.Errorf("Processor output missing ANSI code for matched index")
	}
}

func TestProcessor_OutputHasEightRows(t *testing.T) {
	bannerPath := writeTempBanner(t)
	result := master.Processor("A", bannerPath, "", nil)
	lines := strings.Split(strings.TrimRight(result, "\n"), "\n")
	if len(lines) != 8 {
		t.Errorf("expected 8 art rows, got %d", len(lines))
	}
}