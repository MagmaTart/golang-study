package main

import "fmt"

func unicode_test() {
	var str string = "가나다"
	for idx, ch := range str {
		fmt.Println(idx, ch)
	}
}

func str_print_test() {
	var str string = "스트링"
	for idx, ch := range str {
		fmt.Println(idx, string(ch))
	}
}

func str_loop_test() {
	var str string = "퇴근"
	for i := 0; i < len(str); i++ {
		fmt.Printf("%d : %x\n", i, str[i])
	}
}

func str_modify_test() {
	var str string = "퇴근"
	b_str := []byte(str)
	b_str[2] += 5
	fmt.Println(string(b_str))
}

func str_concat_test() {
	str1 := "테"
	str2 := "스트"
	fmt.Println(str1 + str2)
	str1 += str2
	fmt.Println(str1)
}

func main() {
	unicode_test()
	str_print_test()
	str_loop_test()
	str_modify_test()
	str_concat_test()
}
