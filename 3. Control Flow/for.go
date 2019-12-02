package main

import "fmt"

func for_test1() {
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println("Sum of 1~100 :", sum)
}

func for_test2() {
	sum := 0
	i := 1
	for i <= 100 {
		sum += i
		i++
	}
	fmt.Println("Sum of 1~100 :", sum)
}

func for_test3() {
	strs := []string{"AAA", "BBB", "CCC"}
	for idx, name := range strs {
		fmt.Println(idx, name)
	}
}

func main() {
	for_test1()
	for_test2()
	for_test3()
}
