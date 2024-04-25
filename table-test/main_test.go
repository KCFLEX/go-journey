package main

import "testing"

/*
In Go, a table test is a testing technique
used to verify the behavior of a function or
method against multiple input values and
expected output values.
It involves creating a table of test cases,
where each row represents a specific input and
expected output combination.
*/
type table struct {
	Data   []int //slice
	answer int
}

func TestSum(t *testing.T) {
	testCases := []table{
		{
			Data:   []int{1, 2, 3},
			answer: 6,
		},
		{
			Data:   []int{4, 5, 6},
			answer: 15,
		},
		{
			Data:   []int{7, 8, 9},
			answer: 24,
		},
	}

	for _, v := range testCases {
		result := Sum(v.Data...)
		if result != v.answer {
			t.Errorf("result should be %v, instead it was %v", v.answer, result) // use t.Error to signal failure
		}
	}

}
