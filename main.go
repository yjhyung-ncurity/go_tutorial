package main

import "fmt"

//PART 3 main function - not using

//func main() {
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
// result_test := assignment.NewNum()
// fmt.Println(result_test)
////// NOT FINISHED ///////

//}

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

//PART 4 Struts main function

//Lesson 1. Define struct
/*
type person struct { //define custom type called person in struct
	//Define a field (properties that person struct might have like lastname, firstname)
	firstName string
	lastName  string
}

func main() {

		//Lesson 1. Define & Declare struct and its fields

		//Ways to declare
		1. alex := person("Alex", "Andersen")
		2. alex := person{firstName: "Alex", lastName: "Andersen"}
		3.
		var alex person

		alex.firstName = "Alex"
		alex.lastName = "Andersen"
		fmt.Println(alex)
		fmt.Printf("%+v\n", alex)
}
*/

//Lesson 2. How to Embedd a struct into another struct
/*
type contactInfo struct {
	email   string
	zipcode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo //custom type called contactInfo declared above or just declare it as "contactInfo"
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Smith",
		contact: contactInfo{
			email:   "JSmit@gmail.com",
			zipcode: 94000,
		},
	}

	fmt.Printf("%+v", jim)
}
*/

//Lesson 3. Creating a receiver using Struct && Struct with Pointers
/*
type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Walker",
		contactInfo: contactInfo{
			email:   "jw1123@naver.com",
			zipCode: 80524,
		},
	}


	//fmt.Printf("%+v\n", jim)
	jimPointer := &jim
	jimPointer.updateName("James")
	jim.print()
}

//receiver function - to print out a person's info

func (p person) print() {
	fmt.Printf("%+v\n", p)
}

//receiver function - to update a person's firstname
func (pointerToPerson *person) updateName(newFirstName string) {
	//p.firstName = newFirstName //update that person's firstName (no pointer -> result: name changes but does not show on print)

	//when using the pointer, remove the existing receiver "p person" and replace it with pointerToPerson *person
	(*pointerToPerson).firstName = newFirstName
}
*/

/*
//Lesson 4. Pointer Shortcut
type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Walker",
		contactInfo: contactInfo{
			email:   "jw1123@naver.com",
			zipCode: 80524,
		},
	}

	fmt.Printf("%+v\n", jim)
	jim.updateName("James")
	jim.print()
}

//receiver function - to print out a person's info

func (p person) print() {
	fmt.Printf("%+v\n", p)
}

//receiver function - to update a person's firstname
func (pointerToPerson *person) updateName(newFirstName string) {
	//p.firstName = newFirstName //update that person's firstName (no pointer -> result: name changes but does not show on print)

	//when using the pointer, remove the existing receiver "p person" and replace it with pointerToPerson *person
	(*pointerToPerson).firstName = newFirstName
}
*/

//BONUS: JWLEE STRUCT POINTERS TUTORIAL - what ncurity actually use
/*
type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Walker",
		contactInfo: contactInfo{
			email:   "jw1123@naver.com",
			zipCode: 80524,
		},
	}

	jim.firstName = "James" //이렇게 사용
	jim.print()
}

//receiver function - to print out a person's info

func (p person) print() {
	fmt.Printf("%+v\n", p)
}
*/

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

//PART 5 Map

//1. declaring a map
//2. how to delete a map
// func main() {
//1. declaring a map - literal syntax (declaring with initial values)
// colors := map[string]string{
// 	"red":   "#ff0000",
// 	"green": "#4bf745",
// 	"white": "#ffffff",
// }

//2. declaring a map - using var syntax
// var colors map[string]string //declared an empty map with no key value inside

//3. declaring a map - using make syntax
// colors := make(map[string]string) //declared an empty map

// colors["white"] = "#ffffff"
// colors["black"] = "#000000"

//4. delete using delete()
// delete(colors, "white")
// 	fmt.Println(colors)
// }

//3. Manipulating Maps
// function that accepts a map, iterates over and prints out every key value pair inside of the map
// shows how to iterate over a collection of key value
// how to pass a map off to another function
func main() {
	//1. declaring a map - literal syntax (declaring with initial values)
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
	}

	//fmt.Println(colors)
	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for ", color, "is", hex)
	}
}
