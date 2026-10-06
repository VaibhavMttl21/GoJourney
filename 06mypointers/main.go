package main

import "fmt"

func main(){
	fmt.Println("Welcome to class of pointers")

	// one := 2
	// var ptr *int // datatype of pointer which is storing int type
	// fmt.Println("Value of pointer is ", ptr)

	myNumber := 23
	var ptr  = &myNumber
	fmt.Println("Value of pointer is ", ptr)
	fmt.Println("Value of pointer is ", *ptr) // value at that memory address

	*ptr = *ptr*2
	fmt.Println("New Value is ", myNumber)
}

