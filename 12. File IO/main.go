package main

import "fmt"
import "os"

func stdout_test() {
	fmt.Fprintln(os.Stdout, "STDOUT TEST")
}

func stdin_add_test() {
	a, b := 0, 0
	fmt.Fscanf(os.Stdin, "%d", &a)
	fmt.Fscanf(os.Stdin, "%d", &b)
	fmt.Fprintf(os.Stdout, "%d + %d = %d\n", a, b, a+b)	
}

func sscanf_test() {
	a, b := 0, 0
	str := "10 + 20"
	fmt.Sscanf(str, "%d + %d", &a, &b)
	fmt.Printf("%d + %d = %d\n", a, b, a+b)
}

func file_read_test() error {
	filename := "test.txt"
	file, err := os.Open(filename)
	if err != nil {		// 에러 발생시
		return err
	}
	defer file.Close()
	var read_str string
	if _, err = fmt.Fscanf(file, "%s", &read_str); err == nil {
		// Fscanf에서 에러 미 발생시
		fmt.Println(read_str)
	}
	return err
}

func main() {
	// stdout_test()
	// stdin_add_test()
	// sscanf_test()
	err := file_read_test()
	fmt.Println(err)
}