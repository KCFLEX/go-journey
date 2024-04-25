package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

/*
func myFunc(wg *sync.WaitGroup) { //we always need to use pointer sematics for waitgroup
	time.Sleep(3 * time.Second) // this is saying we need to pause the execution of the progarm for one second
	fmt.Println("finished Executing Goroutine")
	wg.Done() // it is saying the function has finished execution
}

func main() {
	fmt.Println("Go Tutorial")
	var wg sync.WaitGroup
	wg.Add(1)
	go myFunc(&wg) // we always need to use pointer sematics for waitgroup
	wg.Wait()      // this is telling the program to not go passed this line until wg.Add becomes 0
	fmt.Println("finished executing my go program")
}
*/

var urls = []string{
	"https://google.com",
	"https://tutorialedge.net",
	"https://instagram.com",
}

func fetchStatus(w http.ResponseWriter, r *http.Request) { //end point
	var wg sync.WaitGroup
	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			resp, err := http.Get(url)
			if err != nil {
				fmt.Fprintf(w, "%v\n", err)
			}
			fmt.Fprintf(w, "%v\n", resp.Status)
			wg.Done()
		}(url)
	}
	wg.Wait()
}

func main() {
	fmt.Println("Go Tutorial")
	http.HandleFunc("/", fetchStatus)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
