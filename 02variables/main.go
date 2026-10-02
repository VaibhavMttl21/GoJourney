package main

import "fmt"

const LoginToken string = "qwertyuiop" // public due to capital letters

func main() {
	var username string = "vaibhav"
	fmt.Println(username)
	fmt.Printf("variable is of type: %T \n",username)

	var isLoggedin bool = false
	fmt.Println(isLoggedin)
	fmt.Printf("variable is of type: %T \n",isLoggedin)

	var smallVal int = 2666666666
	fmt.Println(smallVal)
	fmt.Printf("variable is of type: %T \n",smallVal)

	var smallfloat float64 = 255.24234234234234
	fmt.Println(smallfloat)
	fmt.Printf("variable is of type: %T \n",smallfloat)

	// default values and aliases

	var anothervariable int
	fmt.Println(anothervariable)
	fmt.Printf("variable is of type: %T \n",anothervariable)

	// implicit type
	var website = "google.com"
	fmt.Println(website)
	
	// no var style
	numberofUsers := 3000
	fmt.Println(numberofUsers)

	fmt.Println(LoginToken)
}
