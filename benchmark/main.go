package main

import (
	"example/go-journey/benchmark/saying"
	"fmt"
)

func main() {
	p := saying.Greet(" Olga")

	fmt.Println(p)
}
