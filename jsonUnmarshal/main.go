package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	type Person struct { // this is the structure we want the data to be in
		First string `json:"First"`
		Last  string `json:"Last"`
		Age   int    `json:"Age"`
	}
	jsonformat := `[{"First":"Ade","Last":"Black","Age":21},{"First":"Bello","Last":"Nmadu","Age":21}]`
	converttobyte := []byte(jsonformat) // convert json to slice of bytes []byte
	//fmt.Println(jsonformat)
	//fmt.Println(converttobyte)

	var people []Person // we need to create a place to store the new data which is what we did here
	/*json.Unmarshal() function in Go is used for this purpose.
	It takes a JSON-encoded byte array or string as input and
	 a pointer to a Go value as output. */
	err := json.Unmarshal(converttobyte, &people) // we unmarshal the data and point to the    people []Person  address to store the data
	//                                             // people []Person is the go value we are pointing to
	if err != nil {
		log.Fatal("error")
	}
	fmt.Println("data", people)
	for i, v := range people {
		fmt.Printf("\n Person no:%v", i)
		fmt.Printf("  %v, %v, %v", v.First, v.Last, v.Age)
	}
}
