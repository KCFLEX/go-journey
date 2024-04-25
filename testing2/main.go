package main

import "fmt"

func main() {
	fmt.Println(Sum(1, 2, 3, 4, 5, 6, 7, 8))
}

// Sum could take in an infinite number of parameters of type int
// and sums them all together
func Sum(ii ...int) int { // (ii ...int) this is called a variaadic parameter
	summ := 0 // to hold the sum of the values
	for _, v := range ii {
		summ += v
	}
	return summ
}
