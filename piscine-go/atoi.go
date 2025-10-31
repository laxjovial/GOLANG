package piscine

func Atoi(s string) int {
	if s == "" {
		return 0
	}

	result := 0
	sign := 1
	start := 0

	if s[0] == '-' {
		sign = -1
		start = 1
	} else if s[0] == '+' {
		start = 1
	}

	for i := start; i < len(s); i++ {
		char := s[i]

		if char < '0' || char > '9' {
			return 0
		}

		digit := int(char - '0')
		result = result*10 + digit
	}

	return result * sign
}
