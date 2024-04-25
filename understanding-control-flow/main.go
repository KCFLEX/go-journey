package main

import (
	"fmt"
	"math"
	"math/rand"
)

 func main() {
	// SEQUENCE
	fmt.Println("this is the first statement to run")
	fmt.Println("this is the second statement to run")
	x := 40 //this is the third statement to run
	y := 5  //this is the fourth statement to run
	fmt.Printf("x = %v \n y = %v\n", x, y)

	//CONDITONAL 
	//If statements
	// switch statements
     
	if (x <= 22) {
		fmt.Println("condition  executed")
	} else {
		fmt.Println("condition not  executed")
	}
   z := rand.Intn(x)

   fmt.Printf("z = %v\n", z)
   
   fmt.Print(rand.Intn(40), "\n")
   fmt.Print(rand.Intn(40), "\n")
   fmt.Print(rand.Intn(40), "\n")
   fmt.Print(rand.Intn(40), "\n")
   fmt.Print(rand.Intn(40), "\n")
   fmt.Print(rand.Intn(40), "\n")
   fmt.Print(rand.Intn(40), "\n")  




	fmt.Println("@@",
		pow(3, 2, 10),
		pow(3, 3, 28),
	)
    for i := 0; i < 100; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Printf("%v\n", i)
	}

 }

 func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}