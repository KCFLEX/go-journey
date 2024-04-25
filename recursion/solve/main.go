package main

import "fmt"

var symbol = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

func getSymbolValue(val uint8) int { // Getting the initial value by getting the symbol value of the last character in the string.
	return symbol[fmt.Sprintf("%c", val)]
}

func romanToInt(s string) int {
	lastIndex := len(s) - 1
	value := getSymbolValue(s[lastIndex]) // to get the last letter of the string
	for i := 0; i < lastIndex; i++ {
		int1 := getSymbolValue(s[i])
		int2 := getSymbolValue(s[i+1])
		if int1 >= int2 {
			value += int1
		} else if int1*5 == int2 || int1*10 == int2 {
			value -= int1
		}
	}
	return value
}
func main() {
	//f := []byte{'H', 'a', 'n', 'n', 'a', 'h'}
	//fmt.Println(f)
	//fmt.Println(f[0])

	//reverseString(f)

	//fmt.Println(string(f))
	//romanToInt()

	//c := I + V
	//fmt.Println(c)

	//j := "thirst"
	//fmt.Println(string(j[1]))

	m := getSymbolValue('M')
	fmt.Println(m)
	tt := romanToInt("MCMXCIV")
	fmt.Println(tt)
	//fmt.Println(m)
	rr := []string{"M", "C", "M", "X", "C", "IV"}
	t := len(rr)

	fmt.Println(t)

}

/*
func reverseString(s []string) []string {
	var c []string
	for i := len(s) - 1; i >= 0; i-- {
		fmt.Println(s[i])
		c = append(c, s[i])
	}
	return c
}


func reverseString(s []byte) {
	left, right := 0, len(s)-1 // Initializing left pointer at index 0 and right pointer at the last index of the input byte slice.
	for left < right {         // Starting a loop that runs until the left pointer is less than the right pointer.
	 s[left], s[right] = s[right], s[left]
		left++  // Incrementing the left pointer to move towards the right end of the slice.
		right-- // Decrementing the right pointer to move towards the left end of the slice.
	}
}
*/
