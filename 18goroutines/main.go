package main

import (
	"fmt"
	"net/http"
	"sync"
	// "time"
)
var signals = []string{"test"}
var wg sync.WaitGroup // usually these are pointers
var mt sync.Mutex // usually these are pointer 
func main() {
	// go greeter("Hello")
	// if i remove the sleep time it jsut not waited that thread to comeback the hello thread
	// that's why we see only World to print we have to wait for that hello thrad to comeback
	// greeter("World")
	//And this is actually the key idea of concurrency: both tasks are making progress during overlapping periods, but their execution order isn't fixed.
	// go greeter("Hello")

	// time.Sleep(30 * time.Millisecond)

	// greeter("World")
	websiteList := []string{
		"https://lco.dev",
		"https://go.dev",
		"https://google.com",
		"https://fb.com",
		"https://github.com",
	}
	for _, web := range websiteList {
		go getStatuscode(web)
		wg.Add(1)
	}
	wg.Wait()
	fmt.Println(signals)
}

// func greeter(s string) {
// 	for i := 0; i < 6; i++ {
// 		time.Sleep(3 * time.Millisecond)
// 		fmt.Println(s)
// 	}
// }

func getStatuscode(endpoint string) {
	defer wg.Done()
	res, err := http.Get(endpoint)

	if err != nil {
		fmt.Println("OOps in endpoint")
	}else {
		mt.Lock()
		signals = append(signals,endpoint)
		mt.Unlock()
		fmt.Printf("%d status code for %s\n", res.StatusCode, endpoint)
	}
}
