package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func sub(a int, b int) int {
	return a - b
}

func main() {
	a := 50
	b := 30

	fmt.Println("Add :", add(a, b))
	fmt.Println("Sub :", sub(a, b))
}
