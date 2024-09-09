package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected deck to have 52 cards, but got %v", len(d))
	}

	if d[0] != "A of Hearts" || d[len(d) - 1] != "K of Clubs" {
		t.Errorf("Expected first card to be an A of Hearts and last card to be a K of Clubs, but got %v instead.", d[0] + " and " +d[len(d) - 1])
	}
}

func TestDeal(t *testing.T) {
	c := newDeck()

	h, d := deal(c, 5)

	if len(d) != len(c) - 5 {
		t.Errorf("Expected deal to have %v cards", len(c) - 5)
	}

	if len(h) != 5 {
		t.Errorf("Expected hand to have %v cards", 5)
	}
}

func TestShuffle(t *testing.T) {
	c := newDeck()

	d := c
	d.shuffle()

	if len(d) != len(c) {
		t.Errorf("Expected both decks to have same length")
	}

	if d.toString() != c.toString() {
		t.Errorf("Expected both decks to have different arrangement")
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting.txt")

	d := newDeck()

	d.saveToFile("_decktesting.txt")

	dFromFile := newDeckFromFile("_decktesting.txt")

	if len(dFromFile) != len(d) {
		t.Errorf("Expected to get 52 cards, but instead got %v", len(dFromFile))
	}

	if dFromFile.toString() != d.toString() {
		t.Errorf("Expected loaded deck to have same content as initial deck")
	}

	os.Remove("_decktesting.txt")
}