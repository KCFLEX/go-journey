package main

import "fmt"

/*func (r reciever) identifier(p parameter(s)) (return(s)) {
...
}*/
// in go everything is passed by value

func main() {
	foo()
	fmt.Println(add(1, 2))
	fmt.Println(convT(32))
	fmt.Println(convW(205))
	fmt.Println(dogYears("Scott", 11))
	fmt.Println("--------------")
	xi := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	//unfurling a slice: is passing the elements of a slice as individual arguments
	//this means turning the slice of number to individual/seperate numbers
	multiply(xi...)

}

func foo() {
	fmt.Println("i am from foo")
}

func add(a, b int) /*return = int */ int {
	return a + b
}

func convT(celsius int) /*return = int */ int {
	var f int = (celsius * 9 / 5) + 32
	return f
}

func convW(lbs float64) /*return = float64 */ float64 {
	var kg float64 = lbs * 0.45359237
	return kg

}

func dogYears(name string, h int) (string, int) {
	var dogAge = h * 7
	return fmt.Sprintf("my dog's name is %v age is", name), dogAge

}

// variadic parameters in our functions
func multiply(ii ...int /*variadic parameter*/) int {
	fmt.Println(ii)
	fmt.Printf("%T\n", ii)
	fmt.Println("--------------")
	o := 0
	for _, v := range ii {
		o += v // to sum all the numbers in the slice together

	}
	fmt.Println(o)
	return o
}
