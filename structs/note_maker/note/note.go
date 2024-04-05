package note

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type Note struct {
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func New(title, content string) *Note {
	return &Note{
		title,
		content,
		time.Now(),
	}
}

func (n *Note) Print() {
	fmt.Printf("Your note:\n%s\n%s", n.Title, n.Content)
}

func (n *Note) SaveToJson(filename string) (result bool) {
	jsonObjBytes, err := json.MarshalIndent(n, "", "\t")

	if err != nil {
		fmt.Println("Unable to convert to valid JSON object.")
		result = false
	}

	err = os.WriteFile(filename, jsonObjBytes, 0644)
	if err != nil {
		fmt.Println("Unable to save JSON Object to the filesystem.")
		result = false
	}

	result = true

	return
}
