package main

import (
	"os"
	"unicode"

	"github.com/01-edu/z01"
)

const size = 9

func main() {
	if len(os.Args) != 10 {
		printError()
		return
	}

	board := [size][size]byte{}

	for i := 0; i < size; i++ {
		row := os.Args[i+1]
		if len(row) != size {
			printError()
			return
		}
		for j := 0; j < size; j++ {
			ch := row[j]
			if ch != '.' && (!unicode.IsDigit(rune(ch)) || ch == '0') {
				printError()
				return
			}
			board[i][j] = ch
		}
	}

	if !isValidBoard(board) || !solveSudoku(&board) {
		printError()
		return
	}

	printBoard(board)
}

func printError() {
	for _, r := range "Error\n" {
		z01.PrintRune(r)
	}
}

func isValid(board *[size][size]byte, row, col int, ch byte) bool {
	for i := 0; i < size; i++ {
		if board[row][i] == ch || board[i][col] == ch ||
			board[3*(row/3)+i/3][3*(col/3)+i%3] == ch {
			return false
		}
	}
	return true
}

func solveSudoku(board *[size][size]byte) bool {
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if board[i][j] == '.' {
				for ch := byte('1'); ch <= '9'; ch++ {
					if isValid(board, i, j, ch) {
						board[i][j] = ch
						if solveSudoku(board) {
							return true
						}
						board[i][j] = '.'
					}
				}
				return false
			}
		}
	}
	return true
}

func isValidBoard(board [size][size]byte) bool {
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			ch := board[i][j]
			if ch != '.' {
				board[i][j] = '.'
				if !isValid(&board, i, j, ch) {
					return false
				}
				board[i][j] = ch
			}
		}
	}
	return true
}

func printBoard(board [size][size]byte) {
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			z01.PrintRune(rune(board[i][j]))
			if j < size-1 {
				z01.PrintRune(' ')
			}
		}
		z01.PrintRune('\n')
	}
}
