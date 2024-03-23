package main

import (
	"fmt"
	"math"
)

// shape interface
type shape interface {
	area() float64
	circumf() float64
}

type square struct {
	length float64
}

type circle struct {
	radius float64
}

// square methods
func (s square) area() float64 {
	return s.length * s.length
}

func (s square) circumf() float64 {
	return s.length * 4
}


// circle methods
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) circumf() float64 {
	return 2 * math.Pi * c.radius
}

// common func
func printShapeInfo(shape shape){
	fmt.Printf("area of the %T is: %0.2f \n", shape, shape.area())
	fmt.Printf("circumf of the %T is: %0.2f \n", shape, shape.circumf())
}

func main() {

	shapes := []shape { // this slice is a set of circle and square strcuts
		square{length: 4},
		circle{radius: 5},
		circle{radius: 7},
	}

	for _, shape := range shapes {
		printShapeInfo(shape)
		fmt.Println("---")
	}

}
