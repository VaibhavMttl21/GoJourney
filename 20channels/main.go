package main

import (
	"fmt"
	// "sync"
)

// func main() {
// 	fmt.Println("channels in go lang")

// 	myCh := make(chan int, 2)
// 	wg := &sync.WaitGroup{}

// 	// myCh <- 5
// 	// fmt.Println(<-myCh)
// 	wg.Add(2)
// 	// receive only
// 	go func(ch <-chan int, wg*sync.WaitGroup){
// 		defer wg.Done()
// 		// close(myCh)
// 		val, isChannelOpen := <- myCh
// 		fmt.Println(isChannelOpen)
// 		fmt.Println(val)
// 		// fmt.Println(<-myCh)
// 		// fmt.Println(<-myCh)
// 	}(myCh,wg)
// 		// send only
// 	go func(ch chan<- int, wg*sync.WaitGroup){
// 		defer wg.Done()
// 		// myCh<- 5
// 		// myCh<- 6
// 		// myCh<- 0
// 		close(myCh)
// 	}(myCh,wg)

// 	wg.Wait()
// }

// func main() {

//     ch := make(chan int)

//     go func() {
//         ch <- 10
//         ch <- 20
//         ch <- 30

//         close(ch)
//     }()

//     // for value := range ch {
//     //     fmt.Println(value)
//     // }
// 	fmt.Println(<-ch)
// }
            //  SENDER                    RECEIVER
            //    │                          │
            //    │      ch <- 10            │
            //    ├─────────────────────────→│
            //    │                          │
            //    │                    receives 10
            //    │                          │
            //    │      ch <- 20            │
            //    ├─────────────────────────→│
            //    │                          │
            //    │                    receives 20
            //    │                          │
            //    │      ch <- 30            │
            //    ├─────────────────────────→│
            //    │                          │
            //    │                    receives 30
            //    │                          │
            //    │      close(ch)           │
            //    ├─────────────────────────→│
            //    │                          │
            //    │                    range ends

func producer(ch chan<- int) {
    for i := 1; i <= 5; i++ {
        ch <- i
    }

    close(ch)
}

func consumer(ch <-chan int) {
    for value := range ch {
        fmt.Println(value)
    }
}

func main() {

    ch := make(chan int) // go automatically treats it with restricted direction

    go producer(ch)

    consumer(ch)
}