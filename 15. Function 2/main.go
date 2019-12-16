package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func append_two_strings(str1 string) func(string) string {
	return func(str2 string) string {
		return str1 + " " + str2
	}
}

func cubic(n int) int {
	return n * n * n
}

func get_cubic(calculator func(int) int, n int) int {
	return calculator(n)
}

func main() {
	// result := add(10, 20)
	// anonymous_add := func(a, b int) int {
	// 	return a + b
	// }
	// result := anonymous_add(10, 20)
	// fmt.Println(result)

	// first_name := "Soomin"
	// last_name := "Lee"
	// appender := append_two_strings(first_name)
	// name := appender(last_name)
	// fmt.Println(name)

	result := get_cubic(cubic, 3)
	fmt.Println(result)
}
