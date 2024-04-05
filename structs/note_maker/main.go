package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"de.anikate/note_maker/note"
)

const jsonFileName string = "note.json"

func main() {
	var newNote *note.Note
	title, content := getNoteInput()

	newNote = note.New(title, content)
	newNote.Print()

	result := newNote.SaveToJson(jsonFileName)

	if !result {
		fmt.Println("Failed to save your note.")
	} else {
		fmt.Println("Note saved successfully!")
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
