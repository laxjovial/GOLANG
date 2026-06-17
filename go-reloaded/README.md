# go-reloaded

A command-line text processing tool written in Go. It reads an input file, applies a set of formatting and transformation rules, and writes the result to an output file.

## Usage

```bash
go run . <input_file> <output_file>
```

**Example:**
```bash
go run . sample.txt result.txt
```

---

## Features

### Word Transformations

Place a command **after** the word (or words) you want to transform.

| Command | Effect | Example |
|---|---|---|
| `(cap)` | Capitalizes the preceding word | `hello (cap)` → `Hello` |
| `(up)` | Uppercases the preceding word | `hello (up)` → `HELLO` |
| `(low)` | Lowercases the preceding word | `HELLO (low)` → `hello` |
| `(hex)` | Converts hex to decimal | `1E (hex)` → `30` |
| `(bin)` | Converts binary to decimal | `10 (bin)` → `2` |

To apply a transformation to **multiple preceding words**, add a number:

```
(cap, N)  (up, N)  (low, N)
```

**Examples:**
```
"This is so exciting (up, 2)"          →  "This is SO EXCITING"
"it was the age of foolishness (cap, 6)" →  "It Was The Age Of Foolishness"
"IT WAS THE (low, 3) winter"           →  "it was the winter"
```

---

### Punctuation Spacing

Punctuation marks (`. , ! ? : ;`) are always attached to the preceding word, with a single space before the next word.

```
"what do you think ?"   →  "what do you think?"
"hello ,world"          →  "hello, world"
```

Consecutive punctuation groups (like `...` or `!!`) are kept together without spaces between them:

```
"I was thinking ... You were right"  →  "I was thinking... You were right"
"BAMM !!"                            →  "BAMM!!"
```

---

### Article Correction (a → an)

`a` is automatically replaced with `an` when the next word begins with a vowel (`a e i o u`) or `h`.

```
"bearing a untold story"  →  "bearing an untold story"
"He is a hero"            →  "He is an hero"
```

---

### Single Quote Formatting

Spaces inside single-quote pairs are removed so the quotes sit directly against the enclosed text.

```
"she is ' awesome '"
→  "she is 'awesome'"

"As Elton John said: ' I am the most well-known homosexual in the world '"
→  "As Elton John said: 'I am the most well-known homosexual in the world'"
```

---

## Full Example

**Input:**
```
it (cap) was the best of times, it was the worst of times (up) , it was the age of wisdom, it was the age of foolishness (cap, 6) , it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, IT WAS THE (low, 3) winter of despair.
```

**Output:**
```
It was the best of times, it was the worst of TIMES, it was the age of wisdom, It Was The Age Of Foolishness, it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, it was the winter of despair.
```

---

## Running Tests

```bash
go test ./...
```

For verbose output with each test name:

```bash
go test -v ./...
```

Tests are in `main_test.go` and cover all transformation functions individually as well as the full pipeline.

---

## File Structure

```
.
├── main.go        # All processing logic
├── main_test.go   # Unit tests
└── README.md
```