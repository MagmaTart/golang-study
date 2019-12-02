package main

import "fmt"

func break_test() {
	for num := 0; num < 100; num++ {
		fmt.Println("num :", num)
		if num >= 50 {
			break
		}
	}
}

func continue_test() {
	for num := 0; num < 100; num++ {
		if num%2 == 0 {
			continue
		}
		fmt.Println("num :", num)
	}
}

func goto_test() {
	for num := 0; num < 100; num++ {
		fmt.Println("num :", num)
		if num >= 70 {
			goto END
		}
	}
END:
	fmt.Println("JUMPED END")
}

func break_label_test() {
LOOP_A:
	for a := 0; a < 100; a++ {
LOOP_B:
		for b := 0; b < 100; b++ {
			fmt.Printf("%d * %d = %d\n", a, b, a*b)
			if a >= 10 && b >= 50 {
				break LOOP_A
			} else if b >= 50 {
				break LOOP_B
			}
		}
	}
}

func main() {
	// break_test()
	// continue_test()
	// goto_test()
	break_label_test()
}
