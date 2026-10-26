package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	// "golang.org/x/oauth2/paypal"
)

func main() {
	// performgetrequest()
	// performpostjsonrequest()
	perfromformrequest()
}

func performgetrequest() {
	myurl := "http://localhost:3000/get"

	res, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close() // it's our job to close

	fmt.Println("Status Code", res.StatusCode)
	fmt.Println("Content Length", res.ContentLength)
	var responseString strings.Builder
	content, err := io.ReadAll(res.Body)
	bytecount, err := responseString.Write(content)
	if err != nil {
		panic(err)
	}
	// fmt.Println("this is my content", string(content))
	fmt.Println("Bytecount is", bytecount)
	fmt.Println(responseString.String())
	// fmt.Println(responseString)
}

func performpostjsonrequest() {
	const myurl = "http://localhost:3000/post"

	payload := strings.NewReader(`
	{
		"coursename":"golang",
		"price":0,
		"platform":"abc.com"
	}
	`)

	resp, err := http.Post(myurl, "application/json", payload)

	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	data, _ := io.ReadAll(resp.Body)
	fmt.Println("Data is:", string(data))
}

func perfromformrequest() {
	const myurl = "http://localhost:3000/postform"

	//formdata

	data := url.Values{}
	data.Add("firstname", "vaibhav")
	data.Add("lastname", "Mittal")
	data.Add("email", "abc@gmail.com")

	resp, err := http.PostForm(myurl, data)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	content,err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println("Content is",string(content))
}
