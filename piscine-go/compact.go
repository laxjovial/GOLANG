package piscine

func Compact(ptr *[]string) int {
	if ptr == nil {
		return 0
	}
	slice := *ptr
	count := 0
	for i := 0; i < len(slice); i++ {
		if slice[i] != "" {
			slice[count] = slice[i]
			count++
		}
	}
	*ptr = slice[:count]
	return count
}
