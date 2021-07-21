package main

import (
	"fmt"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	fmt.Printf("xxxxxxxxx %v", len(d))

	if len(d) != 16 {
		fmt.Errorf("expected deck length of 20, but got %v", len(d))
	}
}
