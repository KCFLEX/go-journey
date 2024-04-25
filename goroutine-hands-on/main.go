package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func Sayhello(wg *sync.WaitGroup) {
	time.Sleep(2 * time.Second)
	fmt.Println("hello")
	wg.Done()
}

func Sayhi(wg *sync.WaitGroup) {
	time.Sleep(1 * time.Second)
	fmt.Println("hi")
	wg.Done()
}

func main() {
	fmt.Println("nice to meet you")
	wg.Add(2)
	go Sayhello(&wg)
	go Sayhi(&wg)
	wg.Wait()
}
