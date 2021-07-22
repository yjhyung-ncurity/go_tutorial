package main

import (
	"fmt"

	"ncurity.com/test/assignment"
)

func main() {
	//1. Declaring a Variable
	//var card string = "Ace of Spades"
	//1-1. Alternative way of declaring variable above (use := operator)
	//card := "Ace of Spades"

	//2. Getting the value by calling other function
	//2-1. calling functions giving fixed return values "Five of Diamonds"
	//card := newCard()

	//2-2. creating multiple cards using slice
	//cards := []string{"Ace of Diamonds", newCard()}
	//cards = append(cards,"Six of Spades")

	//for i, card := range cards {
	//fmt.Println(i, card)
	//}

	//2-3. creating 'deck' type GO Approach
	// cards := deck{"Ace of Diamonds", newCard()}
	// cards = append(cards, "Seven of Spades")

	//2-4. calling new deck of cards created in deck.go file
	// cards := newDeck()

	// hand, remainingCards :=  deal(cards, 5)

	// hand.print()
	// remainingCards.print()

	//2-5. converting deck type into a slice of string then into a string
	// cards := newDeck()
	// fmt.Println(cards.toString())
	// cards.saveToFile("my_cards")

	// 2-6. Reading File from a HardDrive
	// cards := newDeckFromFile("my_cards")
	// cards.print()

	// 2-7. Shuffling the cards using pseudo random generator rand.Intn()
	// cards := deck.NewDeck()
	// cards.Print()
	//result: randomized but always shows the same random results

	//assignment #1 : create a slice of ints from 0 ~ 10 then iterate through a loop and return "odd" if it's odd and "even" if it's even
	result_test := assignment.NewNum()
	fmt.Println(result_test)

}
