package main

import (
    "fmt"
)

// Shape interface
type Shape interface {
    Area() float64
}

// Rectangle struct
type Rectangle struct {
    Width  float64
    Height float64
}

// Circle struct
type Circle struct {
    Radius float64
}

// Area method for Rectangle
func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

// Area method for Circle
func (c Circle) Area() float64 {
    return 3.14 * c.Radius * c.Radius
}

func main() {
    // Create instances of Rectangle and Circle
    rect := Rectangle{Width: 10, Height: 5}
    circle := Circle{Radius: 7}

    // Print the area of each shape
    fmt.Printf("Area of Rectangle: %.2f\n", rect.Area())
    fmt.Printf("Area of Circle: %.2f\n", circle.Area())
}