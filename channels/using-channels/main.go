package main

import "fmt"

func main() {
	c := make(chan int)

	// send
	go foo(c)

	// recieve
	doo(c)
	fmt.Println("about to exit")
}

func foo(d chan<- int) {
	d <- 42 // sending values
}

func doo(d <-chan int) { // recieving values
	fmt.Println("value recieved from channel:", <-d) // pulling the value off
}
