package tests

import (
	"testing"

	"acad.learn2earn.ng/git/oelaho/ascii-art-color/master"
)

func TestANSICode_KnownColors(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		wantCode string
		wantOk   bool
	}{
		{"black", "black", "\033[30m", true},
		{"red", "red", "\033[31m", true},
		{"green", "green", "\033[32m", true},
		{"yellow", "yellow", "\033[33m", true},
		{"blue", "blue", "\033[34m", true},
		{"magenta", "magenta", "\033[35m", true},
		{"cyan", "cyan", "\033[36m", true},
		{"white", "white", "\033[37m", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			code, ok := master.ANSICode(tt.input)
			if ok != tt.wantOk {
				t.Errorf("ANSICode(%q) ok = %v, want %v", tt.input, ok, tt.wantOk)
			}
			if code != tt.wantCode {
				t.Errorf("ANSICode(%q) code = %q, want %q", tt.input, code, tt.wantCode)
			}
		})
	}
}

func TestANSICode_UnknownColors(t *testing.T) {
	unknowns := []string{"", "purple", "orange", "BLUE", "Blue", "RED", " red", "red "}
	for _, input := range unknowns {
		t.Run(input, func(t *testing.T) {
			code, ok := master.ANSICode(input)
			if ok {
				t.Errorf("ANSICode(%q) expected ok=false, got true with code=%q", input, code)
			}
			if code != "" {
				t.Errorf("ANSICode(%q) expected empty code, got %q", input, code)
			}
		})
	}
}

func TestResetConstant(t *testing.T) {
	if master.Reset != "\033[0m" {
		t.Errorf("Reset = %q, want %q", master.Reset, "\033[0m")
	}
}