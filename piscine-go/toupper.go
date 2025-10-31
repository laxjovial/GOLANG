package piscine

func ToUpper(s string) string {
	var result string
	for _, i := range s {
		if i >= 'a' && i <= 'z' {
			result += string(i - 32)
		} else {
			result += string(i)
		}
	}
	return result
}
