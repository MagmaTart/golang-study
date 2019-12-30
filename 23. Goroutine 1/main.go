package main

import "fmt"

func test() {
  for {
    fmt.Println("TEST FUNCTION")
  }
}

func main() {
  go test()
  for {
    fmt.Println("MAIN FUNCTION")
  }
}
