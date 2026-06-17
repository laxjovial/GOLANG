package tests

import (
	"testing"

	"acad.learn2earn.ng/git/oelaho/ascii-art-color/master"
)

func TestParseColorFlag_ValidFlags(t *testing.T) {
	tests := []struct {
		flag     string
		wantCode string
		wantOk   bool
	}{
		{"--color=red", "\033[31m", true},
		{"--color=blue", "\033[34m", true},
		{"--color=green", "\033[32m", true},
		{"--color=yellow", "\033[33m", true},
		{"--color=magenta", "\033[35m", true},
		{"--color=cyan", "\033[36m", true},
		{"--color=white", "\033[37m", true},
		{"--color=black", "\033[30m", true},
	}

	for _, tt := range tests {
		t.Run(tt.flag, func(t *testing.T) {
			code, ok := master.ParseColorFlag(tt.flag)
			if ok != tt.wantOk {
				t.Errorf("ParseColorFlag(%q) ok = %v, want %v", tt.flag, ok, tt.wantOk)
			}
			if code != tt.wantCode {
				t.Errorf("ParseColorFlag(%q) code = %q, want %q", tt.flag, code, tt.wantCode)
			}
		})
	}
}

func TestParseColorFlag_InvalidFlags(t *testing.T) {
	tests := []struct {
		name string
		flag string
	}{
		{"missing prefix", "color=red"},
		{"wrong prefix", "-color=red"},
		{"empty value", "--color="},
		{"empty string", ""},
		{"unknown color", "--color=purple"},
		{"uppercase color", "--color=RED"},
		{"no equals", "--color"},
		{"space in value", "--color=re d"},
		{"random string", "hello"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			code, ok := master.ParseColorFlag(tt.flag)
			if ok {
				t.Errorf("ParseColorFlag(%q) expected ok=false, got true with code=%q", tt.flag, code)
			}
			if code != "" {
				t.Errorf("ParseColorFlag(%q) expected empty code, got %q", tt.flag, code)
			}
		})
	}
}