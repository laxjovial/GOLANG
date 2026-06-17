package main

import (
	"fmt"
	"os"
	"strings"

	"acad.learn2earn.ng/git/oelaho/ascii-art-color/master"
)

const defaultBanner = "standard.txt"

const usageMsg = `Usage: go run . [OPTION] [STRING] [BANNER]

EX: go run . --output=<fileName.txt> something standard`

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		exitUsage()
	}

	cfg, ok := parseArgs(args)
	if !ok {
		exitUsage()
	}

	// Build color set
	var colorSet map[int]bool
	if cfg.color != "" {
		if cfg.substring != "" {
			colorSet = master.SubstringIndexes(cfg.text, cfg.substring)
		} else {
			colorSet = master.AllIndexes(len([]rune(cfg.text)))
		}
	}

	result := master.Processor(cfg.text, cfg.banner, cfg.color, colorSet)

	if cfg.outputFile != "" {
		err := os.WriteFile(cfg.outputFile, []byte(result), 0644)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error: could not write to file:", cfg.outputFile)
			os.Exit(1)
		}
	} else {
		fmt.Print(result)
	}
}

// config holds all parsed argument values.
type config struct {
	color      string // ANSI code, empty if not set
	outputFile string // file name, empty if not set
	substring  string // substring to color, empty if not set
	text       string // the main input string (required)
	banner     string // banner file path
}

// parseArgs separates flags from positional args and builds a config.
// Returns (config, false) on any invalid input.
func parseArgs(args []string) (config, bool) {
	cfg := config{banner: defaultBanner}

	var positional []string

	for _, arg := range args {
		switch {
		case strings.HasPrefix(arg, "--color="):
			ansi, ok := master.ParseColorFlag(arg)
			if !ok {
				return cfg, false
			}
			cfg.color = ansi

		case strings.HasPrefix(arg, "--output="):
			filename, ok := master.ParseOutputFlag(arg)
			if !ok {
				return cfg, false
			}
			cfg.outputFile = filename

		case strings.HasPrefix(arg, "--"):
			// Unknown flag
			return cfg, false

		default:
			positional = append(positional, arg)
		}
	}

	// Positional args: [text] or [text, banner] or [substring, text] or [substring, text, banner]
	// Rule: text is required. Banner is always last if it matches a known banner name.
	// Substring only makes sense when --color is set.
	switch len(positional) {
	case 1:
		cfg.text = positional[0]

	case 2:
		if isBannerName(positional[1]) {
			cfg.text = positional[0]
			cfg.banner = positional[1] + ".txt"
		} else if cfg.color != "" {
			// --color=x sub "text"  (no banner)
			cfg.substring = positional[0]
			cfg.text = positional[1]
		} else {
			return cfg, false
		}

	case 3:
		// substring text banner  (requires --color)
		if cfg.color == "" {
			return cfg, false
		}
		if !isBannerName(positional[2]) {
			return cfg, false
		}
		cfg.substring = positional[0]
		cfg.text = positional[1]
		cfg.banner = positional[2] + ".txt"

	default:
		return cfg, false
	}

	if cfg.text == "" {
		return cfg, false
	}

	return cfg, true
}

// isBannerName returns true for the three known banner font names.
func isBannerName(s string) bool {
	switch s {
	case "standard", "shadow", "thinkertoy":
		return true
	}
	return false
}

func exitUsage() {
	fmt.Println(usageMsg)
	os.Exit(1)
}