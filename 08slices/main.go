package main

import "fmt"

func main() {

fmt.Println("Welcome to slices in golang")

var fruitlist = []string{"Apple","Tomato","Peach"}

fmt.Printf("type of fruitlist is %T\n",fruitlist)

fruitlist = append(fruitlist,"Mango","Banana")

fmt.Println(fruitlist)

// fruitlist = append(fruitlist[1:3])

fmt.Println(fruitlist)

highscore := make([]int,4)

highscore[0]= 234

highscore[1]= 945

highscore[2]= 465

highscore[3]= 877

highscore = append(highscore,555,666,777)

fmt.Println(highscore)

// sort.Ints(highscore)

// fmt.Println(highscore)

// how to remove a value from slices based on index

var courses = []string{"reactjs","javascript","swift","python","ruby"}

fmt.Println(courses)

var index int = 2

courses = append(courses[:index],courses[index+1:]...)

fmt.Println(courses)

}