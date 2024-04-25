package main

import "fmt"

const ArraySize = 7

// hashTable structure
type hashTable struct {
	array [ArraySize]*bucket
}

// methos for the hash table
// insert
// Search
// Delete
// bucket structure
// bucket structure : linked list
type bucket struct {
	head *Bucketnode
}

// Bucket node is a linked list that holds the key
type Bucketnode struct {
	Key  string
	next *Bucketnode
}

// insert
//func (hh *hashTable) Insert(key string) {
//    index := hash(key)
//}
// Search
//func (hh *hashTable) Search(key string) bool {
//   index := hash(key)
//}
// Delete
//func (hh *hashTable) Delete(key string) {//
//   index := hash(key)
//}

// hash function
func hash(Key string) int {
	// get the haskey code for each divider sum it up
	//and divide by the size of the array and get the remainder
	sum := 0
	for _, v := range Key { // this code loops through each character of the string and turns it into an integer and then sums it up together and adds it to sum
		sum += int(v)
	}
	return sum % ArraySize // to get the remainder of the arraysize
}

// init
func Init() *hashTable {
	result := &hashTable{}
	for i := range result.array {
		result.array[i] = &bucket{}
	}
	return result

}

func main() {
	testHashtable := Init()
	fmt.Println(testHashtable)
}
