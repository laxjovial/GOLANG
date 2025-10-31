package piscine

func RecursiveFactorial(nb int) int {
	if nb < 0 {
		return 0
	}
	if nb == 0 {
		return 1
	}
	if nb > 20 {
		return 0
	}
	prevFactorial := RecursiveFactorial(nb - 1)
	if prevFactorial > 0 && (nb*prevFactorial)/nb != prevFactorial {
		return 0
	}
	return nb * prevFactorial
}
