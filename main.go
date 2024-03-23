package main

import "fmt"

func main() {

	mybill := newBill("mario's bill")

	mybill.updateTip(10)
	mybill.addItem("choco", 5.99)
	mybill.addItem("pie", 4.99)
	mybill.addItem("cake", 7.59)

	// this is a receiver function which is associated with bill objects
	fmt.Println(mybill.format())

}
