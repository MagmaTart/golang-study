package main

import "fmt"

func main() {
	pi := 3.141592
	radius := 3.0		// 3으로 할 경우 pi와 타입이 달라져 곱셈 불가능

	fmt.Println("Area of circle with radius", radius)
	fmt.Println(" --->", radius*radius*pi)
}
