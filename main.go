package main

import (
	"fmt"
	"math"
	"strings"
)

func sayGreeting(n string) {
	fmt.Printf("Good morning %v \n", n)
}

func sayBye(n string) {
	fmt.Printf("Goodbye %v \n", n)
}

func cycleNames(n []string, f func(string)) {
	for _, v := range n {
		f(v)
	}
}

func circleArea(r float64) float64 {
	return math.Pi * r * r
}

func getInitials(n string) (string, string){
	s := strings.ToUpper(n)
	names := strings.Split(s, " ")
	
	var initials []string

	for _, v := range names {
		initials = append(initials, v[:1])
	}

	if len(initials) > 1 {
		return initials[0], initials[1]
	}

	return initials[0], "_"
}

func main() {
	// sayGreeting("mario")
	// sayBye("ayesha")

	// cycleNames([]string{"cloud", "tifa", "barret"}, sayGreeting)
	// cycleNames([]string{"cloud", "tifa", "barret"}, sayBye)

	// a1 := circleArea(7)
	// a2 := circleArea(10.5)

	// fmt.Println(a1, a2)
	// fmt.Printf("circle 1 is %0.3f and circle 2 is %0.2f \n", a1, a2)

	s1, s2 := getInitials("tifa lockhart")
	fmt.Println(s1, s2)



}
