package main

import (
	"fmt"
)

func main() {
	foo()

	// the func below os an anonymous function.  a nameless func

	func() {
		fmt.Println("anonymous func")
	}()

	d := func(b int) int {
		a := b * b
		return a
	}

	fmt.Println(d(0))

	fmt.Println("--------------------")
	c := func() {
		z := 6
		for i := 0; i < 13; i++ {
			h := i * z
			fmt.Printf("%v * %v = %v\n", i, z, h)
		}
	}
	c()
	fmt.Println("--------------------")
	// bar function call
	u := bar(5)
	g := u(10000)
	fmt.Println(g)
	fmt.Printf("%T\n", d)
	fmt.Printf("%T\n", c)
	fmt.Printf("%T\n", foo)
	fmt.Printf("%T\n", bar)

}

func foo() {
	fmt.Println("say food")
}

// this it a function that returns a function
func bar(t int) func(i int) int {
	return func(i int) int {
		return t * i
	}
}

// this is a named function with an identifier
// func (r reciever) identifier(p parameters) (r returns) {
//code to excute
//}

// an anonymous function
// func(p parameters) (r returns) {
// code to excute
//}

// note function are first class citizens like strings, integers, arrays, slice, maps, and so on an could also be returned as well
// a callback is passing in a function in as an  argument
