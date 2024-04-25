package main

import "fmt"

var count int

// this node represents a binary search tree
type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

// insert will add a node to the tree
// the key should not already be in the tree
func (n *Node) Insert(k int) {
	//compare key with reciever
	if n.Key < k {
		//move  right
		if n.Right == nil {
			n.Right = &Node{Key: k}
		} else {
			n.Right.Insert(k)
		}
	} else if n.Key > k {
		//move left
		if n.Left == nil {
			n.Left = &Node{Key: k}
		} else {
			n.Left.Insert(k)
		}
	}
}

// search takes any key value
//and returns true if ther is a node with that value

func (n *Node) Search(k int) bool {
	count++
	if n == nil {
		return false
	}
	if n.Key < k {
		// move right
		return n.Right.Search(k)
	} else if n.Key > k {
		// move left
		return n.Left.Search(k)
	}
	return true

}

func main() {
	tree := &Node{Key: 100}
	fmt.Println(tree)
	tree.Insert(56)
	tree.Insert(20)
	tree.Insert(24)
	tree.Insert(70)
	tree.Insert(60)
	tree.Insert(78)
	tree.Insert(44)
	tree.Insert(200)
	fmt.Println(tree.Search(78))
	fmt.Println(count)
}
