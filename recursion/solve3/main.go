package main

import "fmt"

func main() {
	rr := titleToNumber("ZB")
	fmt.Println(rr)
	ff := reverseStr("abcdefg")
	fmt.Println(ff)
}

func titleToNumber(columnTitle string) int {
	result := 0
	power := 1
	for i := len(columnTitle) - 1; i >= 0; i-- {
		charValue := int(columnTitle[i] - 'A' + 1)
		result += charValue * power // This line adds the contribution of the current character to the overall result. It multiplies the numerical value of the character by the power of 26 corresponding to its position from the right and adds it to the result.
		power *= 26                 // This line updates the power by multiplying it by 26. This is necessary because each position from the right represents a higher power of 26 in the Excel column numbering system.
	}
	return result
}

func reverseStr(s string) string {
	sbyte := []rune(s)
	first, last := 0, len(sbyte)-1
	for first <= last {
		sbyte[first], sbyte[last] = sbyte[last], sbyte[first]
		first++
		last--
	}
	return string(sbyte)
}
