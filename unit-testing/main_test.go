package main

import (
	"log"
	"testing"
)

// add test
func TestAdd(t *testing.T) {
	goal := add(3, 3)
	total := 6
	if goal != total {
		log.Fatalf("error the total should be %v not ", total)
	}

}

// substract test
func TestSubstract(t *testing.T) {
	result := subtract(7, 3)
	expected := 4
	if result != expected {
		t.Errorf("error - subtract(7, 3) should be %v instead it %v", expected, result)
	}
}

// doMath test
func TestDoMath(t *testing.T) {
	result := doMath(2, 2, add)
	expected := 4
	if result != expected {
		t.Errorf("error - doMath(2, 2, add) should be %v and not %v", expected, result)
	}
}
