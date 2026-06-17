package main

import "testing"

func TestParseArgs(t *testing.T) {
	tests := []struct {
		name    string
		args    []string
		wantIn  string
		wantBan string
		wantAli string
		wantErr bool
	}{
		{
			name:    "basic input only",
			args:    []string{"cmd", "hello"},
			wantIn:  "hello",
			wantBan: "standard",
			wantAli: "",
			wantErr: false,
		},
		{
			name:    "input + banner",
			args:    []string{"cmd", "hello", "shadow"},
			wantIn:  "hello",
			wantBan: "shadow",
			wantAli: "",
			wantErr: false,
		},
		{
			name:    "input + align center",
			args:    []string{"cmd", "--align=center", "hello"},
			wantIn:  "hello",
			wantBan: "standard",
			wantAli: "center",
			wantErr: false,
		},
		{
			name:    "input + align right + banner",
			args:    []string{"cmd", "--align=right", "hello", "shadow"},
			wantIn:  "hello",
			wantBan: "shadow",
			wantAli: "right",
			wantErr: false,
		},
		{
			name:    "invalid align",
			args:    []string{"cmd", "--align=diagonal", "hello"},
			wantErr: true,
		},
		{
			name:    "missing input string",
			args:    []string{"cmd"},
			wantErr: true,
		},
		{
			name:    "align but missing input",
			args:    []string{"cmd", "--align=center"},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotIn, gotBan, gotAli, err := ParseArgs(tt.args)

			if tt.wantErr {
				if err == nil {
					t.Errorf("expected error but got nil")
				}
				return
			}

			if err != nil {
				t.Errorf("unexpected error: %v", err)
				return
			}

			if gotIn != tt.wantIn {
				t.Errorf("input = %q; want %q", gotIn, tt.wantIn)
			}

			if gotBan != tt.wantBan {
				t.Errorf("banner = %q; want %q", gotBan, tt.wantBan)
			}

			if gotAli != tt.wantAli {
				t.Errorf("align = %q; want %q", gotAli, tt.wantAli)
			}
		})
	}
}
