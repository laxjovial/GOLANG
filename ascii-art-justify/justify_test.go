package main

import "testing"

// tests for the function - ApplyAlignment
func TestApplyAlignment(t *testing.T) {
	tests := []struct {
		name   string
		text   string
		align  string
		width  int
		expect string
	}{
		{
			name:   "left alignment",
			text:   "hello",
			align:  "left",
			width:  20,
			expect: "hello",
		},
		{
			name:   "center alignment",
			text:   "hi",
			align:  "center",
			width:  10,
			expect: "    hi", // (10 - 2) / 2 = 4 spaces
		},
		{
			name:   "right alignment",
			text:   "hi",
			align:  "right",
			width:  10,
			expect: "        hi", // 10 - 2 = 8 spaces
		},
		{
			name:   "justify returns unchanged",
			text:   "hello world",
			align:  "justify",
			width:  10,
			expect: "hello world",
		},
		{
			name:   "unknown alignment defaults to text",
			text:   "hello",
			align:  "diagonal",
			width:  10,
			expect: "hello",
		},
		{
			name:   "negative padding center safe",
			text:   "longtext",
			align:  "center",
			width:  2,
			expect: "longtext",
		},
		{
			name:   "negative padding right safe",
			text:   "longtext",
			align:  "right",
			width:  2,
			expect: "longtext",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ApplyAlignment(tt.text, tt.align, tt.width)
			if got != tt.expect {
				t.Errorf("got %q, want %q", got, tt.expect)
			}
		})
	}
}

// test for the function - GetTerminal
func TestGetTerminalWidth_Fallback(t *testing.T) {

	width := GetTerminalWidth()

	if width <= 0 {
		t.Errorf("expected positive width, got %d", width)
	}

	// most systems should return 80 fallback or real terminal width
	if width < 10 {
		t.Errorf("terminal width too small, got %d", width)
	}
}
