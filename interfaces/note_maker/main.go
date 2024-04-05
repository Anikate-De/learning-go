package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"de.anikate/note_maker/note"
	"de.anikate/note_maker/todo"
)

/*
In Go, interfaces are defined as types.
It can be considered as a contract that a type can choose to implement.
Actual Definition: An interface is a collection of method signatures that a Type can implement.

This is the general convention for naming interfaces in Go:
- Name should be the method name + "er" suffix.
- If there are embedded interfaces, the name should be an attribute (-able)

Interfaces in Go are implicit.
i.e., we don't explicitly specify that a type implements an interface.
rather, if a type has all the methods of an interface, it is considered to implement that interface.
*/
type saver interface {
	SaveToJson() bool // Method signature, no implementation
}

type printer interface {
	Print()
}

type outputtable interface {
	saver   // Embedding the saver interface
	printer // Embedding the printer interface
}

func main() {
	var newNote *note.Note
	title, content := getNoteInput()
	task := getTodoInput()

	newNote = note.New(title, content)

	newTodo := todo.New(task)

	save(newNote, "Note")
	save(newTodo, "Todo")

}

/*
By accepting an outputtable interface, we can be sure that argument passed to this function will have the methods defined in the interface.
*/
func save(snippet outputtable, prompt string) {
	snippet.Print()
	result := snippet.SaveToJson()

	if !result {
		fmt.Println("Failed to save your", prompt)
	} else {
		fmt.Println(prompt, "saved successfully!")
	}
}

func tidyInput(input string) (val string) {
	val = strings.TrimSuffix(input, "\n")
	val = strings.TrimSuffix(val, "\r")

	return
}

func getNoteInput() (title, content string) {

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Your note:")

	fmt.Print("Title: ")
	title, _ = reader.ReadString('\n')
	title = tidyInput(title)

	fmt.Print("Content: ")
	content, _ = reader.ReadString('\n')
	content = tidyInput(content)

	return
}

func getTodoInput() (task string) {

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("TODO Task: ")
	task, _ = reader.ReadString('\n')
	task = tidyInput(task)

	return
}
