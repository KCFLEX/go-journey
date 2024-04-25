package main

import "fmt"

// generics
func addFloat(a float64, b float64) float64 {
	c := a + b
	return c
}

func addint(a int, b int) int {
	c := a + b
	return c
}

type mynum interface {
	~int | ~float64 // this (~) tells the interface to include all values of type int
}

func AddT[T mynum](a T, b T) T { //[T num] is a type constraint
	c := a + b
	return c
}

type Alias int

func main() {
	var n Alias = 42
	r := addFloat(2.2, 3.5)
	fmt.Println(r)
	p := addint(1, 2)
	fmt.Println(p)

	poo := AddT(2.2, 3.5)
	fmt.Println(poo)
	yoo := AddT(n, 2)
	fmt.Println(yoo)

}
