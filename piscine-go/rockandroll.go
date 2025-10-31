package piscine

func RockAndRoll(n int) string {
	if n < 0 {
		return "error: number is negative\n"
	}

	DivideBy2 := n%2 == 0
	DivideBy3 := n%3 == 0

	if DivideBy2 && DivideBy3 {
		return "rock and roll\n"
	}

	if DivideBy2 {
		return "rock\n"
	}

	if DivideBy3 {
		return "roll\n"
	}

	return "error: non divisible\n"
}
