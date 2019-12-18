package main

import "fmt"

type Person struct{
	name string
	age int
	is_male bool
}

type Account struct{
	Person
	id, pw string
}

func main() {
	// var profile1 = Person{"Jenny", 20, false}
	// var profile2 = Person{"Robert", 25, true}
	// var profile3 = Person{name: "James", is_male: true}

	// age_sum := profile1.age + profile2.age + profile3.age
	// fmt.Println(age_sum)

	var acc1 Account
	acc1.name = "Jake"
	acc1.age = 37
	acc1.is_male = true
	acc1.id = "iamjake"
	acc1.pw = "jake8846"

	fmt.Println(acc1.name, acc1.age, acc1.is_male, acc1.id, acc1.pw)
}
