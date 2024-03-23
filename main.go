package main

import (
	"fmt"
	"net/http"
	"sync"
)

var signals = []string {"test"}

var wg sync.WaitGroup // pointer
var mutex sync.Mutex // pointer

func main() {

	websitelist := []string {
		"https://go.dev",
		"https://google.com",
		"https://fb.com",
		"https://github.com",
	}

	for _, web := range websitelist {
		wg.Add(1)
		go getStatusCode(web)
	}

	wg.Wait()
	fmt.Println(signals)


	// go greeter("Hello")
	// greeter("World!")
}

func getStatusCode(endpoint string){
	defer wg.Done()

	res, err := http.Get(endpoint)
	if err != nil {
		fmt.Println("OOPS in endpoint!")
	}
	mutex.Lock()
	signals = append(signals, endpoint)
	mutex.Unlock()

	fmt.Printf("%v status code for %v \n", res.StatusCode, endpoint)

}








// func greeter(s string) {
// 	for i := 0; i < 5; i++ {
// 		time.Sleep(3 * time.Second)
// 		fmt.Println(s)
// 	}
// }