package main

import (
	"fmt"
)

var symbol = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

func getSymbolValue(val uint8) int {
	return symbol[fmt.Sprintf("%c", val)] /*The expression fmt.Sprintf("%c", val) is used to convert the uint8 value val to its corresponding ASCII character representation. The %c verb in fmt.Sprintf is used to format the uint8 value as a character.
	The resulting character is then used as the key to look up the value in the symbol map. The value retrieved from the map is of type int, as indicated by the return type of the function.*/
}

func romanToInt(s string) int {
	lastIndex := len(s) - 1 // this to get the last index postion
	// now get the value of the lastIndex by using the getSymbolValue() function
	lastIndexValue := getSymbolValue(s[lastIndex])
	// use a for loop to iterate through the string
	for i := 0; i < lastIndex; i++ {
		// now get the value of the of the first letter of the string using getSymbolValue() function
		firstIndexValue := getSymbolValue(s[i])
		secondIndexValue := getSymbolValue(s[i+1])
		// now we use an if to write our conditions
		if firstIndexValue >= secondIndexValue {
			lastIndexValue += firstIndexValue
		} else if firstIndexValue*5 == secondIndexValue || firstIndexValue*10 == secondIndexValue {
			lastIndexValue -= firstIndexValue
		}
	}
	return lastIndexValue
}

func main() {
	//ggg := reverseVowels("leetcode")
	//fmt.Println(ggg)

	aaa := romanToInt("MCMXCIV")
	fmt.Println(aaa)
	fmt.Println("--------")
	f := []byte{'h', 'e', 'l', 'l', 'o'}
	reverseString(f)

	fmt.Println(string(f))
}

func reverseString(s []byte) {
	first, last := 0, len(s)-1 // get the positions of the first and last string
	for first < last {
		s[first], s[last] = s[last], s[first]
		first++
		last--
	}
}

/*
func reverseVowels(s string) string {
	low, high := 0, len(s)-1
	vowels := "aeiouAEIOU"
	sRune := []rune(s)
	for low < high {
		// Move the `low` pointer to the right until it points to a vowel
		for low < high && !strings.Contains(vowels, string(sRune[low])) {
			low++
		}
		for low < high && !strings.Contains(vowels, string(sRune[high])) {
			high--
		}
		if low < high {
			sRune[low], sRune[high] = sRune[high], sRune[low]
			low++
			high--
		}
	}
	return string(sRune)
}
*/
