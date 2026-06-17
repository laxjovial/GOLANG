package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

var puncts = map[rune]bool{
	'.': true, ',': true, '!': true, '?': true, ':': true, ';': true,
}

var vowels = map[rune]bool{
	'a': true, 'e': true, 'i': true, 'o': true, 'u': true,
	'A': true, 'E': true, 'I': true, 'O': true, 'U': true,
	'h': true, 'H': true,
}

// --- Transformation functions ---

func toUpper(s string) string  { return strings.ToUpper(s) }
func toLower(s string) string  { return strings.ToLower(s) }
func capitalize(s string) string {
	r := []rune(strings.ToLower(s))
	if len(r) == 0 {
		return s
	}
	r[0] = unicode.ToUpper(r[0])
	return string(r)
}

func hexToDec(s string) string {
	dec, err := strconv.ParseInt(s, 16, 64)
	if err != nil {
		return s
	}
	return fmt.Sprintf("%d", dec)
}

func binToDec(s string) string {
	dec, err := strconv.ParseInt(s, 2, 64)
	if err != nil {
		return s
	}
	return fmt.Sprintf("%d", dec)
}

var transformers = map[string]func(string) string{
	"up":  toUpper,
	"low": toLower,
	"cap": capitalize,
	"hex": hexToDec,
	"bin": binToDec,
}

// --- Command parsing ---

// parseCommand parses "(cmd)" or "(cmd, N)" from a token or pair of tokens.
// Returns cmd name, count, and whether it consumed two tokens.
func parseCommand(tokens []string, i int) (cmd string, count int, twoTokens bool, ok bool) {
	token := tokens[i]

	// Single token: "(up)" or "(up,2)" or "(up, 2)"
	if strings.HasPrefix(token, "(") && strings.HasSuffix(token, ")") {
		inner := token[1 : len(token)-1]
		parts := strings.SplitN(inner, ",", 2)
		cmd = strings.TrimSpace(parts[0])
		if _, exists := transformers[cmd]; !exists {
			return "", 0, false, false
		}
		count = 1
		if len(parts) == 2 {
			n, err := strconv.Atoi(strings.TrimSpace(parts[1]))
			if err != nil {
				return "", 0, false, false
			}
			count = n
		}
		return cmd, count, false, true
	}

	// Two tokens: "(up," and "2)"
	if strings.HasPrefix(token, "(") && !strings.HasSuffix(token, ")") && i+1 < len(tokens) {
		next := tokens[i+1]
		if strings.HasSuffix(next, ")") {
			combined := token + " " + next
			inner := combined[1 : len(combined)-1]
			parts := strings.SplitN(inner, ",", 2)
			cmd = strings.TrimSpace(parts[0])
			if _, exists := transformers[cmd]; !exists {
				return "", 0, false, false
			}
			count = 1
			if len(parts) == 2 {
				n, err := strconv.Atoi(strings.TrimSpace(parts[1]))
				if err != nil {
					return "", 0, false, false
				}
				count = n
			}
			return cmd, count, true, true
		}
	}

	return "", 0, false, false
}

// --- Main processing steps ---

// applyCommands scans tokens left-to-right, applies transformations in place,
// and removes command tokens from the slice.
func applyCommands(words []string) []string {
	for i := 0; i < len(words); {
		cmd, count, twoTokens, ok := parseCommand(words, i)
		if ok {
			fn := transformers[cmd]
			// Apply to `count` words before index i
			start := i - count
			if start < 0 {
				start = 0
			}
			for j := start; j < i; j++ {
				words[j] = fn(words[j])
			}
			// Remove command token(s)
			if twoTokens {
				words = append(words[:i], words[i+2:]...)
			} else {
				words = append(words[:i], words[i+1:]...)
			}
			// Don't advance i — re-check same position
		} else {
			i++
		}
	}
	return words
}

// adjustPunctuation ensures punctuation is attached to the preceding word
// and separated from the following word by one space.
func adjustPunctuation(s string) string {
	runes := []rune(s)
	var result []rune

	for i := 0; i < len(runes); i++ {
		ch := runes[i]

		if puncts[ch] {
			// Remove trailing space before this punctuation
			for len(result) > 0 && result[len(result)-1] == ' ' {
				result = result[:len(result)-1]
			}
			result = append(result, ch)

			// Consume any spaces after this punctuation
			for i+1 < len(runes) && runes[i+1] == ' ' {
				i++
			}

			// If next char is also punctuation, don't add space (handles "..." "!?" etc.)
			if i+1 < len(runes) && !puncts[runes[i+1]] {
				result = append(result, ' ')
			}
		} else {
			result = append(result, ch)
		}
	}

	return string(result)
}

// fixArticles converts "a" -> "an" when the next word starts with a vowel or 'h'.
func fixArticles(s string) string {
	words := strings.Fields(s)
	for i := 0; i+1 < len(words); i++ {
		next := []rune(words[i+1])
		if len(next) == 0 {
			continue
		}
		if vowels[next[0]] {
			switch words[i] {
			case "a":
				words[i] = "an"
			case "A":
				words[i] = "An"
			}
		}
	}
	return strings.Join(words, " ")
}

// fixSingleQuotes removes spaces inside single-quote pairs.
func fixSingleQuotes(s string) string {
	runes := []rune(s)
	var result []rune
	inQuote := false

	for i := 0; i < len(runes); i++ {
		ch := runes[i]

		if ch == '\'' {
			if !inQuote {
				// Opening quote: skip spaces after it
				inQuote = true
				result = append(result, '\'')
				for i+1 < len(runes) && runes[i+1] == ' ' {
					i++
				}
			} else {
				// Closing quote: trim spaces before it
				for len(result) > 0 && result[len(result)-1] == ' ' {
					result = result[:len(result)-1]
				}
				result = append(result, '\'')
				inQuote = false
			}
			continue
		}

		result = append(result, ch)
	}

	return string(result)
}

func processText(input string) string {
	// Step 1: tokenize
	words := strings.Fields(input)

	// Step 2: apply (cmd) and (cmd, N) transformations
	words = applyCommands(words)

	// Step 3: rejoin
	text := strings.Join(words, " ")

	// Step 4: fix punctuation spacing
	text = adjustPunctuation(text)

	// Step 5: fix a/an
	text = fixArticles(text)

	// Step 6: fix single quotes
	text = fixSingleQuotes(text)

	return text
}

func main() {
	if len(os.Args) != 3 {
		log.Fatal("Usage: go run main.go <inputFile> <destinationFile>")
	}
	inputFile, destinationFile := os.Args[1], os.Args[2]

	fileBytes, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	finalText := processText(string(fileBytes))

	if err := os.WriteFile(destinationFile, []byte(finalText), 0644); err != nil {
		log.Fatal(err)
	}

	fmt.Println("File processed successfully!")
}