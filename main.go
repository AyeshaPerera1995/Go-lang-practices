package main

import "fmt"

func main() {

	var ages [3]int = [3]int{20, 25, 30}
	var scoress = [3]int{10, 45, 80}
	// arrays have fixed legth

	names := [4]string{"yoshi", "mario", "peach", "bowser"}
	names[1] = "ayesha"

	fmt.Println(ages, len(ages), scoress)
	fmt.Println(len(names), names)

	// slices >> slices don't have fix length. you can append values
	var scores = []int{100, 50, 60}
	scores[2] = 25
	scores = append(scores, 40, 20)
	// we can append elements or another slice using above function

	fmt.Println(scores, len(scores))

	// slice ranges
	rangeOne := names[1:3]
	rangeTwo := names[2:]
	rangeThree := names[:]
	rangeFour := names[:3]

	fmt.Println(rangeOne, rangeTwo, rangeThree, rangeFour)

	// you can add another value to range slice
	rangeOne = append(rangeOne, "selva")
	fmt.Println(rangeOne)

}
