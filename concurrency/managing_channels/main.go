package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

var globalNumber = 0

func main() {

	// Make slices of channels
	doneChans := make([]chan bool, 5)
	errChans := make([]chan error, 5)

	for i := 0; i < 5; i++ {
		// Initialize the channels
		doneChans[i] = make(chan bool)
		errChans[i] = make(chan error)
		go incrementGlobal(doneChans[i], errChans[i])
	}

	// Go allows us to use the select statement to wait for multiple channels
	// It is quite similar to the switch case statement and only waits for the first channel to return a value, of all cases
	for index := 0; index < 5; index++ {
		// Here, we are waiting for either the done channel or the error channel
		select {
		case <-doneChans[index]:
			fmt.Println("Goroutine", index, "finished")
		case err := <-errChans[index]:
			fmt.Println("Goroutine", index, "failed with error", err)
		}
	}

}

/*
Accept both a done channel and an error channel
We can use the error channel to send an error message if something goes wrong
*/
func incrementGlobal(doneChannel chan bool, errorChannel chan error) {
	time.Sleep(time.Second * 2) // Simulate a long running task

	// Simulate a random error
	if rand.Intn(2) == 1 {
		// Send an error message to the error channel
		errorChannel <- errors.New("random error")
		return
	}

	globalNumber++
	fmt.Println("Global number is now", globalNumber)

	// If the goroutine finishes successfully, send a message to the done channel
	doneChannel <- true
}
