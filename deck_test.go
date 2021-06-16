package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length is 16, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card is Ace of Spades, but got %v", d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected last card is Four of Clubs, but got %v", d[len(d)-1])
	}
}

func TestSavetoDeckAndNewDeckTestFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFormFile("_decktesting")
	if len(loadedDeck) != 16 {
		t.Errorf("Expected deck length is 16, but got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}
