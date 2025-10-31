package piscine

func LoafOfBread(str string) string {
	count := 0
	for _, char := range str {
		if char != ' ' {
			count++
		}
	}
	if count > 0 && count < 5 {
		return "Invalid Output\n"
	}
	if count == 0 {
		return "\n"
	}
	result := ""
	i := 0
	first := true
	for i < len(str) {
		word := ""
		for len(word) < 5 && i < len(str) {
			if str[i] != ' ' {
				word += string(str[i])
			}
			i++
		}
		if len(word) > 0 {
			if !first {
				result += " "
			}
			result += word
			first = false
		}
		i++
	}
	return result + "\n"
}
