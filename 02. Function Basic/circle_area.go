package main

import "fmt"

func circle_area(radius float64) float64 {
	pi := 3.141592
	return pi * radius * radius
}

func main() {
	radius := 7.0
	area := circle_area(radius)
	fmt.Println(area)
}
