package main

import "fmt"

type node struct {
	data int
	next *node
}

type linkedList struct {
	head   *node
	length int
}

func (l *linkedList) prepend(n *node) {
	second := l.head // this is to keep reference to the second node incase if we update the the first node
	l.head = n       // this updates the value of l.head to the new node n *node
	// by doing this we are effectively updating the head field
	// this embraces n as the new starting point
	l.head.next = second // we are updating the next field of the first node to the second node
	l.length++
}

func (l linkedList) printListData() {
	// this is to print the list data
	toprint := l.head
	for l.length != 0 {
		fmt.Printf("%d", toprint)
		toprint = toprint.next // to update the toprint to the next
		l.length--             // decrease length by 1
	}
	fmt.Println("\n")
}
func main() {
	myList := linkedList{}
	node1 := &node{data: 42} // & is used to obtain the memory address of the node struct.
	node2 := &node{data: 43}
	node3 := &node{data: 44}
	node4 := &node{data: 45}
	node5 := &node{data: 46}
	node6 := &node{data: 47}
	myList.prepend(node1)
	myList.prepend(node2)
	myList.prepend(node3)
	myList.prepend(node4)
	myList.prepend(node5)
	myList.prepend(node6)
	fmt.Println(myList)
	myList.printListData()
}

// note there are differences between a pointer reciever and a value reciever
// pointer reciever is denoted like this (p *pointer) this is for when you would like to make changes to the reciever
// value reciever is denoted like this (r reciever) this is for when there is no need to make changes to the reciever
