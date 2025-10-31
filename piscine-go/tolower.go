package piscine

func ToLower(s string) string {
	var result string
	for _, i := range s {
		if i >= 'A' && i <= 'Z' {
			result += string(i + 32)
		} else {
			result += string(i)
		}
	}
	return result
}
