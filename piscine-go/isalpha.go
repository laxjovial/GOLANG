package piscine

func IsAlpha(s string) bool {
	for _, i := range s {
		isUpper := i >= 'A' && i <= 'Z'
		isLower := i >= 'a' && i <= 'z'
		isDigit := i >= '0' && i <= '9'
		if !isUpper && !isLower && !isDigit {
			return false
		}
	}
	return true
}
