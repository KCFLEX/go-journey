package main

import (
	"fmt"
	"log"
	"math"
)

//var mathError = errors.New("math error: square root of negative number")

func main() {
	//fmt.Printf("%T\n", mathError)
	fmt.Println("--------")
	r, err := sqrt(-10)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(r)
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		mathError2 := fmt.Errorf("math error: square root of negative number: %v", f)
		return 0, mathError2
	}
	c := math.Sqrt(f)
	return c, nil // you are returning an error you can use nil
}
