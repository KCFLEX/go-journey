package main

import "fmt"

func main() {
	c := make(chan int, 2)

	c <- 42 // send value to channel c
	c <- 43

	fmt.Println(<-c)
	fmt.Println(<-c)

	// this code is not working because the channel is locked

}
