package main

import "fmt"

func main() {
	
	menu := map[string]float64 {  // keys are strings and the values are floats
		"soup": 4.99,
		"pie": 7.99,
		"salad": 6.99,
		"toffee pudding": 3.55,
	}
	fmt.Println(menu)
	fmt.Println(menu["pie"])

	for k, v := range menu {
		fmt.Println(k, "-", v)
	}

	// ints as key type
	phonebook := map[int]string{
		267584967: "mario",
		345784398: "luigi",
		356482456: "peach",
	}

	fmt.Println(phonebook)
	fmt.Println(phonebook[345784398])

	// update an item
	phonebook[345784398] = "bowser"
	fmt.Println(phonebook)

}
