package main

import "fmt"

type Person struct {
	First string
	Last  string
	Age   int
}

func (p *Person) Speak() {
	fmt.Printf("Hi i am %v\n", p.First)
}

type Human interface {
	Speak()
}

func saySomething(hh Human) {
	hh.Speak()
}

func main() {
	person1 := Person{
		First: "Olga",
		Last:  "Olenik",
		Age:   23,
	}
	person1.Speak()
	saySomething(&person1)
}
