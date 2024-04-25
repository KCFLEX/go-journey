package main

import "fmt"

func main() {

	// anonymous func
	func() {
		fmt.Println("Olga is a slut")
	}()

	// func literal
	g := func(c int) {
		for i := 0; i <= 12; i++ {
			d := c * i
			fmt.Printf("%v * %v = %v\n", c, i, d)
		}
	}

	g(4)

	fmt.Println("------------")
	f := add(6)
	t := f(4)
	fmt.Println(t)
	y := returnF()   // remember when you assign a function to a variable that variable becomes a function and the variable can be called like a function
	fmt.Println(y()) // y() is a variable that became a function and is called like a function
	fmt.Println("------------")
	h := opera(5, 6, multiply)
	fmt.Println(h)
	h = opera(55, 5, division)
	fmt.Println(h)
	fmt.Println("------------")
	bb := outerF()
	bb()
	fmt.Println("------------")
	d := recur(5)
	fmt.Println(d)
}

// func return
func add(t int) func(i int) int {
	return func(i int) int {
		return t + i
	}
}

func returnF() func() int {
	return func() int {
		return 42
	}
}

// callback
func multiply(a int, b int) int {
	return a * b
}

func division(a int, b int) int {
	return a / b
}

func opera(a int, b int, f func(int, int) int) int {
	return f(a, b)
}

// closure is when we enclosed the scope of a variable in some code block
// also can be known as enclosing one function in another function
// closures

func outerF() func() {
	word := "closure"

	innerF := func() {
		fmt.Println(word)
	}

	return innerF
}

// rcursive
func recur(num int) int {
	if num == 1 {
		return num
	}
	return num * recur(num-1)
}

// note: callback function is a function passed as an argument
// to be executed at a specific point or event.
// To summarize, a wrapper function wraps or modifies another function's behavior,
