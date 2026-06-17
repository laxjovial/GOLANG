package master

const Reset = "\033[0m"

// ANSICode maps a color name to its ANSI escape code.
// Returns ("", false) if the color is not supported.
func ANSICode(name string) (string, bool) {
	colors := map[string]string{
		"black":   "\033[30m",
		"red":     "\033[31m",
		"green":   "\033[32m",
		"yellow":  "\033[33m",
		"blue":    "\033[34m",
		"magenta": "\033[35m",
		"cyan":    "\033[36m",
		"white":   "\033[37m",
	}
	code, ok := colors[name]
	return code, ok
}
