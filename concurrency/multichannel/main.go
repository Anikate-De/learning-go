package main

import (
	"fmt"
	"time"
)

func main() {

	done := make(chan bool)
	go greet(done)
	go slowGreet(done)
	go greet(done)

	/*
		Wait for the message from the channel, here we are waiting for 3 messages
		Otherwise, the program will exit before ALL the goroutines finish
		You may want to use a slice of channels using make, and a for loop to define multiple channels
	*/
	<-done
	<-done
	<-done

	/*
		Or you can iterate over the values returned by the channel
		```
		for i := range done {}
		```

		However, this will throw an error if the channel is not closed
		You can close the channel by using the close function in the slowest goroutine

		But this only makes sense if your program is aware of the slowest goroutine
	*/

}

func greet(doneChan chan bool) {
	fmt.Println("Hello there!")
	doneChan <- true
}

func slowGreet(doneChan chan bool) {
	time.Sleep(time.Second * 2) // Simulate a long running task
	fmt.Println("I was sleeping")

	// Send a message to the channel
	doneChan <- true
	close(doneChan) // Close the channel, it cannot be used to send messages anymore
}
