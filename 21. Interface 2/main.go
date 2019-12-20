package main

import "fmt"

type Animal interface {
	eat()
}

type Human interface {
	Animal
	speak()
}

type Person struct {
	name string
}

func (p Person) eat() {
	fmt.Println("Person : Eat")
}

func (p Person) speak() {
	fmt.Println("Person : Speak")
}

func HumanTest(h Human) {
	h.eat()
	h.speak()
}

func main() {
	var person = Person{"Soomin Lee"}
	HumanTest(person)
}
