package main

import (
	"fmt"
	"sort"
)

// we are sorting a slice of structs
// go to line 56 to see the slice of structs
type Person struct {
	First string
	Age   int
}
type ByAge []Person

// the methods below are implemented
// to sort a slice of structs  which type ByAge is
// type ByAge is a slice of struct
// we are doing this to sort by age
func (b ByAge) Len() int {
	return len(b)
}

func (b ByAge) Swap(i int, j int) {
	b[i], b[j] = b[j], b[i]
}

func (b ByAge) Less(i int, j int) bool {
	return b[i].Age < b[j].Age
}

// now we want to sort by name
// we just need to repeat the same process but change
// the body of the less method
// from return b[i].Age < b[j].Age to  return b[i].First < b[j].First
// this shows that we are no more targeting the Age field but the First field
type ByName []Person

func (bn ByName) Len() int { // this checks for the length of the slice
	return len(bn)
}

// find the explanation of the concept of swap below on line 63
func (bn ByName) Swap(i int, j int) { // this swaps values
	bn[i], bn[j] = bn[j], bn[i]
}

func (bn ByName) Less(i int, j int) bool { // this is where it does the comparison
	return bn[i].First < bn[j].First
}

// note only when the Sort has these methods above can it function
func main() {
	p1 := Person{"James", 32}
	p2 := Person{"Moneypenny", 27}
	p3 := Person{"Q", 64}
	p4 := Person{"M", 56}

	people := []Person{p1, p2, p3, p4} // this is a slice of structs
	sort.Sort(ByAge(people))
	fmt.Println(people)
	fmt.Println("----------------")
	sort.Sort(ByName(people))
	fmt.Println(people)
	// to explain the concept of swap
	x := 2
	y := 3
	fmt.Println(x, y)
	x, y = y, x       // by doing this we are swapping the values of x and y
	fmt.Println(x, y) // output 3, 2
}
