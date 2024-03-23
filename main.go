package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getInput(prompt string, r *bufio.Reader ) (string, error){
	fmt.Print(prompt)
	input, err := r.ReadString('\n')

	return strings.TrimSpace(input),err
}

func createBill() bill {
	// create a reader which is going to read from standard input (from terminal)
	reader := bufio.NewReader(os.Stdin)
	
	// call this func
	name, _ := getInput("Create a new bill name: ", reader)

	b := newBill(name)
	fmt.Println("Created the bill - ", b.name)
	return b
}

func propmtOptions(b bill){
	reader := bufio.NewReader(os.Stdin)

	opt, _ := getInput("Choose option (a - add an item, s - save bill, t - add tip): ", reader)
	fmt.Println(opt)

	switch opt {
		case "a":
			name, _ := getInput("What's the item name: ", reader)
			price, _ := getInput("What's the item price: ", reader)
			

			fmt.Println(name, price)

		case "s":
			fmt.Println("save bill")
		case "t":
			tip, _ := getInput("Enter tip amount ($): ", reader)

			fmt.Println(tip)
		default:
			fmt.Println("invalid input. Try again!")
			propmtOptions(b)

	}

}

func main() {

	mybill := createBill()
	propmtOptions(mybill)



	// ------------------------------------------------------------------------

	// mybill := newBill("mario's bill")

	// mybill.updateTip(10)
	// mybill.addItem("choco", 5.99)
	// mybill.addItem("pie", 4.99)
	// mybill.addItem("cake", 7.59)

	// // this is a receiver function which is associated with bill objects
	// fmt.Println(mybill.format())

}
