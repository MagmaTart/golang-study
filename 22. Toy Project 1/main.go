package main

import "github.com/nsf/termbox-go"

func main() {
  // Termbox testing codes
  err := termbox.Init()
  if err != nil {
    panic(err)
  }
  defer termbox.Close()
  termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
  termbox.SetCell(10, 10, rune('A'), termbox.ColorBlack, termbox.ColorWhite)
  termbox.Flush()
  for {}
}
