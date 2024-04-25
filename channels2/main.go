package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() { // main  thread
	var wg sync.WaitGroup
	data := make(chan int) // unbuffered channel
	/*
		go func() {
			data <- 45 // sending data to the chanel
		}()
		n := <-data// recieving data from the channel
		fmt.Println(n) // this gave a deadlock because the channel was not sending and recieving at the same time
		// for it to work we have to put it in a goroutine
	*/
	// another example below
	go func() { // background thread
		for i := 0; i < 1000; i++ {
			wg.Add(1)
			go func() {
				result := Dowork()
				data <- result
				defer wg.Done()
			}()
		}
		wg.Wait()
		close(data)
	}()
	for n := range data { // this is another syntax you can use to grab values out of the channel
		fmt.Println(n)
	}
}

func Dowork() int {
	time.Sleep(time.Second)
	return rand.Intn(100)

}
