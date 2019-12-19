package main

import "fmt"

type Color int
type Grade int

func main() {
	const CA = 100
	const CB int = 200
	const CS = "MagmaTart"

	const (
		CC = 1234
		CD = 5678
		CS2 = "Soomin"
	)

	const (
		_           = iota
		RED   Color = iota
		GREEN
		BLUE
	)

	fmt.Println(RED, GREEN, BLUE)

	const (
		A Grade = 100 - (10 * iota)
		B
		C
		D
		F
	)

	fmt.Println(A, B, C, D, F)
}
