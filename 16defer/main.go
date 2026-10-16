package main

import "fmt"

func main() {

// fmt.Println("defer in golang")

defer fmt.Println("World")

defer fmt.Println("one")

defer fmt.Println("two")

fmt.Println("Hello")

myDefer()

}

// 0,1,2,3,4 stack

// Hello, 4,3,2,1,0,two,one,World

func myDefer(){

for i:=0; i<5;i++{

defer fmt.Print(i)

}

}