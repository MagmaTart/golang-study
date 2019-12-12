package main

import "fmt"

func print_slice(slc []int) {
	for i := 0; i < len(slc); i++ {
		fmt.Printf("%d ", slc[i])
	}
	fmt.Println()
}

func slice_by_value_test(slc []int) {
	fmt.Printf("Func 1 : ")
	print_slice(slc)
	slc[0] = 10
	slc = append(slc, 6)
	fmt.Printf("Func 2 : ")
	print_slice(slc)
}

func slice_by_pointer_test(slc *[]int) {
	fmt.Printf("Func 1 : ")
	print_slice(*slc)
	(*slc)[0] = 10
	*slc = append(*slc, 6)
	fmt.Printf("Func 2 : ")
	print_slice(*slc)
}

func add_and_mul(a, b int) (int, int) {
	return a + b, a * b
}

func named_add_and_mul(a, b int) (add int, mul int) {
	add = a + b
	mul = a * b
	return
}

func add_all_values_in_slice(slc... int) int {
	total := 0
	for i := 0; i < len(slc); i++ {
		total += slc[i]
	}
	return total
}

func main() {
	// slc := []int{1, 2, 3, 4, 5}
	// fmt.Printf("Main 1 : ")
	// print_slice(slc)
	// slice_by_value_test(slc)
	// slice_by_pointer_test(&slc)
	// fmt.Printf("Main 2 : ")
	// print_slice(slc)

	// add_result, mul_result := add_and_mul(30, 50)
	// add_result, mul_result := named_add_and_mul(30, 50)
	// fmt.Printf("Add : %d, Mul : %d\n", add_result, mul_result)

	// result := add_all_values_in_slice(5, 6, 7, 8, 9)
	// fmt.Printf("Add result : %d\n", result)

	slc := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	result := add_all_values_in_slice(slc...)
	fmt.Printf("Add result : %d\n", result)
}