package main

import (
	"fmt"
	"strings"
)

const s = `The key with TDD and BDD is to not
let methodological purity become the enemy of the good enough.
To goal is to write well-tested software that we can reliably
augment over many versions. A simple suite of working and complete
unit tests, like those shown above, are better than spending a
bunch of time arguing over proper TDD methods, or implementing
fancy-looking BDD tests that aren't actually complete.`

func main() {
	// trying to convert the first letter of every word to uppercase

	xs := strings.Split(s, " ")

	fmt.Println(xs)
	fmt.Println("--------")
	for i, v := range xs {
		xs[i] = strings.ToUpper(string(v[0])) + v[1:]
	}
	fmt.Println(xs)

	jj := strings.Join(xs, " ")
	fmt.Println(jj)

}
