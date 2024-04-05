package main

import (
	"fmt"
	"os"
)

func main() {

	doneChans := make([]chan bool, 5)
	errChans := make([]chan error, 5)

	for i := 0; i < 5; i++ {
		doneChans[i] = make(chan bool)
		errChans[i] = make(chan error)
		go readFile(doneChans[i], errChans[i])
	}

	for index := 0; index < 5; index++ {
		select {
		case <-doneChans[index]:
		case err := <-errChans[index]:
			fmt.Println("Goroutine", index, "failed with error", err)
		}
	}

}

func readFile(doneChannel chan bool, errorChannel chan error) {

	file, err := os.Open("dummy.txt")

	if err != nil {
		errorChannel <- err
		return
	}

	/*
		The defer keyword is used to schedule a function call to be run after the function completes
		Here, we are deferring the file.Close() function call
		This means that the file will be closed after the readFile function completes
	*/
	defer file.Close()

	fmt.Println("File opened successfully")

	// We need to close the file after we are done with it
	// But we want to make sure that the file is closed even if an error occurs
	// file.Close()

	doneChannel <- true
}
