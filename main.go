package main

import "fmt"

func updateName(x *string) {
	*x = "wedge"
}

func main() {
	// pointer is just a pointer to a memory location
	name := "ayesha"

	// updateName(name)

	fmt.Println("memory address of name is: ", &name)
	// prints the memory location of the name

	m := &name
	// fmt.Println("memory address:", m)
	// fmt.Println("value of memory address:", *m)

	// pass a pointer
	fmt.Println(name)
	updateName(m)
	fmt.Println(name)



}
