package piscine

import "github.com/01-edu/z01"

const boardSize = 8

func EightQueens() {
	var board [boardSize]int
	solve(0, &board)
}

func solve(col int, board *[boardSize]int) {
	if col == boardSize {
		printSolution(*board)
		return
	}
	for row := 0; row < boardSize; row++ {
		if isSafe(row, col, *board) {
			board[col] = row
			solve(col+1, board)
		}
	}
}

func isSafe(row, col int, board [boardSize]int) bool {
	for prevCol := 0; prevCol < col; prevCol++ {
		prevRow := board[prevCol]
		if prevRow == row {
			return false
		}
		if prevRow-prevCol == row-col || prevRow+prevCol == row+col {
			return false
		}
	}
	return true
}

func printSolution(board [boardSize]int) {
	for _, row := range board {
		z01.PrintRune(rune(row + '1'))
	}
	z01.PrintRune('\n')
}
