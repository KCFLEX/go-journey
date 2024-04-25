package main

import (
	"fmt"
	"math"
)

// type Square
type Square struct {
	length float64
	width  float64
}

// type Circle
type Circle struct {
	radius float64
}

// Area method
func (s Square) area() float64 {
	squareArea := s.length * s.width
	return squareArea
}

// Area method
func (c Circle) area() float64 {
	circleArea := math.Pi * math.Pow(c.radius, 2)
	return circleArea
}

// Shape interface
type Shape interface {
	area() float64 // please always remember to put the type
}

func info(sh Shape) float64 {
	return sh.area()
}

func main() {
	sh1 := Square{
		length: 3,
		width:  5,
	}
	sh2 := Circle{
		radius: 5,
	}

	value1 := info(sh1)
	value2 := info(sh2)
	fmt.Printf("The area of the square is = %v while the are of the circle is = %v", value1, value2)
}

/*● create a type SQUARE
● create a type CIRCLE
● attach a method to each that calculates AREA and returns it
○ circle area= π r 2
○ squarearea=L*W
● create a type SHAPE that defines an interface as anything that has the AREA method
● create a func INFO which takes type shape and then prints the area
● create a value of type square
● create a value of type circle
● use func info to print the area of square
● use func info to print the area of circle*/
