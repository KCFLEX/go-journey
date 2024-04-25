package main

import "fmt"

func main() {
	b := [5]int{}
	for i := 0; i < 5; i++ {
       b[i] = i
	}
	for i, v := range b {
		fmt.Printf("index: %v   value: %v\n", i, v)
	}
      
    fmt.Println("------------")
	c := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	for i, v := range c {
		fmt.Printf("index: %v   value: %v\n", i, v)
	}
   
	d := make([]int, len(c))
   fmt.Println(d)

   copy(d, c)
   fmt.Println(d)
   d = append(d[0:5])
   fmt.Println(d)

   fmt.Println("------------")
   e := make([]int, len(c))
   fmt.Println(e)
   copy(e, c)
   fmt.Println(e)
   e = append(e[5:])
   fmt.Println(e)
   fmt.Println("------------")
   f := make([]int, len(c))
   fmt.Println(f)
   copy(f, c)
   fmt.Println(f)
   f = append(f[2:7])
   fmt.Println(f)
    fmt.Println("------------")
	h := make([]int, 10)
     fmt.Println(h)

	 copy(h, c)
	 fmt.Println(h)
	 h = append(h[1:6])
   fmt.Println(h)
   fmt.Println("----------")
   fmt.Println(c[0:5])
   fmt.Println(c[5:])
   fmt.Println(c[2:7])
   fmt.Println(c[1:6])
    c = append(c, 52,53,54,55)
	fmt.Println(c)
y := []int{56,57,58,59,60}

c = append(c, y...)
fmt.Println(c)
fmt.Println("----------")
x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
fmt.Println(x)
x = append(x[:3],x[9:]...)
fmt.Println(x)

US := make([]string, 0, 50)
    US = append(US, "Alabama",
    "Alaska",
    "Arizona",
    "Arkansas",
    "California",
    "Colorado",
    "Connecticut",
    "Delaware",
    "Florida",
    "Georgia",
    "Hawaii",
    "Idaho",
    "Illinois",
    "Indiana",
    "Iowa",
    "Kansas",
    "Kentucky",
    "Louisiana",
    "Maine",
    "Maryland",
    "Massachusetts",
    "Michigan",
    "Minnesota",
    "Mississippi",
    "Missouri",
    "Montana",
    "Nebraska",
    "Nevada",
    "New Hampshire",
    "New Jersey",
    "New Mexico",
    "New York",
    "North Carolina",
    "North Dakota",
    "Ohio",
    "Oklahoma",
    "Oregon",
    "Pennsylvania",
    "Rhode Island",
    "South Carolina",
    "South Dakota",
    "Tennessee",
    "Texas",
    "Utah",
    "Vermont",
    "Virginia",
    "Washington",
    "West Virginia",
    "Wisconsin",
    "Wyoming")
    length := len(US)
    capacity := cap(US)
    fmt.Printf("%v\n %v\n %v\n", US, length, capacity)
    for i := 0; i < len(US); i++ {
        fmt.Println(i, US[i])

    }
    fmt.Printf("%v\n %v\n", length, capacity)
    fmt.Println("----------")

     // create a slice of a slice of string 
     of :=  []string{"James", "Bond", "Shaken", "not stirred"}
     ph :=  []string{"Miss", "Moneypenny", "i'm 008"}
         bo := [][]string{of, ph}
          for i, v := range bo {
            fmt.Printf("%v\n %v\n", i, v)
            for a, b := range v {
                fmt.Printf("%v\n %v\n", a, b)
            }
          }
              
        
    }

	
