package main

import (
    "fmt"
    "time"
)

func main() {

    ch1 := make(chan int)
    ch2 := make(chan int)

    go func() {
        time.Sleep(1 * time.Second)
        ch1 <- 10
    }()

    go func() {
        time.Sleep(2 * time.Second)
        ch2 <- 20
    }()

    for {
        select {
        case msg := <-ch1:
            fmt.Println("ch1:", msg)

        case msg := <-ch2:
            fmt.Println("ch2:", msg)
        }
    }
}

// func main() {

//     ch := make(chan int)

//     select {
//     case value := <-ch:
//         fmt.Println(value)

//     default:
//         fmt.Println("No value available")
//     }
// }