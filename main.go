package main

import (
	"fmt"
	"net/http"
	"sync"
)

var wg = sync.WaitGroup{}

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
	// go greeter("Hello")
	// greeter("World!")
}

func getStatusCode(endpoint string){
	defer wg.Done()
	
	res, err := http.Get(endpoint)
	if err != nil {
		fmt.Println("OOPS in endpoint!")
	}
	fmt.Printf("%v status code for %v \n", res.StatusCode, endpoint)

}








// func greeter(s string) {
// 	for i := 0; i < 5; i++ {
// 		time.Sleep(3 * time.Second)
// 		fmt.Println(s)
// 	}
// }