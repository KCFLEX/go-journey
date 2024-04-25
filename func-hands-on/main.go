package main

import "fmt"

// basic funcs
func main() {
	b := foo()
	fmt.Println(b)

	c, d := bar()
	fmt.Println(c, d)
}

func foo() int {
	return 23
}
func bar() (int, string) {

	return 45, "John"
}
