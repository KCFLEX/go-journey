package main

import "fmt"

// recursion is when a function calls itself or when a function call is within itself
func main() {
	fmt.Println("-------")
	s := factorial(6)
	fmt.Println(s)
	d := factLoop(6)
	fmt.Println(d)
}

func factorial(num int) int {
	if num == 1 {
		return 1
	}
	return num * factorial(num-1)
}

// note: anything that could be done with recursion could be done with a loop
func factLoop(n int) int {
	total := 1
	for n > 0 {
		// this also means the current value of n is multiplied with the current value of total
		//and the result is stored in the total variable
		total *= n // this is also the same as total = total * n
		n--
	}
	return total
}
