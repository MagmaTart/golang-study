package main

import "fmt"

func switch_test1() {
	num := 2
	switch num {
	case 1:
		fmt.Println("num == 1")
	case 2, 3:
		fmt.Println("num == 2 or num == 3")
	default:
		fmt.Println("num > 3")
	}
}

func switch_test2() {
	num := 2
	switch mul := num * 2; {
	case mul < 4:
		fmt.Println("mul < 4")
	case mul >= 4 && mul < 6:
		fmt.Println("4 <= mul < 6")
	default:
		fmt.Println("mul >= 6")
	}
}

func switch_test3() {
	num := 10
	switch {
	case num < 10:
		fmt.Println("num < 10")
	case num >= 10 && num < 20:
		fmt.Println("10 <= num < 20")
	default:
		fmt.Println("num > 20")
	}
}

func switch_test4() {
	num := 1
	switch num {
	case 1:
		fmt.Println("num : 1")
		fallthrough
	case 2:
		fmt.Println("num : 2")
		fallthrough
	default:
		fmt.Println("num : 3")
	}
}

func switch_test5(num interface{}) {	// 아직 interface가 뭔지 모르겠당
	switch num.(type) {
	case int:
		fmt.Println("num type is int")
	case float64:
		fmt.Println("num type is float64")
	default:
		fmt.Println("num type is unknown")
	}
}

func main() {
	switch_test1()
	switch_test2()
	switch_test3()
	switch_test4()
	switch_test5(9297)
}
