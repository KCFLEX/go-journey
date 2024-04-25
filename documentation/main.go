package main

import (
	"example/go-journey/documentation/dog"
	"fmt"
)

type Canine struct {
	Name string
	Age  int
}

func main() {
	Scott := Canine{
		Name: "Scott",
		Age:  dog.Years(1),
	}
	fmt.Println(Scott)
}
