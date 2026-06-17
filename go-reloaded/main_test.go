package main

import (
	"strings"
	"testing"
)

// --- applyCommands ---

func TestApplyCommands_Cap(t *testing.T) {
	tests := []struct{ input, want string }{
		{"hello (cap) world", "Hello world"},
		{"it (cap) was a test", "It was a test"},
		{"HELLO (cap) world", "Hello world"},
	}
	for _, tc := range tests {
		words := splitWords(tc.input)
		got := joinWords(applyCommands(words))
		if got != tc.want {
			t.Errorf("input=%q\n  want=%q\n   got=%q", tc.input, tc.want, got)
		}
	}
}

func TestApplyCommands_Up(t *testing.T) {
	tests := []struct{ input, want string }{
		{"ready set go (up)", "ready set GO"},
		{"say hello (up) there", "say HELLO there"},
	}
	for _, tc := range tests {
		words := splitWords(tc.input)
		got := joinWords(applyCommands(words))
		if got != tc.want {
			t.Errorf("input=%q\n  want=%q\n   got=%q", tc.input, tc.want, got)
		}
	}
}

func TestApplyCommands_Low(t *testing.T) {
	tests := []struct{ input, want string }{
		{"I should stop SHOUTING (low)", "I should stop shouting"},
		{"HELLO (low) world", "hello world"},
	}
	for _, tc := range tests {
		words := splitWords(tc.input)
		got := joinWords(applyCommands(words))
		if got != tc.want {
			t.Errorf("input=%q\n  want=%q\n   got=%q", tc.input, tc.want, got)
		}
	}
}

func TestApplyCommands_MultiWord(t *testing.T) {
	tests := []struct{ input, want string }{
		{"it was the age of foolishness (cap, 6)", "It Was The Age Of Foolishness"},
		{"This is so exciting (up, 2)", "This is SO EXCITING"},
		{"IT WAS THE (low, 3) winter", "it was the winter"},
		{"hello (cap, 1) world", "Hello world"},
	}
	for _, tc := range tests {
		words := splitWords(tc.input)
		got := joinWords(applyCommands(words))
		if got != tc.want {
			t.Errorf("input=%q\n  want=%q\n   got=%q", tc.input, tc.want, got)
		}
	}
}

func TestApplyCommands_Hex(t *testing.T) {
	tests := []struct{ input, want string }{
		{"1E (hex) files were added", "30 files were added"},
		{"Simply add 42 (hex) done", "Simply add 66 done"},
		{"value is ff (hex)", "value is 255"},
	}
	for _, tc := range tests {
		words := splitWords(tc.input)
		got := joinWords(applyCommands(words))
		if got != tc.want {
			t.Errorf("input=%q\n  want=%q\n   got=%q", tc.input, tc.want, got)
		}
	}
}

func TestApplyCommands_Bin(t *testing.T) {
	tests := []struct{ input, want string }{
		{"It has been 10 (bin) years", "It has been 2 years"},
		{"result is 1010 (bin) ok", "result is 10 ok"},
	}
	for _, tc := range tests {
		words := splitWords(tc.input)
		got := joinWords(applyCommands(words))
		if got != tc.want {
			t.Errorf("input=%q\n  want=%q\n   got=%q", tc.input, tc.want, got)
		}
	}
}

func TestApplyCommands_HexAndBin(t *testing.T) {
	input := "Simply add 42 (hex) and 10 (bin) and you will see the result is 68."
	want := "Simply add 66 and 2 and you will see the result is 68."
	words := splitWords(input)
	got := joinWords(applyCommands(words))
	if got != want {
		t.Errorf("input=%q\n  want=%q\n   got=%q", input, want, got)
	}
}

// --- adjustPunctuation ---

func TestAdjustPunctuation_Comma(t *testing.T) {
	tests := []struct{ input, want string }{
		{"hello ,world", "hello, world"},
		{"a ,b ,c", "a, b, c"},
	}
	for _, tc := range tests {
		got := adjustPunctuation(tc.input)
		if got != tc.want {
			t.Errorf("input=%q\n  want=%q\n   got=%q", tc.input, tc.want, got)
		}
	}
}

func TestAdjustPunctuation_QuestionAndExclamation(t *testing.T) {
	tests := []struct{ input, want string }{
		{"what do you think ?", "what do you think?"},
		{"BAMM !!", "BAMM!!"},
	}
	for _, tc := range tests {
		got := adjustPunctuation(tc.input)
		if got != tc.want {
			t.Errorf("input=%q\n  want=%q\n   got=%q", tc.input, tc.want, got)
		}
	}
}

func TestAdjustPunctuation_Ellipsis(t *testing.T) {
	tests := []struct{ input, want string }{
		{"I was thinking ... You were right", "I was thinking... You were right"},
		{"Punctuation tests are ... kinda boring ,what do you think ?", "Punctuation tests are... kinda boring, what do you think?"},
	}
	for _, tc := range tests {
		got := adjustPunctuation(tc.input)
		if got != tc.want {
			t.Errorf("input=%q\n  want=%q\n   got=%q", tc.input, tc.want, got)
		}
	}
}

func TestAdjustPunctuation_ColonSemicolon(t *testing.T) {
	tests := []struct{ input, want string }{
		{"note : important", "note: important"},
		{"one ; two", "one; two"},
	}
	for _, tc := range tests {
		got := adjustPunctuation(tc.input)
		if got != tc.want {
			t.Errorf("input=%q\n  want=%q\n   got=%q", tc.input, tc.want, got)
		}
	}
}

// --- fixArticles ---

func TestFixArticles_BeforeVowel(t *testing.T) {
	tests := []struct{ input, want string }{
		{"There it was. A amazing rock!", "There it was. An amazing rock!"},
		{"bearing a untold story", "bearing an untold story"},
		{"A apple a day", "An apple a day"},
	}
	for _, tc := range tests {
		got := fixArticles(tc.input)
		if got != tc.want {
			t.Errorf("input=%q\n  want=%q\n   got=%q", tc.input, tc.want, got)
		}
	}
}

func TestFixArticles_BeforeH(t *testing.T) {
	got := fixArticles("He is a hero")
	want := "He is an hero"
	if got != want {
		t.Errorf("want=%q got=%q", want, got)
	}
}

func TestFixArticles_NoChange(t *testing.T) {
	got := fixArticles("I have a cat")
	want := "I have a cat"
	if got != want {
		t.Errorf("want=%q got=%q", want, got)
	}
}

// --- fixSingleQuotes ---

func TestFixSingleQuotes_SingleWord(t *testing.T) {
	got := fixSingleQuotes("she is ' awesome '")
	want := "she is 'awesome'"
	if got != want {
		t.Errorf("want=%q got=%q", want, got)
	}
}

func TestFixSingleQuotes_MultiWord(t *testing.T) {
	got := fixSingleQuotes("As Elton John said: ' I am the most well-known homosexual in the world '")
	want := "As Elton John said: 'I am the most well-known homosexual in the world'"
	if got != want {
		t.Errorf("want=%q got=%q", want, got)
	}
}

func TestFixSingleQuotes_AlreadyCorrect(t *testing.T) {
	got := fixSingleQuotes("it is 'fine'")
	want := "it is 'fine'"
	if got != want {
		t.Errorf("want=%q got=%q", want, got)
	}
}

// --- processText (full pipeline) ---

func TestProcessText_FullPipeline(t *testing.T) {
	input := "it (cap) was the best of times, it was the worst of times (up) , it was the age of wisdom, it was the age of foolishness (cap, 6) , it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, IT WAS THE (low, 3) winter of despair."
	want := "It was the best of times, it was the worst of TIMES, it was the age of wisdom, It Was The Age Of Foolishness, it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, it was the winter of despair."
	got := processText(input)
	if got != want {
		t.Errorf("want=%q\n got=%q", want, got)
	}
}

func TestProcessText_HexBin(t *testing.T) {
	input := "Simply add 42 (hex) and 10 (bin) and you will see the result is 68."
	want := "Simply add 66 and 2 and you will see the result is 68."
	got := processText(input)
	if got != want {
		t.Errorf("want=%q\n got=%q", want, got)
	}
}

func TestProcessText_AToAn(t *testing.T) {
	input := "There is no greater agony than bearing a untold story inside you."
	want := "There is no greater agony than bearing an untold story inside you."
	got := processText(input)
	if got != want {
		t.Errorf("want=%q\n got=%q", want, got)
	}
}

func TestProcessText_Punctuation(t *testing.T) {
	input := "Punctuation tests are ... kinda boring ,what do you think ?"
	want := "Punctuation tests are... kinda boring, what do you think?"
	got := processText(input)
	if got != want {
		t.Errorf("want=%q\n got=%q", want, got)
	}
}

// --- helpers used in tests ---

func splitWords(s string) []string {
	return strings.Fields(s)
}

func joinWords(words []string) string {
	return strings.Join(words, " ")
}