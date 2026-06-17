package master

import "strings"

// ParseColorFlag validates a --color=<value> flag and returns the resolved
// ANSI escape code. Returns ("", false) if the flag is malformed or the
// color value is unrecognised.
func ParseColorFlag(flag string) (string, bool) {
	if !strings.HasPrefix(flag, "--color=") {
		return "", false
	}
	parts := strings.SplitN(flag, "=", 2)
	if len(parts) != 2 || parts[1] == "" {
		return "", false
	}
	return ANSICode(parts[1])
}

// ParseOutputFlag validates a --output=<fileName.txt> flag and returns the
// file name. Returns ("", false) if the flag is malformed or the file name
// is empty.
func ParseOutputFlag(flag string) (string, bool) {
	if !strings.HasPrefix(flag, "--output=") {
		return "", false
	}
	parts := strings.SplitN(flag, "=", 2)
	if len(parts) != 2 || parts[1] == "" {
		return "", false
	}
	return parts[1], true
}