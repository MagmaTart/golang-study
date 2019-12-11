package main

import "fmt"
import "os"
import "bufio"

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

func file_bufio_scanner_test() error {
	filename := "test.txt"
	file, err := os.Open(filename)
	if err != nil {		// 에러 발생시
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	return err
}

func main() {
	// err := file_bufio_read_test()
	// err = file_bufio_write_test()
	err := file_bufio_scanner_test()
	fmt.Println(err)
}