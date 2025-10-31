package piscine

import "github.com/01-edu/z01"

func PrintCombN(n int) {
	if n <= 0 || n >= 10 {
		return
	}
	combination := make([]int, n)
	printCombHelper(n, 0, 0, combination)
	z01.PrintRune('\n')
}

func printCombHelper(n, pos, start int, combination []int) {
	if pos == n {
		printCombination(combination, n)
		return
	}
	for i := start; i <= 9; i++ {
		combination[pos] = i
		printCombHelper(n, pos+1, i+1, combination)
	}
}

func printCombination(combination []int, n int) {
	for i := 0; i < n; i++ {
		z01.PrintRune(rune(combination[i] + '0'))
	}
	if !isLastCombination(combination, n) {
		z01.PrintRune(',')
		z01.PrintRune(' ')
	}
}

func isLastCombination(combination []int, n int) bool {
	for i := 0; i < n; i++ {
		if combination[i] != 10-n+i {
			return false
		}
	}
	return true
}
