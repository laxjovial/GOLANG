package piscine

import "fmt"

func DealAPackOfCards(deck []int) {
	cardsPerPlayer := 3
	for i := 0; i < 4; i++ {
		fmt.Printf("Player %d: ", i+1)
		start := i * cardsPerPlayer
		for j := 0; j < cardsPerPlayer; j++ {
			fmt.Printf("%d", deck[start+j])
			if j < cardsPerPlayer-1 {
				fmt.Printf(", ")
			}
		}
		fmt.Printf("\n")
	}
}
