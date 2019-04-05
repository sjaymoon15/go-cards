package main

import "fmt"

func main() {
	cards := deck{newCard(), newCard()}
	cards = append(cards, "Six of Spades")

	for card := range cards {
		fmt.Println(card)
	}
	// fmt.Println(cards)
}

func newCard() string {
	return "Five of Diamonds"
}
