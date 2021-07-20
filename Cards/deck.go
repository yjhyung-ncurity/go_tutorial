package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new type of 'deck'
// which is a slice of strings
type deck []string

//returns new deck of cards (52 cards - 4 types , 13 numbers)
func newDeck() deck {
	cards := deck{} //no cards inside of it (initialize)

	//instead of writing out all 52 types of cards, create a two slice and use double iterations
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards //return what i created
}

//creating receiver for iterating loop
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

//deal function - dealing deck of cards
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

//turning string set of slice into a string
//converting process : custom type -> slice of string -> string (toString)
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)

	if err != nil {
		//error handling option 1 : log the error and return a call to newDeck()
		//error handling option 2 : log the error and entirely quit the program
		fmt.Println("Error: ", err)
		os.Exit(1) //quiting program entirely
	}

	//converting : byte slice -> string -> slice of strings -> deck type(custom type)
	s := strings.Split(string(bs), ",")

	return deck(s)

}

//randomize the deck order
func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)

		d[i], d[newPosition] = d[newPosition], d[i] //swap the elements in both positions
	}
}
