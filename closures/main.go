package main

import "fmt"

func main() {
	g := decrementor()
	fmt.Println(g())
	fmt.Println(g())
	fmt.Println(g())
	fmt.Println(g())
	fmt.Println(g())
	fmt.Println(g())
	fmt.Println(g())
	fmt.Println(g())
	fmt.Println(g())
	fmt.Println(g())
	fmt.Println(g())
	fmt.Println(g())
	fmt.Println(g())
	fmt.Println(g())
	fmt.Println(g())
	fmt.Println(g())
	fmt.Println(g())
	fmt.Println(g())
	fmt.Println(g())
	fmt.Println(g())
	fmt.Println(g())
	fmt.Println(g())
	fmt.Println(g())
	fmt.Println(g())
	fmt.Println(g())
	fmt.Println(g())
	fmt.Println(g())
	fmt.Println(g())
	fmt.Println(g())
	fmt.Println(g())
	fmt.Println(g())
	fmt.Println(g())
	fmt.Println(g())
	fmt.Println(g())
	fmt.Println(g())
	fmt.Println(g())
	fmt.Println(g())
	fmt.Println(g())
	fmt.Println(g())

	h := decrementor()
	fmt.Println(h())
	fmt.Println(h())
	fmt.Println(h())
	fmt.Println(h())
	fmt.Println(h())
	fmt.Println(h())
	fmt.Println(h())
	fmt.Println(h())
	fmt.Println(h())
	fmt.Println(h())
	fmt.Println(h())
	fmt.Println(h())
	fmt.Println(h())
	fmt.Println(h())
	fmt.Println(h())
	fmt.Println(h())
	fmt.Println(h())
	fmt.Println(h())
	fmt.Println(h())
	fmt.Println(h())
	fmt.Println(h())
	fmt.Println(h())
	fmt.Println(h())
	fmt.Println(h())
	fmt.Println(h())
	fmt.Println(h())
	fmt.Println(h())
	fmt.Println(h())
	fmt.Println(h())
	fmt.Println(h())
	fmt.Println(h())
	fmt.Println(h())
	fmt.Println(h())
	fmt.Println(h())
	fmt.Println(h())
	fmt.Println(h())
	fmt.Println(h())
	fmt.Println(h())
	fmt.Println(h())
	fmt.Println(h())
	fmt.Println(h())
	fmt.Println(h())
	fmt.Println(h())
	fmt.Println(h())
	fmt.Println(h())
	fmt.Println(h())
	fmt.Println(h())
	fmt.Println(h())

}

// what you see below is a closure
func decrementor() func() int {
	x := 100
	return func() int {
		x--
		return x
	}
}

// closure: is when an outer function encloses an inner function
