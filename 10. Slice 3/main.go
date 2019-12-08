package main

import "fmt"

func print_slice(slice []int) {
	for i := 0; i < len(slice); i++ {
		fmt.Printf("%d ", slice[i])
	}
	fmt.Println()
}

func slice_element_insert_append() {
	slc := []int{1, 2, 3, 4, 5}
	print_slice(slc)
	idx := 3
	slc = append(slc[:idx+1], slc[idx:]...)
	slc[idx] = 6
	print_slice(slc)
}

func slice_element_insert_copy() {
	slc := []int{1, 2, 3, 4, 5}
	print_slice(slc)
	idx := 3
	slc = append(slc, 6)
	copy(slc[idx+1:], slc[idx:])
	slc[idx] = 6
	print_slice(slc)
}

func slice_slice_insert() {
	slc := []int{1, 2, 3, 4, 5}
	ttt := []int{7, 8, 9}
	print_slice(slc)
	idx := 3
	slc = append(slc, ttt...)
	copy(slc[idx+len(ttt):], slc[idx:])
	copy(slc[idx:], ttt)
	print_slice(slc)
}

func main() {
	slice_element_insert_append()
	slice_element_insert_copy()
	slice_slice_insert()
}