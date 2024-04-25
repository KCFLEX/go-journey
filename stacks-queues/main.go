package main

import "fmt"

// stack represents stack that hold a slice
type Stack struct {
	elements []int
}

// push
func (s *Stack) Push(i int) {
	s.elements = append(s.elements, i)

}

// pop will remove a value at the end
// and return the removed value
func (s *Stack) Pop() int {
	lastIndex := len(s.elements) - 1 // this is denoting the last index
	popped := s.elements[lastIndex]
	s.elements = s.elements[:lastIndex]
	return popped
}

type Queue struct {
	elements []int
}

// Enqueue
func (q *Queue) Enqueue(i int) {
	q.elements = append(q.elements, i)
}

// Dequeue removes the value at the front
// and returns the removed value
func (q *Queue) Dequeue() int {
	firstIndex := q.elements[0]
	q.elements = q.elements[1:]
	return firstIndex
}

func main() {
	// what we have here is a stack data structure
	boo := Stack{}
	fmt.Println(boo)
	boo.Push(200)
	boo.Push(300)
	boo.Push(400)
	boo.Push(500)
	fmt.Println(boo)
	fmt.Println("------------------")
	fmt.Println(boo.Pop())
	fmt.Println(boo)
	fmt.Println("------------------")
	foo := Queue{}
	fmt.Println(foo)
	foo.Enqueue(10)
	foo.Enqueue(20)
	foo.Enqueue(30)
	foo.Enqueue(40)
	fmt.Println(foo)
	foo.Dequeue()
	fmt.Println(foo)

	fmt.Println("------------------")
	numbers := []int{1, 2, 3, 4, 5}
	newSlice := numbers[1:]
	fmt.Println(newSlice)
	numbers[2] = 9
	fmt.Println(numbers)
	newSlice[2] = 9
	fmt.Println(newSlice)

}
