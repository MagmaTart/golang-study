package main

import "fmt"

type Person struct{
	name string
	age int
	gender string
}

func (person Person) Print() {
	fmt.Println("Name is", person.name)
	fmt.Println("Age is", person.age)
	fmt.Println("Gender is", person.gender)
}

func (person *Person) change_name(new string) {
	person.name = new
}

func main() {
	var me = Person{"Lee Soomin", 20, "Male"}
	me.Print()
	me.change_name("MagmaTart")
	me.Print()
}