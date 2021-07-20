/*Intro. A Tour of Go*/

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main(){
	bs := []byte{71,111}
	fmt.Println(bs)
	fmt.Printf("%s", bs)

	fmt.Println("Hello, Ncurity Platform!")
	fmt.Println("The time is", time.Now())
	fmt.Println("Favorite Number is ... ", rand.Intn(10))
}