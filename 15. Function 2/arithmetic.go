package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}

func mul(a, b int) int {
	return a * b
}

func div(a, b int) int {
	return a / b
}

func arithmetic(op func(int, int) int, a int, b int) int {
	return op(a, b)
}

func main() {
	a, b := 100, 30
	add_result := arithmetic(add, a, b)
	sub_result := arithmetic(sub, a, b)
	mul_result := arithmetic(mul, a, b)
	div_result := arithmetic(div, a, b)
	fmt.Println(add_result, sub_result, mul_result, div_result)
}
