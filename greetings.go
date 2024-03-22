package main

import "fmt"

var points = []int{20, 90, 100, 45, 70}

func sayHello(n string) { // since these are in same package (main), you can use these things in main() func.
	fmt.Println("hello", n)
}

func showScore(){
	fmt.Println(score)
}