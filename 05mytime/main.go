package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Printf("welcome to timestudy of golang \n")

	presenttime := time.Now()
	fmt.Println(presenttime)

	fmt.Println(presenttime.Format("01-02-2006 15:04:05 Monday"))

	createdate := time.Date(2020,time.January,10,23,23,0,0,time.UTC)
	fmt.Println(createdate)
	fmt.Println(createdate.Format("01-02-2006 Monday"))
}
