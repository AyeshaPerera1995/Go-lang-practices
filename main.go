package main

import "fmt"

func main(){
	fmt.Println("Hello!")

	// strings ----------------------------------------------------------------
	var nameOne string = "mario"
	var nameTwo = "luigi"
	var nameThree string // this has an empty string

	fmt.Println(nameOne, nameTwo, nameThree)

	nameOne = "peach"
	nameThree = "bowser"

	fmt.Println(nameOne, nameTwo, nameThree)

	nameFour := "yoshi" 
	// this is a short hand of initializing a variable and you can use it only in the very first time.
	// you can't use this way, outside a function

	fmt.Println(nameFour)






	// ints ----------------------------------------------------------------	
	var ageOne int = 20
	var ageTwo = 30
	ageThree := 40

	fmt.Println(ageOne, ageTwo, ageThree)

	// bits and memory
	var numOne int8 = 25
	var numTwo int8 = -128
	var numThree uint16 = 256

	fmt.Println(numOne, numTwo, numThree)

	var scoreOne float32 = 25.45
	var scoreTwo float64 = 2444444444444444444444444444444445.458777777
	scoreThree := 1.5

	fmt.Println(scoreOne, scoreTwo, scoreThree)

}