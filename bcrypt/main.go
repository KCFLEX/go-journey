package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

/*
In Go, bcrypt is a package that provides functions for hashing and comparing passwords
using the bcrypt algorithm. Bcrypt is a popular password hashing algorithm that is
designed to be slow and computationally expensive, making it resistant to brute-force attacks.
*/
func main() {
	s := `password123`
	// we are hashing our password
	bytes, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost) // this will return a slice of bytes
	if err != nil {
		fmt.Println("error")
	}
	fmt.Println(s)
	fmt.Println(bytes)
	// if someone gave me a password i would compare it to the hashed password by doing the following below to check if the password is correct
	loginPword1 := `password123`
	err = bcrypt.CompareHashAndPassword(bytes, []byte(loginPword1))
	if err != nil {
		fmt.Println("login not happening")
	}
	fmt.Println("You're logged in ")
}
