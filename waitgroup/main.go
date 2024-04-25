package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("OS\t", runtime.GOOS)
	fmt.Println("ARCH\t", runtime.GOARCH)
	fmt.Println("CPUs\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())
	wg.Add(1) // to add one concurent function
	go boo()
	coo()
	fmt.Println("\nCPUs\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())
	wg.Wait() // this waits until all the things we have added are done
}

func boo() {
	for i := 0; i < 10; i++ {
		fmt.Printf("boo: %v  ", i)
	}
	wg.Done()
}

func coo() {
	for i := 0; i < 10; i++ {
		fmt.Printf("\n coo: %v", i)
	}
}
