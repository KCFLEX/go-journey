// variadic funcs
package main

import "fmt"

func main() {
	defer fmt.Println("this sentence was deferred")
	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9} //slice of ints
	fmt.Println(foo(x...))
	defer fmt.Println("----")
	numbers := []int{10, 20, 30, 40, 50}
	fmt.Println(bar(numbers))
}

func foo(m ...int) int {
	total := 0
	for _, v := range m {
		total += v
	}
	return total

}

func bar(k []int) int {
	t := 0
	for _, v := range k {
		t += v
	}
	return t
}

/*● create a func with the identifier foo that
○ takes in a variadic parameter of type int
○ pass in a value of type []int into your func (unfurl the []int)
○ returns the sum of all values of type int passed in
● create a func with the identifier bar that
○ takes in a parameter of type []int
○ returns the sum of all values of type int passed in*/

/*Use the “defer” keyword to show that a deferred
func runs after the func containing it exits.*/
// deferred function run in last in first out order
