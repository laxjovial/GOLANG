# ascii-art-justify

Ascii-art-justify is a Go-based program that generates ASCII art from input text while applying text alignment and justification.

The program reads predefined banner templates (such as standard, shadow, or thinkertoy) and transforms input strings into stylized ASCII representations. It then formats the output using alignment options like left, right, center, or justify, ensuring a structured and visually consistent display.

## Usage

```bash
Usage: go run . [OPTION] [STRING] [BANNER]

go run . [--align=OPTION] "STRING" [BANNER]

Valid align options: left, right, center, justify

```

### Arguments description

| Argument | Description |
|---|---|
| `--align=OPTION` | Optional. Alignment mode: `left`, `right`, `center`, or `justify` |
| `STRING` | The text to render as ASCII art |
| `BANNER` | Optional. Banner font to use. Defaults to `standard` |

### Available Banners

- `standard` 
- `shadow` 
- `thinkertoy` 

### How It Works

- `The program reads the selected banner file.`
- `Each character in the input string is mapped to its ASCII representation.`
- `Lines are constructed row by row from the banner.`
- `The resulting ASCII text is formatted based on the chosen alignment.`
- `The final output is printed to the terminal.`

## Requirements

- Go 1.25.0

Install dependencies:

```bash
go mod tidy
```

## Authors
Odidi Esther  
Elaho Osarietin
