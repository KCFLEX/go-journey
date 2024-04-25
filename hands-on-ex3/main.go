package main

import (
	"fmt"
	"math/rand"
)

// the init func always runs before the main func and it takes no parameters

	



func main()  {

	//x = rand.Intn(300)
	//fmt.Printf("%T\t %v\n", x, x)
	for i := 0; i < 100; i++ {
	x := rand.Intn(10)
	fmt.Printf("this is x: %v\n", x)
	y := rand.Intn(10)
	fmt.Printf("this is y: %v\n", y)
	/*switch{
	case x < 100:
		fmt.Printf("less than 100: %v\n",x)
	case x > 101 && x < 200:
		fmt.Printf("greater than 100 and less than 200: %v\n",x)
    case x > 201 && x < 250:
		fmt.Printf("greater than 200 and less than 250: %v\n",x)
	default: 
	fmt.Println("error")
	}
*/
	/*if (x < 4 && y < 4) {
	} else if (x > 6 && y > 6) {
		fmt.Printf("both are greater than 6: %v\n %v", x, y)
	  } else if (x >= 4 && x <= 6) {
		fmt.Printf("in between 4 and 6: %v\n", x)
	  } else if y != 5 {
		fmt.Printf("y is not 5: %v\n", y)
	  } else {
		fmt.Printf("none of the prev case were met")

	  }
*/

	  switch {
	   case x < 4 && y < 4:
        fmt.Printf("both are less than 4: %v\n %v", x, y)
	   case x > 6 && y > 6:
		fmt.Printf("both are greater than 6: %v\n %v", x, y)
	   case x >= 4 && x <= 6:
		fmt.Printf("in between: %v\n", x)
	   case  y != 5:
		fmt.Printf("y is not 5: %v\n", y)
	   default: fmt.Printf("none of the prev case were met") 
	   }
	}
}