package main

//main function
func main(){
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
		cards := newDeck()
		cards.print()
}

//2. defining new functions - always have to declare which datatype to return
// func newCard() string { 
// //2-1. returning fixed values 
// 	return "Five of Diamonds"
// }

/* var card string = "Ace of Spades"
var    : creating new variable
card   : name of the variable
string : only this data type will be assigned to the variable 
"Ace of Spades" : assigned value to the variable 
*/

/* Why Assign Datatype?
Go = static type language (similar to C++ and Java) must assign data types 
	(Where as Dynamic Types language (Javascript , Ruby, Python) do not care assigning data types)
*/

/* Basic Go Data Types
	1. bool 	: true/false
	2. string	: "Hi" , "What's Up Everyone?"
	3. int 		: 0 , -10000 , 99999
	4. float64  : 10.001 , 0.000008 , -100.301  
*/

/*Assigining / Re-assigning variable 
	:= operator is used to assign a variable for the first time (initialized value)
	= operator is used to re-assign a variable that has been already declared 

	ex) 
		card := "Initializing Value"
		card = "Re-assigning Value"
*/

/*Go 2 Data Structures : Array , Slice
	1. Array : Fixed length list of records 
	2. Slice : An array that can grow or shrink (flexible)
	
	- Both Array and Slice can only assign identical datatype to a record inside of it
		ex) allowed : "A" "B" "C"
		    not allowed : "A" 123 true
*/

/* Iterate (Loop) Slice
for index, card := range cards {
	fmt.Println(card)
}

index : index of array
card  : current index of sets of array
range cards : slice of 'cards' to loop over
*/

/*Object Oriented Approach vs Go Approach
!! GO IS NOT AN OBJECT ORIENTED PROJECT !!

1. Object Oriented Project (Python , Ruby , Java)
 Deck Class - acts like a blue print which includes properties , methods attach to it to describe an instance (copy of deck looks like)
	ex) Deck Instance : contains properties called 'cards' in a form of an array of strings
	    Functions attached to Deck Instance : print() , shuffle() , saveToFile() to manipulate he lsit of cards attached to the deck instance


2. Go Approach 
 - define a "deck" type in slice of strings (special customized type)
 - Functions with "deck" refer it as a 'receiver' (think of it as a method)

*/

