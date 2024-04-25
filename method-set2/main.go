package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
}

func (c *circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

type shape interface { // this is a collection of method signatures or a store for storing methods
	area() float64 // note a method can't be stored if it was't created in the first place and in this case float64 is a return not a type
}

func info(s shape) {
	fmt.Println("area", s.area())
}

func main() {
	c := &circle{
		radius: 5,
	}
	info(c)
	//fmt.Println(c.area())

}
