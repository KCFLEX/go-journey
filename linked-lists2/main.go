package main

import "fmt"

type Node struct {
	value int
	next  *Node
}

type LinkedList struct {
	head *Node
}

func (ll *LinkedList) Insert(value int) {
	newNode := &Node{value: value, next: nil}

	if ll.head == nil {
		ll.head = newNode
	} else {
		current := ll.head
		for current.next != nil {
			current = current.next
		}
		current.next = newNode
	}
}

func (ll *LinkedList) Display() {
	current := ll.head
	for current != nil {
		fmt.Printf("%d ", current.value)
		current = current.next
	}
	fmt.Println()
}

func main() {
	linkedList := LinkedList{}

	linkedList.Insert(10)
	linkedList.Insert(20)
	linkedList.Insert(30)
	linkedList.Insert(40)

	linkedList.Display() // Output: 10 20 30 40
}

/*Define the Node struct:
Create a new Go file, e.g., linkedlist.go.
Declare a struct named Node that represents a node in the linked list.
Inside the Node struct, define two fields: value of type int to store the value of the node, and next of type *Node to store the reference to the next node.
Define the LinkedList struct:
Declare another struct named LinkedList that represents the linked list itself.
Inside the LinkedList struct, define a field named head of type *Node to store the reference to the first node in the list.
Implement the Insert method:
Add a method named Insert to the LinkedList struct.
The Insert method should take a value as a parameter.
Inside the method, create a new Node with the provided value.
If the linked list is empty (i.e., head is nil), set the head to the new node.
Otherwise, traverse the list until you reach the last node and set the next field of the last node to the new node.
Implement the Display method:
Add a method named Display to the LinkedList struct.
Inside the method, initialize a variable current with the head of the list.
Start a loop that continues until current becomes nil.
In each iteration, print the value of the current node and update current to the next node.
Print a newline character after the loop to separate the output.
Create an instance of the LinkedList struct:
In the main function, create a new instance of the LinkedList struct using the LinkedList{} syntax.
Insert values into the linked list:
Call the Insert method on the linked list instance to insert values into the list.
You can insert values in any order or sequence you prefer.
Display the linked list:
Call the Display method on the linked list instance to print the values of all nodes in the list.
Run the program:
Save the file and run the Go program using the go run linkedlist.go command.
By following these step-by-step instructions, you should be able to build and run a basic linked list in Go. Feel free to refer to the previous examples and explanations for guidance.*/
