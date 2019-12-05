package main

import "fmt"

func slice_test1() {
	slice := make([]int, 3)
	slice[0] = 5
	slice[1] = 6
	slice[2] = 7
	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}
}

func slice_test2() {
	big := []int{1, 2, 3, 4, 5}
	small := big[1:4]
	for i := 0; i < len(big); i++ {
		fmt.Printf("%d ", big[i])
	}
	fmt.Println()
	for i := 0; i < len(small); i++ {
		fmt.Printf("%d ", small[i])
	}
}

func print_string_slice(slice []string) {
	for i := 0; i < len(slice); i++ {
		fmt.Printf("%s ", slice[i])
	}
	fmt.Println()
}

func slice_test3() {
	strs := []string{"AAA", "BBB", "CCC"}
	print_string_slice(strs)
	strs = append(strs, "DDD", "EEE")
	print_string_slice(strs)
}

func slice_test4() {
	strs := []string{"AAA", "BBB", "CCC"}
	new := []string{"FFF", "GGG"}
	print_string_slice(strs)
	strs = append(strs, new...)
	print_string_slice(strs)
}

func main() {
	slice_test1()
	slice_test2()
	slice_test3()
	slice_test4()
}
