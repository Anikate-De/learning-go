package main

import (
	"fmt"
	"time"
)

func main() {

	/*
		GoRoutines are lightweight threads managed by the Go runtime.
		They are used to run functions concurrently.

		We can define a GoRoutine by using `go`
		go funcName()

		However, the go keyword just 'dispatches' the function to run concurrently.
		There is no guarantee that the function will run before the main function ends.
		Here, if we do not use a channel, we will see no output from greet()
		go greet()
	*/
	/*
		Channels
		Channels are used to communicate between GoRoutines.
		They are used to send and receive data between GoRoutines. (like Send/ReceivePort in Dart)

		Use the make function to create a channel
		ch := make(chan int)
	*/

	done := make(chan bool)
	go slowGreet(done)
	go greet()

	<-done // Wait for the message from the channel

	// The output of the above code will be:
	// Hello there!
	// I was sleeping
}

func greet() {
	fmt.Println("Hello there!")
}

func slowGreet(doneChan chan bool) {
	time.Sleep(time.Second * 2) // Simulate a long running task
	fmt.Println("I was sleeping")

	// Send a message to the channel
	doneChan <- true
}
