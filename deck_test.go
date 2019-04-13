package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected 16 but got %v", len(d))
	}

	if d[0] != "Ace of Spaces" {
		t.Errorf("Expected Ace of Spaces for the first element of deck but got %v", d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected Four of Clubs for the last element of deck but got %v", d[len(d)-1])
	}
}

func TestSaveToFileNewDeckFromFile(t *testing.T) {
	os.Remove("_testingDeck")
	d := newDeck()
	d.saveToFile("_testingDeck")
	loadedDeck := newDeckFromFile("_testingDeck")
	if len(loadedDeck) != 16 {
		t.Errorf("Expected 16 but got %v", len(loadedDeck))
	}

	os.Remove("_testingDeck")
}
