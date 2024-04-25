package main

import (
	"fmt"
)

func main() {
/* for i := 0; i < 42; i++ {
	x := rand.Intn(5)
	fmt.Printf("i = %v\n", i)
    
  switch x {
   case 0:
	fmt.Printf("This is = %v\n", x)
   case 1:
	fmt.Printf("This is = %v\n", x)
   case 2:
	fmt.Printf("This is = %v\n", x)
   case 3:
	fmt.Printf("This is = %v\n", x)
   case 4: 
   fmt.Printf("This is = %v\n", x)
   default: 
   fmt.Printf("Out of Bounds")	
  }
}
y := 100
for y > 0 {
    fmt.Printf("y = %v\n", y)
	y--
}
*/
z := 100

for {
	
	
	if z <= 0 {
		break
	}
	fmt.Println(z)
	z--
}
  
}