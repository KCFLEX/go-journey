package main

import (
	"testing"
)

// Tests must
// ● be in a file that ends with “_test.go”
// ● put the file in the same package as the one being tested
// ● be in a func with a signature “func TestXxx(*testing.T)”
func TestSum(t *testing.T) {
	totalsum := Sum(1, 2, 3, 4, 5)
	if totalsum != 15 {
		t.Errorf("totalsum should be %d instead it gave %v", 15, totalsum)
	}
}
