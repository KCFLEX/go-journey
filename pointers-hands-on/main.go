package main

import "fmt"

var (
	a, b, c *string
	d       *int
)

func init() {
	p := "Drop by drop, the bucket gets filled."
	q := "Persistently, patiently, you are bound to succeed."
	r := "The meaning of life is ..."
	n := 42
	a = &p
	b = &q
	c = &r
	d = &n
}

func main() {
	fmt.Println(*a) // dereference of a 
	fmt.Printf("%v\t%T\n", a, a)
	fmt.Println(*b) // dereference of b 
	fmt.Printf("%v\t %T\n", b, b)
	fmt.Println(*c) // dereference of c 
	fmt.Printf("%v\t%T\n", c, c)
	fmt.Println(*d) // dereference of d
	fmt.Printf("%v\t%T\n", d, d)
}

