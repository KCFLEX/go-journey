package main

import "fmt"

func main() {
	c := make(chan int)

	// send
	go func() {
		for i := 0; i < 5; i++ {
			c <- i // sending values
		}
		close(c)
	}()

	// recieve
	for v := range c {
		fmt.Println(v)
	}
	fmt.Println("about to exit")
}

//ease of programming, channels block, buffered channels, range loop
