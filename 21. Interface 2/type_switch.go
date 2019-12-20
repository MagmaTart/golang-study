package main

import "fmt"

type Rect struct {
  width, height float64
}

type Circle struct {
  radius float64
}

func area(shape interface{}) float64 {
  area := 0.0
  switch shape.(type) {
  case Rect:
    rect := shape.(Rect)
    area = rect.width * rect.height
  case Circle:
    circle := shape.(Circle)
    area = circle.radius * circle.radius * 3.141592
  }
  return area
}

func main() {
  var rect = Rect{5.0, 6.0}
  var circle = Circle{3.0}

  fmt.Println("Rect area :", area(rect))
  fmt.Println("Circle area :", area(circle))
}
