package main

import "fmt"

type dog struct {
	first string
}

func (d dog) walk() {
	fmt.Println("My name is", d.first, "and i am walking")
}

func (d dog) run() {

	fmt.Println("My name is", d.first, "and i am running")

}

//func (d dog) bark() {
//	fmt.Println(d.first, "barked woof woof")
//}

type youngin interface {
	walk()
	run()
	//bark()
}

func youngrun(yu youngin) {
	yu.walk()
	yu.run()
}

func main() {
	d1 := dog{"Henry"}
	youngrun(d1)
	d2 := dog{"paget"}
	fmt.Println("-----------")
	youngrun(d2)
	fmt.Println("-----------")
	d2 = d2.changeName("Ada")
	youngrun(d2)
}

func (d dog) changeName(g string) dog {
	d.first = g
	return d
}
