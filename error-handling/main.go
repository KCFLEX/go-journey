package main

import (
	"fmt"
	"io"
	"os"
)

// if error == nil it means there's no error
func main() {
	//n, err := fmt.Println("Hello")
	//if err != nil { // this means if err is not equal to nil(mean there's an err do what is in between the curly braces {...})
	//	fmt.Println(err)
	//	}
	//fmt.Println(n)
	// different code ------------
	/*	var answer1, answer2, answer3 string

		fmt.Print("Name:")
		_, err := fmt.Scan(&answer1)
		if err != nil {
			panic(err)
		}

		fmt.Print("Fav Food: ")
		_, err = fmt.Scan(&answer2)
		if err != nil {
			panic(err)
		}

		fmt.Print("Fav Sport: ")
		_, err = fmt.Scan(&answer3)
		if err != nil {
			panic(err)
		}

		fmt.Println(answer1, answer2, answer3)
	*/
	// different code ------------
	//Overall, the code snippet below creates a new file named "names.txt", checks for any errors during the file creation, writes the string "wassup" to the file, and ensures that the file is closed properly.
	/*
		f, err := os.Create("names.txt")
		if err != nil {
			fmt.Println(err)
		}
		defer f.Close() // after creating a file please always try to close it

		r := strings.NewReader("James bond") // we used NewReader to pass in the string

		io.Copy(f, r) // io.Copy reads from the r reader and writes the data to the f file. In this case, it will write the string "wassup" to the "names.txt" file.
	*/

	// different code ------------

	f, err := os.Open("names.txt")
	if err != nil {
		fmt.Println(err)
		return // exits the current function
	}
	defer f.Close()

	bs, err := io.ReadAll(f) // io.ReadAll(f) line declares a variable bs of type []byte and assigns it the result of reading all the content from the file descriptor f using the io.ReadAll function. If there is an error while reading the file, it will be assigned to the err variable.
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))
}
