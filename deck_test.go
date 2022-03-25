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
		t.Errorf("Expected our first card to be the Ace of Spades but got %v", d[0])
	}
	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("expected last card to be the Four of Clubs but got %v", d[len(d)-1])
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {

	os.Remove("_decktesting")
	deck := newDeck()
	deck.saveToFile("_decktesting")
	loadedDeck := newDeckFromFile("_deckTesting")

	if len(deck) != len(loadedDeck) {
		t.Errorf("Expected the length of deck and loadedDeck to match.")
	}

	os.Remove("_deckTesting")
}

func TestDeal(t *testing.T) {

	deck := newDeck()
	hand, remaining := deal(deck, 5)

	if len(remaining) != 11 {
		t.Errorf("Expected remaining deck len to be 11 after dealing five cards. Instead, it was %v", len(remaining))
	}

	if len(hand) != 5 {
		t.Errorf("expected 5 cards to be in hand. Instead, got %v", len(hand))
	}

}
