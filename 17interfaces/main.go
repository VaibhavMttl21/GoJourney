package main
import "fmt"

func main() {
  dog := Dog{}
  cat := Cat{}
  makeSound(dog)
  makeSound(cat)
}

type Speaker interface{
  Speak()
}

type Dog struct{}

func (d Dog) Speak(){
  fmt.Println("Woof")
}

type Cat struct{}

func (c Cat) Speak(){
  fmt.Println("Meow")
}

func makeSound(s Speaker){
  s.Speak()
}