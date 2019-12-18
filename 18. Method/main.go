package main

import "fmt"

type Length int

func (n Length) area() Length {
	return n * n
}

func main() {
	var length Length = 15
	fmt.Println(length.area())
}
