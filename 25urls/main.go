package main

import (
	"fmt"
	"net/url"
)

var myURL = "https://lco.dev:3000/learn?course=reactjs&paymentid=asd234fdd"

func main() {
	fmt.Println("Welcome to URLs in Golang")
	fmt.Println(myURL)

	result, err := url.Parse(myURL)
	if err != nil {
		panic(err)
	}

	fmt.Println("Scheme:", result.Scheme)
	fmt.Println("Host:", result.Host)
	fmt.Println("Path:", result.Path)
	fmt.Println("RawQuery:", result.RawQuery)
	fmt.Println("Port:", result.Port())

	qparam := result.Query()

	fmt.Printf("Type of qparam: %T\n", qparam)

	fmt.Println("Course:", qparam.Get("course"))
	fmt.Println("Payment ID:", qparam.Get("paymentid"))

	for key, val := range qparam {
		fmt.Println("Param:", key, val)
	}

	query := url.Values{}
	query.Set("user", "hitesh")

	partsOfURL := &url.URL{
		Scheme:   "https",
		Host:     "lco.dev",
		Path:     "/tutcss",
		RawQuery: query.Encode(),
	}

	fmt.Println(partsOfURL.String())
}