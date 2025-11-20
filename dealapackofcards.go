package piscine

import (
	"fmt"

	"github.com/01-edu/z01"
)

func DealAPackOfCards(deck []int) {
	// We know the deck always has 12 cards and 4 players
	cardsPerPlayer := 3
	players := 4

	for i := 0; i < players; i++ {
		// Print "Player X: "
		fmt.Printf("Player %d: ", i+1)

		for j := 0; j < cardsPerPlayer; j++ {
			card := deck[i*cardsPerPlayer+j]
			if j == cardsPerPlayer-1 {
				fmt.Printf("%d", card)
			} else {
				fmt.Printf("%d, ", card)
			}
		}

		// End line using z01.PrintRune
		z01.PrintRune('\n')
	}
}
