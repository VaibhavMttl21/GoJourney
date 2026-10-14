package main

import(

"fmt"

)

func main() {

fmt.Println("Functions in golang")

greeter()

result := adder(3,5)

fmt.Println("Result is:", result)

prores, mymessage := proAdder(2,5,8,7)

fmt.Println("Pro result is: ",prores)

fmt.Println("Pro Message is: ",mymessage)

// not allowed we have to declare it outside

// func greeterTwo(){

// fmt.Println("Another method")

// }

// greeterTwo()

}

// func greeterTwo(){

// fmt.Println("Another method")

// }

func adder(valone int, valtwo int) int{

return valone+valtwo

}

func proAdder(values...int)(int,string){

total := 0

for _,val := range values{

total += val

}

return total, "hi Pro result function"

}

func greeter(){

fmt.Println("Namastey from golang")

}