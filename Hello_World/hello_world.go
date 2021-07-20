/*Intro. A Tour of Go*/

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main(){

	fmt.Println("Hello, Ncurity Platform!")
	fmt.Println("The time is", time.Now())
	fmt.Println("Favorite Number is ... ", rand.Intn(10))
}