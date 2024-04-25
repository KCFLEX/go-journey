package main

import "fmt"

func main() {
	c := make(chan<- int, 2) // send channel

	//c := make(<-chan int, 2) // recieve channel
	//c := make(chan int) // bi-directional channel
	c <- 42 // send value to channel c
	c <- 43

	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println("----------")
	fmt.Printf("%T\n", c)

	// this code is not working because the channel is locked

}
