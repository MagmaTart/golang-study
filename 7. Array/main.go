package main

import "fmt"

func array_test1() {
	const arr_len = 3
	arr_a := [...]int{1, 2, 3}
	arr_b := [arr_len]int{4, 5, 6}
	for i := 0; i < arr_len; i++ {
		fmt.Println(arr_a[i], arr_b[i])
	}
}

func array_test2() {
	arr := [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	for y := 0; y < len(arr); y++ {
		for x := 0; x < len(arr[y]); x++ {
			fmt.Printf("%d ", arr[y][x])
		}
		fmt.Printf("\n")
	}
}

func main() {
	array_test1()
	array_test2()
}
