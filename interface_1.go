package main

import "fmt"

// define ineterface 
type Shape interface {
    Area() float64
}

// Define a struct: Rectangle
type Rectangle struct {
    Width  float64
    Height float64
}

// Implement the Area() 
func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

// Define another struct: Circle
type Circle struct {
    Radius float64
}

// Implement the Area() method for Circle
func (c Circle) Area() float64 {
    return 3.1416 * c.Radius * c.Radius
}


// print the calculated area
func printArea(s Shape) {
    fmt.Printf("Area: %.2f\n", s.Area())
}

func main() {
    rect := Rectangle{Width: 10, Height: 5}
    circle := Circle{Radius: 7}

    // Use the interface
    printArea(rect)
    printArea(circle)
}
