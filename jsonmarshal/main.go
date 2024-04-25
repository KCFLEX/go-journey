package main

import (
	"encoding/json"
	"fmt"
	"log"
)

//func Marshal(v any) ([]byte, error) => to prove that json.marshal throws a slice of bytes and an error

/*In Go, a byte is a built-in data type that represents
an 8-bit unsigned integer. It is commonly used to store and
manipulate binary data or ASCII characters.
The byte type is an alias for the `uint8` type,
which means it can hold values from 0 to 255.
In Go, bytes are often used when working with
low-level I/O operations, network protocols, or
when dealing with raw data that needs to be manipulated
at the byte level.*/

type Person struct {
	First string
	Last  string
	Age   int
}

func main() {

	person1 := Person{
		First: "Ade", // always make these fields uppercase
		Last:  "Black",
		Age:   21,
	}

	person2 := Person{
		First: "Bello",
		Last:  "Namadu",
		Age:   21,
	}

	civilians := []Person{person1, person2} // this is a slice of structs because Person is a struct
	fmt.Println(civilians)
	fmt.Println("-------------")
	b, err := json.Marshal(civilians) // json.marshal throws a slice of byte and an error
	if err != nil {                   // error handling
		log.Fatal("error")
	}
	fmt.Println(string(b)) // theres a need to convert b to a string
	//because json.marshal throws a slice of bytes so b is a slice of bytes(uint8)

}
