package main

import "fmt"

// Go is a pass-by-value language. that means,
// Go makes "copies" of values when passed into functions

	// 	Group A						Group B
	// (Non-Pointer Values)		 (Pointer Wrapper Value)
	//	----------				----------------
	//	strings						slices
	//	ints						maps
	//	floats						functions
	// 	booleans
	//	arrays
	//	structs


func updateName(x string){
	x = "wedge"
}

func updateLabel(y string) string{
	y = "pending"
	return y
}

func updateMenu(y map[string]float64){
	y["coffee"] = 2.99
	y["pie"] = 4.00
}


func main() {

	// Group A types -> strings, ints, floats, booleans, arrays, structs
	name := "tifa"
	updateName(name) // pass a copy of name variable.
	// everytime when we pass a value to a function, Go makes a copy of that value and pass it. 
	// So the original value never changed.
	fmt.Println(name)


	label := "start"
	label = updateLabel(label)
	fmt.Println(label)




	// Group B types -> slices, maps, functions
	menu := map[string]float64 {
		"pie": 5.95,
		"ice cream": 3.99,
	}

	updateMenu(menu)
	fmt.Println(menu)





}
