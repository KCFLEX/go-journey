package main

import (
	"fmt"
)

func main() {

	d := func() {
		h := 5
		for i := 5; i > 0; i-- {
			c := i * h
			fmt.Printf("%v * %v = %v\n", i, h, c)
		}
	}
	d()
	fmt.Println("--------")
	fmt.Println(f(7))
	fmt.Printf("%T\n", multiply)
	fmt.Printf("%T\n", divide)
	fmt.Printf("%T\n", solve)
	fmt.Println("--------")
	fmt.Println(solve(10, 2, multiply))
	fmt.Println(solve(15, 5, divide))
}

// recursive function
func f(num int) int {
	if num == 1 {
		return 1
	} else {
		var answer int = num * f(num-1)
		return answer
	}
}

// callback functions
// a callback is a function passed into another function as an argument

func multiply(a int, b int) int {
	return a * b
}
func divide(a int, b int) int {
	return a / b
}

func solve(a int, b int, f func(int, int) int) int {
	// here we called the function f and make it look like a function call
	return f(a, b)
}
