package main

import "fmt"

func main() {
	fmt.Println(Sum(1, 2, 3, 4, 5, 6, 7, 8))
}
func Sum(ii ...int) int { // (ii ...int) this is called a variaadic parameter
	summ := 0 // to hold the sum of the values
	for _, v := range ii {
		summ += v
	}
	return summ
}
