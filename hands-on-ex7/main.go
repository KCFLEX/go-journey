package main

import (
	"fmt"
	"math/rand"
)

func main() {
	for i := 0; i < 100; i++ {
	if x := rand.Intn(5); x == 3 {
      fmt.Printf("iteration no %v  x is %v\n", i, x)
	}
}

fmt.Println(true && true)
fmt.Println("-------------")
fmt.Println(true && false)
fmt.Println("-------------")
fmt.Println(true || true)
fmt.Println("-------------")
fmt.Println(true || false)
fmt.Println("-------------")
fmt.Println(!true)



}