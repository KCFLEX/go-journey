package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	_, err := os.Open("no-file-text")
	if err != nil {
		//fmt.Println("err happened\n", err)
		log.Println("err happened\n", err) // this comes with a date file stamp
		//	log.Fatal(err) // this totally terminates the operating and no deferred functions run
		//	panic(err)
	}

	f, err := os.Create("log.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	log.SetOutput(f) // This line sets the output destination for the standard logger from the log package to the file f. This means that any log messages written using the log package will be written to the "log.txt" file.

	f2, err := os.Open("no-file.txt")
	if err != nil {
		//fmt.Println("err happened\n", err)
		log.Println("err happened\n", err) // this comes with a date file stamp
		//	log.Fatal(err)
		//	panic(err)
	}
	defer f2.Close()
	fmt.Println("check the log.txt file in the directory")
}

/* Panic is a built-in function that stops the ordinary flow
of control and begins panicking. When the function F calls panic,
execution of F stops, any deferred functions in F are executed normally,
and then F returns to its caller. To the caller, F then behaves like
a call to panic. The process continues up the stack until
all functions in the current goroutine
have returned, at which point the program crashes. Panics
can be initiated by invoking panic directly.
They can also be caused by runtime errors, such as out-of-bounds array accesses. */
