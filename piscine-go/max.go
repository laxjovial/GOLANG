package piscine

func Max(a []int) int {
	if len(a) == 0 {
		return 0
	}

	maxVal := a[0]

	for i := 1; i < len(a); i++ {
		if a[i] > maxVal {
			maxVal = a[i]
		}
	}

	return maxVal
}
