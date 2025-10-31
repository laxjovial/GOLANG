package piscine

func AlphaCount(s string) int {
	amt := 0
	for _, i := range s {
		if (i >= 'a' && i <= 'z') || (i >= 'A' && i <= 'Z') {
			amt++
		}
	}
	return amt
}
