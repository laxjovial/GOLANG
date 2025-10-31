package piscine

func IterativeFactorial(nb int) int {
	if nb < 0 {
		return 0
	}
	if nb == 0 {
		return 1
	}
	result := 1
	for i := 2; i <= nb; i++ {
		nextResult := result * i
		if nextResult < result {
			return 0
		}
		result = nextResult
	}
	return result
}
