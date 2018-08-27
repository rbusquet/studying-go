package main

import (
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected deck length of 52 but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card to be Ace of Spades but got %v", d[0])
	}

	if d[len(d)-1] != "Jack of Clubs" {
		t.Errorf("Expected last card to be Jack of Clubs but got %v", d[len(d)-1])
	}
}
