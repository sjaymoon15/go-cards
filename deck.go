package main

import "fmt"

// create a new type of deck
// whixh is a slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spaces", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			card := value + " of " + suit
			cards = append(cards, card)
		}
	}
	return cards
}

// receiver d
// any variable type deck has access to this method.
func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}
