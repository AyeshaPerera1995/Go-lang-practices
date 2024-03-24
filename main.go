package main

import (
	"fmt"
	"sync"
)

func main() {

	myCh := make(chan int, 2)
	wg := &sync.WaitGroup{}

	wg.Add(2)

	// RECEIVE DATA FROM CHANNEL ONLY
	go func (ch <-chan int, wg *sync.WaitGroup){
		val, isChannelOpen := <-myCh

		fmt.Println(isChannelOpen)
		fmt.Println(val)

		// fmt.Println(<-myCh)
		// fmt.Println(<-myCh)
		wg.Done()
	}(myCh, wg)


	// SEND DATA TO CHANNEL ONLY
	go func (ch chan<- int, wg *sync.WaitGroup){
		myCh <- 5
		// myCh <- 6
		close(myCh)
		wg.Done()
	}(myCh, wg)

	wg.Wait()

	// myCh <- 5
	// myCh is kind of a box and using arrow, somebody is putting values inside the box.	
	// fmt.Println(<-myCh)
	// this is like values are coming outof the box.(arrow is in other side)

}
