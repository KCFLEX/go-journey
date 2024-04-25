package main

import "fmt"

func main() {
	even := make(chan int)
	odd := make(chan int)
	quit := make(chan bool)

	go func() {
		for i := 0; i < 100; i++ {
			if i%2 == 0 {
				even <- i
			} else {
				odd <- i
			}
		}
		close(even)
		close(odd)
		quit <- true
	}()

	for {
		select {
		case gg := <-even:
			fmt.Println("these numbers are even:", gg)
		case ig := <-odd:
			fmt.Println("these numbers are odd:", ig)
		case <-quit:
			return
		}
	}

}
