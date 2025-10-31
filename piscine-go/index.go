package piscine

func Index(s string, toFind string) int {
	if toFind == "" {
		return 0
	}
	len_of_s := len(s)
	len_of_f := len(toFind)
	if len_of_f > len_of_s {
		return -1
	}
	for i := 0; i <= len_of_s-len_of_f; i++ {
		j := 0
		for j < len_of_f {
			if s[i+j] != toFind[j] {
				break
			}
			j++
		}
		if j == len_of_f {
			return i
		}
	}
	return -1
}
