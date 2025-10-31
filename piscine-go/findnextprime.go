package piscine

func FindNextPrime(nb int) int {
	if nb <= 2 {
		return 2
	}
	if nb%2 == 0 {
		nb++
	}
	for {
		isPrime := true
		for i := 3; i*i <= nb; i += 2 {
			if nb%i == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			return nb
		}
		nb += 2
	}
}
