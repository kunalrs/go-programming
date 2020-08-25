package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	cards := newDeck()
	if len(cards) != 16 {
		t.Errorf("Expected deck length of 16 but got %v\n", len(cards))
	}

	if cards[0] != "Ace of Spades" {
		t.Errorf("Expected first card to be 'Ace of Spaces' but got '%v'\n", cards[0])
	}

	if cards[len(cards)-1] != "Four of Clubs" {
		t.Errorf("Expected first card to be 'Four of Clubs' but got '%v'\n", cards[len(cards)-1])
	}
}

func TestSaveToDeckandNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")
	d := newDeck()
	d.saveToFile("_decktesting")

	loadDeck := newDeckFromFile("_decktesting")

	if len(loadDeck) != 16 {
		t.Errorf("Expected deck length of 16 but got %v\n", len(loadDeck))
	}

	os.Remove("_decktesting")
}
