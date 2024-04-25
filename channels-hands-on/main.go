package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	c := gen()
	//fmt.Printf("%T and %v", c, c)

	wg.Add(1)
	go receive(c)

	fmt.Println("about to exit")
	wg.Wait()
}

func gen() <-chan int {
	c := make(chan int)
	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
		wg.Done()
	}()
	return c
}

func receive(c <-chan int) {
	for v := range c {
		fmt.Println("value no:", v)
	}
}
