package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

type person struct {
	first string
}

// writer interface
func (p person) writeOut(w io.Writer) error {
	_, err := w.Write([]byte(p.first))
	return err
}

func main() {
	// the code below creates a file
	f, err := os.Create("output.txt") // this throws an error
	// so we catch the error with if statement below
	if err != nil {
		log.Fatalf("error %s", err)
	}
	// we need to close a file after opening it by using the syntax below
	defer f.Close()
	///////////////////////////////////////////////////////
	s := []byte("Hello gophers") // this also throws an error
	_, err = f.Write(s)
	if err != nil {
		log.Fatalf("error %s", err)
	}

	// byte buffer
	// this is known as writing to buffer
	b := bytes.NewBufferString("Hello ")
	fmt.Println(b.String())
	b.WriteString("gophers! ")
	fmt.Println(b.String())
	b.WriteString("how are you")
	fmt.Println(b.String())
	b.WriteString(" I am Vincent")
	fmt.Println(b.String())

}
