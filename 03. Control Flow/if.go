package main

import "fmt"

func if_test1() {
	a := 10
	if a > 5 {
		fmt.Println("a is greater than 5")
	}
}

func if_test2() {
	a := 70
	if a < 60 {
		fmt.Println("a < 60")
	} else if a > 60 && a < 70 {
		fmt.Println("60 < a < 70")
	} else {
		fmt.Println("a > 70")
	}
}

func if_test3() {
	a := 5
	if num := a * a; num > 20 {
		fmt.Println("num is greater than 20")
	} else {
		fmt.Println("num is equal or smaller than 20")
	}
}

func main() {
	if_test1()
	if_test2()
	if_test3()
}
