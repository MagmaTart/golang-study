package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func main() {
	// result := add(10, 20)
	anonymous_add := func(a, b int) int {
		return a+b
	}
	result := anonymous_add(10, 20)
	fmt.Println(result)
}