package main

import "fmt"


func main() {
	xi := []int{42, 43, 44, 45, 46, 47}

	for i, v := range xi {
		fmt.Printf("index: %v\t value: %v\n", i, v)
	}


	m := map[string]int{ 
		    "James": 42,
             "Moneypenny": 32,
		 }

		// for  i, v := range m {
		//	fmt.Printf("name: %v\n age: %v\n", i, v)
		// }

		 if v, ok := m["James"]; ok {
			fmt.Printf("age: %v\n", v)
		 }
}