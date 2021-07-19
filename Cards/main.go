package main

import "fmt"

//main function
func main(){
	//var card string = "Ace of Spades"
	//Alternative way of declaring variable above (use := operator)
	//card := "Ace of Spades"
	
	card := newCard()

	fmt.Println(card)
}

//defining new functions
func newCard() string {
	return "Five of Diamond11s"
}

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