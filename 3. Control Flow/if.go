package main

import "fmt"

func main() {
	a := 5
	if num := a*a; num < 25 {
		fmt.Println("num < 25")
	} else {
		fmt.Println("num >= 25")
	}
}
