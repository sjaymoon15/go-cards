package main

import "fmt"

func main() {
	cards := newDeck()
	cards.print()
	fmt.Println("_________")
	cards.shuffle()
	cards.print()
}
