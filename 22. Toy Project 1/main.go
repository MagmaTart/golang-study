package main

import "fmt"
import "github.com/nsf/termbox-go"

func main() {
  err := termbox.Init()
  if err != nil {
    panic(err)
  }
  defer termbox.Close()
  fmt.Println("Init complete")
}
