package main

import (
	"fmt"
)

//func reverseSlice(s []int) {
//	left, right := 0, len(s)-1

//	for left < right {
//		s[left], s[right] = s[right], s[left]

//		left++
//		right--
//	}

//}

func main() {

	slice := []int{2, 7, 11, 15}
	ffm := twoSum(slice, 9)
	fmt.Println(ffm)

	//reverseSlice(slice)

	//fmt.Println("Reversed slice:", slice)

}

/*
func twoSum(nums []int, target int) []int {
	left, right := 0, len(nums)-1
	var newarr []int
	for left < right {
		sum := nums[left] + nums[right]
		if sum == target {
			newarr = append(newarr, left, right)
			break
		} else if sum < target {
			left++
		} else if sum > target {
			right--
		}
	}
	return newarr
}
*/
/*
func twoSum(nums []int, target int) []int {

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}
*/
func twoSum(nums []int, target int) []int {
	left, right := 0, len(nums)-1

	for left < right {
		sum := nums[left] + nums[right]
		if sum < target {
			left += 1
		} else if sum > target {
			right -= 1
		} else {
			return []int{left, right}
		}

	}
	return []int{}
}
