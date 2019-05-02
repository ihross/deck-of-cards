package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected 'Ace of Spades' but got badk %v", d[0])
	}

	if d[len(d)-1] != "King of Clubs" {
		t.Errorf("Expected last card to be 'King of Clubs' but got back %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 56 {
		t.Errorf("Expected 56 cards in the deck but got back %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}
