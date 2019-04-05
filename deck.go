package main

import "fmt"

// create a new type of deck
// whixh is a slice of strings
type deck []string

// receiver d
// any variable type deck has access to this method.
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
