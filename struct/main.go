package main

import "fmt"

// create struct
type person struct {
	first string
	last  string
	age   int
}

type animal struct {
	first  string
	specie string
	person
}

func main() {
	e := person{
		first: "Heath",
		last:  "legder",
		age:   29,
	}
	f := animal{

		first:  "dog",
		specie: "rott",
		person: person{
			first: "Heath",
			last:  "legder",
			age:   29,
		},
	}
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(f.first, f.person.first, e.first)

}
