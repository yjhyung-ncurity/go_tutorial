package main

import (
	"os"
	"testing"

	"ncurity.com/test/deck"
)

func TestNewDeck(t *testing.T) {
	//argument t is test handler
	d := deck.NewDeck()

	//To make sure deck is created with x number of cards (in this case total of 16 cards). If not, returns error message
	if len(d) != 16 {
		//formatted string : can inject the value into a string (%v injects value of length of d)
		t.Errorf("Expected deck length of 20, but got %v", len(d))
	}

	//To make sure the first card is an Ace of Spades. If not, returns error message
	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of deck is Ace of Spades, but got %v", d[0])
	}

	//To make sure the last card is Four of Clubs. If not, returns error message
	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected last card of deck is Four of Clubs, but got %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	//1. Use Remove() from provided by os package to remove any crashed , unnecessary files
	os.Remove("_decktesting")

	new_deck := deck.NewDeck()
	new_deck.SaveToFile("_decktesting")

	loadedDeck := deck.NewDeckFromFile("_decktesting")

	if len(loadedDeck) != 16 {
		t.Errorf("Expected 16 cards in deck, got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}
