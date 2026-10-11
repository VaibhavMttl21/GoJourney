package main

import "fmt"

func main() {

fmt.Println("If else in golang")

loginCount := 23

var result string

if loginCount <10 {

result ="Regular user"

}else if loginCount>10{

result ="Watch Out"

}else{

result = "exactly 10 login count"

}

if 9%2==0{

fmt.Println("Number is even")

}else{

fmt.Println("Number is odd")

}

if num:=3; num<10{

fmt.Println("Num is less than 10")

}else{

fmt.Println("Num is NOT less than 10")

}

fmt.Println(result)


}