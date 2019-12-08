package main

import "fmt"

func print_slice(slice []int) {
	for i := 0; i < len(slice); i++ {
		fmt.Printf("%d ", slice[i])
	}
	fmt.Println()
}

func slice_structure_test() {
	slc := []int{1, 2, 3}
	fmt.Println(len(slc), cap(slc))
}

func slice_make_test() {
	slc_a := make([]int, 7, 20)
	print_slice(slc_a)
	fmt.Printf("slc_a --- len : %d, cap : %d\n", len(slc_a), cap(slc_a))
	slc_b := make([]int, 15)
	print_slice(slc_b)
	fmt.Printf("slc_b --- len : %d, cap : %d\n", len(slc_b), cap(slc_b))
}

func slice_len_cap_test() {
	slc := []int{1, 2, 3, 4, 5, 6, 7, 8}
	print_slice(slc)
	fmt.Printf("slc --- len : %d, cap : %d\n", len(slc), cap(slc))
	slc = slc[3:]
	print_slice(slc)
	fmt.Printf("slc --- len : %d, cap : %d\n", len(slc), cap(slc))
	slc = append(slc, 8, 9)
	print_slice(slc)
	fmt.Printf("slc --- len : %d, cap : %d\n", len(slc), cap(slc))
	slc = slc[2:5]
	print_slice(slc)
	fmt.Printf("slc --- len : %d, cap : %d\n", len(slc), cap(slc))
}

func slice_copy_test() {
	slc := []int{1, 2, 3, 4, 5, 6, 7, 8}
	print_slice(slc)
	fmt.Printf("slc --- len : %d, cap : %d\n", len(slc), cap(slc))
	bbb := make([]int, len(slc[2:5]))
	fmt.Printf("bbb --- len : %d, cap : %d\n", len(bbb), cap(bbb))
	copy(bbb, slc[2:5])
	bbb[0] = 100
	bbb[1] = 200
	print_slice(bbb)
	print_slice(slc)
}

func main() {
	// slice_structure_test()
	slice_make_test()
	// slice_len_cap_test()
	// slice_copy_test()
}
