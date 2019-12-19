package main

import "fmt"

type IFileManager interface {
	Read() string
	Write(data string) error
}

type CSVFileManager struct {
	filename string
}

type TXTFileManager struct {
	filename string
}

type JSONFileManager struct {
	filename string
}

func (fm TXTFileManager) Read() string {
	return "Dummy String TXT"
}

func (fm TXTFileManager) Write(data string) error {
	var err error = nil
	fmt.Println("WRITE :", data)
	return err
}

func (fm CSVFileManager) Read() string {
	return "Dummy String CSV"
}

func (fm CSVFileManager) Write(data string) error {
	var err error = nil
	fmt.Println("WRITE :", data)
	return err
}

func (fm JSONFileManager) Read() string {
	return "Dummy String JSON"
}

func FileManagerTest(manager IFileManager) {
	fmt.Println(manager.Read())
	manager.Write("Write test")
}

func main() {
	var csv = CSVFileManager{"Dummy_Filename1"}
	var txt = TXTFileManager{"Dummy_Filename2"}
	var json = JSONFileManager{"Dummy_Filename3"}
	fmt.Println(csv.Read())
	fmt.Println(txt.Read())
	fmt.Println(json.Read())
	csv.Write("TEST_CSV")
	txt.Write("TEST_TXT")

	FileManagerTest(csv)
	FileManagerTest(txt)
	FileManagerTest(json)
}
