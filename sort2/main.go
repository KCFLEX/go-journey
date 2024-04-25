package main

import (
	"fmt"
	"sort"
)

// we are trying to sort a slice
func main() {
	xi := []int{4, 7, 3, 42, 99, 18, 16, 56, 12}
	xs := []string{"James", "Q", "M", "Moneypenny", "Dr. No"}
	sort.Ints(xi)
	fmt.Println(xi)
	sort.Strings(xs)
	fmt.Println(xs)

}
