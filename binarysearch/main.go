package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}
	r := binarySearch(arr, 15)
	fmt.Println(r)
}

func binarySearch(arr []int, target int) int {
	left, right := 0, len(arr)-1

	for left <= right {
		fmt.Println("running")
		mid := (left + right) / 2

		if target == arr[mid] {
			return arr[mid]
		} else if target < arr[mid] {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}
