package main

import "fmt"
import "reflect"

func print_map(map_to_print map[int]string) {
	for key, value := range(map_to_print) {
		fmt.Printf("%d : %s\n", key, value)
	}
}

func map_test() {
	mymap := map[int]string{1: "AAA", 2: "BBB", 3: "CCC"}
	// imap := map[int]int{1: 2, 2: 4, 3: 6}
	value, exists := mymap[2]
	fmt.Println(value, exists)
	value, exists = mymap[4]
	fmt.Println(value, exists)
	mymap[4] = "DDD"
	value, exists = mymap[4]
	fmt.Println(value, exists)
}

func map_delete_test() {
	mymap := map[int]string{1: "AAA", 2: "BBB", 3: "CCC", 4: "DDD"}
	print_map(mymap)
	fmt.Printf("Length : %d\n\n", len(mymap))
	delete(mymap, 3)
	print_map(mymap)
	fmt.Printf("Length : %d\n\n", len(mymap))
}

func map_equal_test() {
	mymap1 := map[int]string{1: "AAA", 2: "BBB", 3: "CCC", 4: "DDD"}
	mymap2 := map[int]string{1: "AAA", 2: "BBB", 3: "CCC", 4: "DDD"}
	mymap3 := map[int]string{1: "EEE", 2: "FFF", 3: "GGG", 4: "HHH"}

	fmt.Println("mymap1 is equal to mymap2? : ", reflect.DeepEqual(mymap1, mymap2))
	fmt.Println("mymap1 is equal to mymap3? : ", reflect.DeepEqual(mymap1, mymap3))
}

func main() {
	// map_test()
	// map_delete_test()
	map_equal_test()
}