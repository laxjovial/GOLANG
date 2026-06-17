# ASCII Art Color

A command-line tool that renders text as ASCII art with optional ANSI color support.

---

## Usage

```bash
go run . [OPTION] [STRING]
```

| Arguments | Description |
|---|---|
| `go run . "string"` | Plain ASCII art |
| `go run . --color=<color> "string"` | Whole string colored |
| `go run . --color=<color> <substring> "string"` | Only matching substrings colored |

---

## Examples

Plain ASCII art:
```bash
go run . "hello"
```

Color the whole string red:
```bash
go run . --color=red "hello"
```

Color only matching substrings:
```bash
go run . --color=red kit "a king kitten have kit"
```
In the example above, `kit` inside `kitten` and the word `kit` at the end are both colored red.

---

## Supported Colors

| Color | Flag |
|---|---|
| Black | `--color=black` |
| Red | `--color=red` |
| Green | `--color=green` |
| Yellow | `--color=yellow` |
| Blue | `--color=blue` |
| Magenta | `--color=magenta` |
| Cyan | `--color=cyan` |
| White | `--color=white` |

Colors use ANSI escape codes and will display correctly in any ANSI-compatible terminal.

---

## Project Structure

```
.
├── main.go               # Entry point — argument routing
├── go.mod
└── master/
    ├── color.go          # ANSI color name → escape code
    ├── flags.go          # --color=<value> flag parsing
    ├── substring.go      # Substring index detection
    ├── processor.go      # ASCII art generation and rendering
    ├── color_test.go
    ├── flags_test.go
    ├── substring_test.go
    └── tests/
```

---

## Running Tests

```bash
go test ./master/...
```

---

## Error Handling

Any invalid flag format or unsupported color prints:

```
Usage: go run . [OPTION] [STRING]  EX: go run . --color=<color> <substring to be colored> "something"
```

---

## Requirements

- Go 1.22.2 or higher
- A `standard.txt` banner file in the project root
- An ANSI-compatible terminal
