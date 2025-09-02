package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first element of the deck being 'Ace of Spades', but got %v", d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected first element of the deck being 'Four of Clubs', but got %v", d[len(d)-1])
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	filename := "_decktesting"
	removeFile := func() { os.Remove(filename) }
	removeFile()
	defer removeFile()

	deck := newDeck()
	deck.saveToFile(filename)

	loadedDeck := newDeckFromFile(filename)

	if len(loadedDeck) != len(deck) {
		t.Errorf("Expected same deck size, but got %v", len(loadedDeck))
	}

	os.Remove(filename)
}
