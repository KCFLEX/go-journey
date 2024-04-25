package main

import (
	"fmt"

	"github.com/KCFLEX/puppy"
)

/*const (
	_ = iota
	kb = 1 << (10 * iota)
	mb
	gb
	tb
	pb
	eb
)
*/
var cat int
var fish string = "tilapia"


func main() { 
	dog := 7

    lion, tiger := 4, 5 

	
	
   fmt.Printf(" %v\t %v\t %v\t %v\t %v\t\n ", cat, fish, dog, lion, tiger)
   
  s1, s2, s3 := puppy.Bark(), puppy.Barks(), puppy.BigBark2()

  fmt.Printf("The 1st dog started barking %v then the 2nd dog followed suit %v\n", s1, s2)
  fmt.Println(puppy.Bark())
  fmt.Println(puppy.Barks())
  fmt.Println(s3)
  fmt.Println(puppy.BigBarking())
  fmt.Println(puppy.Walk())

  

  

  
  
   

}

