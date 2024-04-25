package main

import "fmt"

//memory

func main() {
	x := 42
	fmt.Println(x)
	fmt.Println(&x)                // this gives us the address of x
	y := &x                        // now we are storing the address of x in variable y
	fmt.Printf("%v \t %T\n", y, y) // this would give us the address of x because the address of x was stored in y and it also gives us the type of the value stored in x
	fmt.Println(*y)                // this gives us the value stored iin the address that was stored in y which is 42
	*y = 43                        // we are trying to change the value in that address
	fmt.Println(x)                 // the value of x changes
	// this is because we pointed the address and changed its value
	fmt.Println(*y)
	// we can manipulate the value in the address by using *
	// pointers are also called reference types

	
}



// Dereferencing 'p' gives you the integer that 'p' points to // You can change the value that 'p' points to
// The value of 'i' is now 21
// * and & are operators while this *int is a pointer to an int
// & gives us the address while * gives us the value in that address
