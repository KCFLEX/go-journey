package main

import "fmt"

// input: [1 ,2 ,3], (n) => n * 2
// output: [2, 4, 6]

/*
	func mapValues(values []int, mapFunc func(int) int) []int {
		var newValues []int
		for _, v := range values {
			newValue := mapFunc(v)
			newValues = append(newValues, newValue)
		}
		return newValues
	}
*/
func mapValues(values []int) []int {
	var newValues []int

	for _, v := range values {
		newValue := v * 2
		newValues = append(newValues, newValue)
	}
	return newValues
}

func main() {
	d := mapValues([]int{1, 2, 3, 4})
	fmt.Println(d)

	//	result := mapValues([]int{1, 2, 3}, func(i int) int {
	//		return i * 2
	//	})

	// fmt.Printf("result: %v\n", result)
}
