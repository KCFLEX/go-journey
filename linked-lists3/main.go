package main

import "fmt"

type node struct {
	value int   // to store the value of the node
	next  *node // to store the reference to the next node
}

type linkedList struct { // this representas the linked list itself
	head *node // to store the reference to the first node
}

func (l *linkedList) Insert(value int) {
	newNode := &node{value: value, next: nil}
	if l.head == nil { // this is saying if the head is empty
		l.head = newNode // set head to a new node
	} else {
		lastNode := l.head // instantiating lastnode with linkedlist
		for lastNode.next != nil {
			lastNode = lastNode.next
		}
		lastNode.next = newNode
	}
}

// display method
func (l *linkedList) Display() {
	current := l.head
	for current != nil {
		fmt.Printf("%d ", current.value)
		current = current.next
	}
}

func main() {
	Food := linkedList{} // this is an instance of the linked list
	Food.Insert(20)      // to insert values into the linked list
	Food.Insert(30)      // to insert values into the linked list
	Food.Insert(40)      // to insert values into the linked list
	Food.Insert(50)      // to insert values into the linked list
	Food.Insert(60)      // to insert values into the linked list
	Food.Insert(70)      // to insert values into the linked list
	// call the display method on the instance
	Food.Display()
}
