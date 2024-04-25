package main

import "fmt"

func main() {
	c := make(chan int)

	c <- 42 // send value to channel c
	fmt.Println(<-c)
	// this code is not working because the channel is locked
}
