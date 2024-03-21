package main

import "fmt"

func main() {

	x := 0

	// in Go while alos equal to for
	for x < 5 {
		fmt.Println("value of x is: ", x)
		x++
	}

	for i := 0; i < 5; i++ {
		fmt.Println("value of i is: ", i)
	}

	names := []string{"mario", "luigi", "yoshi", "peach"}

	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	for index, value := range names {
		fmt.Printf("the value at index %v is %v \n", index, value)
	}

	for _, value := range names {
		fmt.Printf("the value is %v \n", value)
	}
	// if we don't wnat to use any variable, you can simply use undersocore

}
