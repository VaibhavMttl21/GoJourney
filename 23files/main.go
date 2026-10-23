package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Welcome to the files in golang")
	content := "abcdefgh"

	file, err := os.Create("./mygofile.txt")

	// if err != nil {
	// 	panic(err)
	// 	// fmt.Println("errror",err)
	// }
	checkNillerror(err)
	length, err := io.WriteString(file, content)

	checkNillerror(err)
	fmt.Println("length is: ", length)
	defer file.Close()
	readfile("./mygofile.txt")
}

func readfile(filename string) {
	databyte, err := os.ReadFile(filename)
	checkNillerror(err)
	fmt.Println("Text data is: ", string(databyte))
}

func checkNillerror(err error) {
	if err != nil {
		panic(err)
		// fmt.Println("errror",err)
	}
}
