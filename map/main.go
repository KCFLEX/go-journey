package main

import "fmt"


func main() {
	//create slice 
	b := make([]string, 0, 20)
	b = append(b, "20002")
	fmt.Println(b)
	// create map
	c := make(map[string][]string)
	c["bond_james"] = []string{"shaken and not stirred", "martinis", "fast cars"} 
	c["moneypenny_jenny"] = []string{"intelligence", "literature", "computer science"} 
	c["no_dr"] = []string{"cats", "ice-cream", "sunsets"} 
	c["fleming_ian"] = []string{"steak", "cigars", "espionage"} 
	
     for key, value := range c {
		fmt.Printf("%v\n", key)
		for i, v := range value {
			fmt.Printf("%v - %v\n", i, v)
		}
	 }
     delete(c, "fleming_ian")
	 for key, value := range c {
		fmt.Printf("%v\n", key)
		for i, v := range value {
			fmt.Printf("%v - %v\n", i, v)
		}
	 }

	 d := []string{
		"apple", "banana", "cherry", "date", "elderberry", "fig", "grape",
        "honeydew", "imbe", "jackfruit", "kiwi", "lemon", "mango", "nectarine",
        "orange", "papaya", "quince", "raspberry", "strawberry", "tangerine",
        "apple", "apple", "banana", "cherry", "date", "elderberry", "elderberry",
        "fig", "grape", "grape", "honeydew", "imbe", "jackfruit", "kiwi", "kiwi",
        "lemon", "mango", "mango", "nectarine", "orange", "papaya", "papaya",
        "quince", "raspberry", "strawberry", "tangerine", "tangerine",
        // Repeat the previous words as needed
        "apple", "banana", "cherry", "date", "elderberry", "fig", "grape",
        "honeydew", "imbe", "jackfruit", "kiwi", "lemon", "mango", "nectarine",
        "orange", "papaya", "quince", "raspberry", "strawberry", "tangerine",
        "apple", "apple", "banana", "cherry", "date", "elderberry", "elderberry",
        "fig", "grape", "grape", "honeydew", "imbe", "jackfruit", "kiwi", "kiwi",
        "lemon", "mango", "mango", "nectarine", "orange", "papaya", "papaya",
        "quince", "raspberry", "strawberry", "tangerine", "tangerine",
	 }
      // you ned to range over these words and show how many times each of them were repeated
       // first you need to create a map
	   e := make(map[string]int)
       // this block makes it that the number increments anytime the a word is repeated as it loops through
	   for _, v := range d {
		e[v]++
	   }

	   for i, v := range e {
		fmt.Println(i, v)
	   }
	
	}