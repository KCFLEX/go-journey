package main

import (
	"fmt"
	"runtime"
)

func main() {
	c := make(chan int)

	for i := 0; i < 10; i++ {
		go func() {
			for i := 0; i < 10; i++ {
				c <- i
			}
		}()
	}

	fmt.Println("ROUTINES", runtime.NumGoroutine())

	for k := 0; k < 100; k++ {
		fmt.Println(k, <-c)
	}
	close(c)

}
