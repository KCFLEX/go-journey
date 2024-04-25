package main

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
)

type Cities struct {
	Name       string
	Population int
}

func main() {
	f, err := os.Open("cities.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	rows := genRows(f)
	// fan out patern
	//more than one worker ccmpeting to consume the same channel
	//       __ worker1
	// rows /__ worker2
	//      \__ worker3
	//       __ workern...
	ur1 := upperCityName(rows)
	ur2 := upperCityName(rows)
	// fanIn patern consolidates the outputs of multiple channels into one
	for c := range fanIn(ur1, ur2) {
		log.Println(c)
	}
}

// this function is to transform the city names to uppercase
func upperCityName(city <-chan Cities) <-chan Cities { // this func take in rows from the channel cities and returns a read-only channel(recieve only)
	out := make(chan Cities)
	go func() {
		for c := range city {
			//strings.ToUpper(c.Name)
			out <- Cities{
				Name:       strings.ToUpper(c.Name),
				Population: c.Population,
			}
			// remember to always close your channel when done
		}
	}()
	return out // make sure to return the out channel
}

// a channel a way for communicating between difernt go routines and allows us to exchange data between different go routines

// a function that takes an io.reader and returns a channel
func genRows(r io.Reader) chan Cities {
	out := make(chan Cities)
	go func() {
		reader := csv.NewReader(r) // csv.NewReader(r) creates a new CSV reader that reads data from the given r input source.
		_, err := reader.Read()
		if err != nil {
			log.Fatal(err)
		}
		for {
			row, err := reader.Read() //This line of code is using the Read() function from the reader object. It is attempting to read a row of data from the reader. The Read() function returns two values: row and err. The row variable will store the data read from the reader, and the err variable will store any error that occurred during the reading process.
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatal(err)
			}
			// the code above reads the data from the row and now we want to pass that data to the out channel that we prevoiusly created
			populationInt, err := strconv.Atoi(row[9])
			if err != nil {
				continue
			}
			out <- Cities{
				Name:       row[1],
				Population: populationInt,
			}
		}
		close(out) //  when ever you are using a channel the only way to tell a goroutine that there's no more data in the channel is when you close the channel
		// once the channel is closed ot means there is no more data
	}()
	return out // make sure to return the out channel
}

func fanIn(chans ...<-chan Cities) <-chan Cities { // fanIn takes in a set of channels it returns just a single channel
	out := make(chan Cities)
	var wg sync.WaitGroup
	wg.Add(len(chans))        // we add how many channels we are reading from
	for _, c := range chans { //This line iterates over the chans , which contains the input channels we want to read from.
		go func(city <-chan Cities) { // This line starts a new goroutine for each input channel c. The goroutine receives the input channel as an argument (city <-chan Cities) and reads all the data of the workers.
			for r := range city { // sends all the data to the out channel
				out <- r
			}
			wg.Done()

		}(c)
	}
	go func() {
		wg.Wait()

	}()
	return out
}
