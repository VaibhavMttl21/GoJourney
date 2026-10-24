package main

import (
	"fmt"
	"io"
	"net/http"
)

var url = "https://example.com"

func main() {
	fmt.Println("golang web requests ")

	response,err := http.Get(url)
	if err!= nil{
		panic(err)
	}
	fmt.Printf("Response is of type %T\n",response)
	// fmt.Println("Response is: ",response)

	defer response.Body.Close() // caller's responsibility to close connection

	databyte,err := io.ReadAll(response.Body)
	if err!= nil{
		panic(err)
	}
	fmt.Println("Response is: ",string(databyte))
}
