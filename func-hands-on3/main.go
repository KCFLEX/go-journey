package main

import (
	"fmt"
)

type Person struct {
	first string
	age   int
}

func (p Person) Speak() {
	fmt.Printf("my name is %v and i am %v years old", p.first, p.age)
}

func main() {
	bro := Person{
		first: "Ikenna",
		age:   19,
	}
	bro.Speak()
	fmt.Printf("\nI love my brother %v with all my heart", bro.first)
}

/*● Create a user defined struct with
○ the identifier “person”
○ the fields:
■ first ■ last ■ age
● attach a method to type person with
○ the identifier “speak”
○ the method should have the person say their name and age
● create a value of type person
● call the method from the value of type person*/
