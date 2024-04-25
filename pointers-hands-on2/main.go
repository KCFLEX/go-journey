package main

import "fmt"

type Person struct {
	first string
}

func (p Person) name1(k string) Person {
	p.first = k
	return p
}

func (p *Person) name2() {
	fmt.Println("my name is", p.first)
}

func main() {

	person2 := &Person{
		first: "Triston",
	}
	person2.name2()

	person1 := Person{}
	person1.name1("andrew")
	fmt.Println(person1)

}
