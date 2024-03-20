package main

import "fmt"

func main(){

	age := 25
	name := "ayesha"
	

	// Print
	// fmt.Print("hello, \n")
	// fmt.Print("world!")

	//Println
	// fmt.Println("my age is", age, "my name is", name)

	//Printf
	fmt.Printf("my age is %v and my name is %v \n", age, name)
	fmt.Printf("my age is %q and my name is %q \n", age, name)
	fmt.Printf("age type is: %T \n", age)
	fmt.Printf("you scored %f points! \n", 225.55)
	fmt.Printf("you scored %0.1f points! \n", 225.55)

	// %_   >>  this underscore will be different specifiers
	// %v  		default format of variable
	// %q  		put quotes around the table. this is only works for strings.
	// %T  		data type of variable
	// %f   	print floating numbers
	// %0.1f  	limit the decimal points to one value 


	// Sprintf  >> save the formatted string to a new variable
	var str = fmt.Sprintf("my age is %v and my name is %v \n", age, name)
	fmt.Println("The formatted string is: ", str)



}