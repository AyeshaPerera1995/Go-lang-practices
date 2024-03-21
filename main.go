package main

import (
	"fmt"
	"sort"
	"strings"
)

func main(){

// ---------------------- strings package ----------------------------------
	greeting := "hello there friends!"

	fmt.Println(strings.Contains(greeting, "llo the"))

	fmt.Println(strings.ReplaceAll(greeting, "hello", "hi"))

	fmt.Println(strings.ToUpper(greeting))

	fmt.Println(strings.Index(greeting, "e"))

	fmt.Println(strings.Split(greeting, " "))



	// ---------------------- sort package ----------------------------------
	ages := []int{45, 20, 35, 30, 75, 60, 50, 25}

	sort.Ints(ages) // sorted the values
	fmt.Println(ages)

	index := sort.SearchInts(ages, 30); 
	// this function sorted the array and find the exact index of that value 
	fmt.Println(index)

	index2 := sort.SearchInts(ages, 28); 
	// [20 25 28 30 35 45 50 60 75]
	// if the number is not availble in the array, it resorted the array with new value and return the respective index
	fmt.Println(index2)


	names := []string{"yoshi", "mario", "peach", "bowser", "luigi"}

	sort.Strings(names)
	fmt.Println(names)

	fmt.Println(sort.SearchStrings(names, "peach"))
	


}