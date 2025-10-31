package piscine

func PodiumPosition(podium [][]string) [][]string {
	n := len(podium)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if podium[j][0] > podium[j+1][0] {
				podium[j], podium[j+1] = podium[j+1], podium[j]
			}
		}
	}
	return podium
}
