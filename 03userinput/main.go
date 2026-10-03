package main

import (
	"fmt"
	"os"
	"bufio"
)

func main() {
	welcome := "Welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your name: ")
	 //comma ok syntax
	 v,_ := reader.ReadString('\n')
	 fmt.Println("Thank for reading", v)
	 fmt.Printf("type of input is: %T", v)

}
