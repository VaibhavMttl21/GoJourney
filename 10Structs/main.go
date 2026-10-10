package main

import "fmt"

func main() {

fmt.Println("Structs in golang")

vaibhav := User{"vaibhav","vaibhav@gamil.com",true,21}

fmt.Println(vaibhav)

fmt.Printf("vaibhav details are: %+v\n",vaibhav)

fmt.Printf("Name is %v and email is %v.",vaibhav.Name,vaibhav.Email)


}

type User struct {

Name string

Email string

Status bool

Age int

}