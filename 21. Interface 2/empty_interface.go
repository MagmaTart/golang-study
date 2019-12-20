package main

import "fmt"

func getType(arg interface{}) {
  switch arg.(type) {
  case int:
    fmt.Println("Argument is Int")
  case float32, float64:
    fmt.Println("Argument is Float")
  case string:
    fmt.Println("Argument is String")
  }
}

func printTypeAndValue(arg interface{}) {
  switch value := arg.(type) {
  case int:
    fmt.Println("Int, value :", value)
  case float32, float64:
    fmt.Println("Float, value :", value)
  case string:
    fmt.Println("String, value :", value)
  }
}

func printString(arg interface{}) {
  str := arg.(string)
  fmt.Println(str)
}

func main() {
  i := 50
  f := 5.791
  s := "Hello"

  // getType(i)
  // getType(f)
  // getType(s)
  // printString(s)
  printTypeAndValue(i)
  printTypeAndValue(f)
  printTypeAndValue(s)
}
