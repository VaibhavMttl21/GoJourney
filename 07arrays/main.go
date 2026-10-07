package main

import "fmt"

func main() {

fmt.Println("Welcome to array in golang")

var fruitlist [4]string

fruitlist[0] ="Apple"

fruitlist[1] = "Orange"

fruitlist[3] = "peach"

fmt.Println("fruit list is:", fruitlist)

fmt.Println("fruit list is:", len(fruitlist))

var veglist= [3]string{"potato","beans","tomato"}

fmt.Println("fruit list is:", veglist)

}