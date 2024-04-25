package main

import "fmt"

func main () {
	// arrays and slices
	c := []string{"Alabama", " Alaska", " Arizona", " Arkansas", " California", 
	"Colorado", " Connecticut", " Delaware", " Florida", "Georgia", "Hawai",
}
 length := len(c)

 fmt.Printf("c = %v\n and the length is = %v\n", c, length)

 for i, v := range c {
     fmt.Printf("index = %v and value = %v\n", i, v)
 }


 d := []int{41, 42, 43}
 //use the append syntax to add more values to your array  
d = append(d, 1,2,3,4,5,6,7,8,9,10,11)
 fmt.Println(d)
 fmt.Printf("-----------\n")
 d = append(d[:4], d[5:]...)
 fmt.Println(d)
 //fmt.Printf("-----------\n")
 //fmt.Printf("d - %v\n", d[4:])
  x := make([]int, 10)
  x[0] = 1
  x[1] = 3
  x[2] = 4
  x[3] = 5
  x[4] = 6

  fmt.Println(x)
}