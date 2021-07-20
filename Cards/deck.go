package main

import "fmt"

// Create a new type of 'deck'
// which is a slice of strings
type deck []string 

//returns new deck of cards (52 cards - 4 types , 13 numbers)
func newDeck() deck {
	cards := deck{} //no cards inside of it (initialize)

	cardSuits := []string{"Spades","Diamonds","Hearts","Clubs"}
	cardValues := []string{"Ace","Two","Three","Four","Five"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	} 

	return cards //return what i created 
}

//creating receiver for iterating loop 
func (d deck) print() {
	for i , card := range d {
		fmt.Println(i, card)
	}
}
