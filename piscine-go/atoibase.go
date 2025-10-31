package piscine

func isValidBaseAtoi(base string) bool {
	if len(base) < 2 {
		return false
	}
	for i, c := range base {
		if c == '+' || c == '-' {
			return false
		}
		for j := i + 1; j < len(base); j++ {
			if base[i] == base[j] {
				return false
			}
		}
	}
	return true
}

func indexInBase(c rune, base string) int {
	for i, b := range base {
		if b == c {
			return i
		}
	}
	return -1
}

func AtoiBase(s string, base string) int {
	if !isValidBaseAtoi(base) {
		return 0
	}

	baseLen := len(base)
	result := 0

	for _, c := range s {

		index := indexInBase(c, base)

		if index == -1 {
			return 0
		}
		result = result*baseLen + index
	}

	return result
}
