package piscine

func indexInBaseConvert(base string, r rune) int {
	for i, br := range base {
		if br == r {
			return i
		}
	}
	return -1
}

func atoiBaseConvert(nbr string, base string) int {
	baseLen := len(base)
	result := 0
	for _, r := range nbr {
		val := indexInBaseConvert(base, r)
		if val == -1 {
			return 0
		}
		result = result*baseLen + val
	}
	return result
}

func itoaBaseConvert(num int, base string) string {
	if num == 0 {
		return string(base[0])
	}

	baseLen := len(base)
	var result []rune

	for num > 0 {
		remainder := num % baseLen
		result = append([]rune{rune(base[remainder])}, result...)
		num /= baseLen
	}

	return string(result)
}

func ConvertBase(nbr, baseFrom, baseTo string) string {
	num := atoiBaseConvert(nbr, baseFrom)
	return itoaBaseConvert(num, baseTo)
}
