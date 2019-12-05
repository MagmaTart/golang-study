package main

import "fmt"

func slice_structure_test() {
	slc := []int{1, 2, 3}
	fmt.Println(len(slc), cap(slc))
}

func slice_len_cap_test() {
	slc := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Printf("slc --- len : %d, cap : %d\n", len(slc), cap(slc))
	slc = slc[2:6]
	fmt.Printf("slc --- len : %d, cap : %d\n", len(slc), cap(slc))
	slc = append(slc, 8)
	fmt.Printf("slc --- len : %d, cap : %d\n", len(slc), cap(slc))
	slc = slc[2:5]
	fmt.Printf("slc --- len : %d, cap : %d\n", len(slc), cap(slc))
}

func main() {
	// slice_structure_test()
	slice_len_cap_test()
}
