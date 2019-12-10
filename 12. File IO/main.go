package main

import "fmt"
import "os"
import "bufio"

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

func file_write_test() error {
	filename := "new.txt"
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	if _, err := fmt.Fprintln(file, "File writing test!"); err != nil {
		// Fprintln에서 에러 발생시
		return err
	}
	return err
}

func file_bufio_read_test() error {
	filename := "test.txt"
	file, err := os.Open(filename)
	if err != nil {		// 에러 발생시
		return err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	if read_str, err := reader.ReadString('\n'); err == nil {
		fmt.Println(read_str)
	} else {
		return err
	}
	return err
}

func file_bufio_write_test() error {
	filename := "new.txt"
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	if _, err := writer.WriteString("bufio writing test!\n"); err != nil {
		return err
	}
	writer.Flush()
	return err
}

func main() {
	// stdout_test()
	// stdin_add_test()
	// sscanf_test()
	// err := file_read_test()
	// err := file_write_test()
	// err := file_bufio_read_test()
	err := file_bufio_write_test()
	fmt.Println(err)
}