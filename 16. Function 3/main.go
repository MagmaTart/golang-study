package main

import "fmt"

func stack_and_print() func(string) {
	var string_stack []string
	return func(str string) {
		string_stack = append(string_stack, str)
		for i := 0; i < len(string_stack); i++ {
			fmt.Printf("%s ", string_stack[i])
		}
		fmt.Println()
	}
} 

func slice_iterator(slc []int) func() int {
	idx := 0
	return func() int {
		ret := slc[idx]
		idx++
		idx %= len(slc)
		return ret
	}
}

func generator() func() int {
	num := 0
	return func() int {
		num++
		return num
	}
}

func main() {
	// stacker := stack_and_print()
	// stacker("Google")
	// stacker("Apple")
	// stacker("Microsoft")

	// slc := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// next_element := slice_iterator(slc)
	// for i := 0; i < 20; i++ {
	// 	fmt.Println(next_element())
	// }

	next_num := generator()
	total := 0
	for i := 0; i < 100; i++ {
		total += next_num()
	}
	fmt.Println("Sum 1 ~ 100 :", total)
}
