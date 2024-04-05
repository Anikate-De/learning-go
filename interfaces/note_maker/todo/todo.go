package todo

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type Todo struct {
	Task      string    `json:"task"`
	CreatedAt time.Time `json:"created_at"`
}

func New(task string) *Todo {
	return &Todo{
		task,
		time.Now(),
	}
}

func (n *Todo) Print() {
	fmt.Printf("Your todo:\n%s\n", n.Task)
}

func (n *Todo) SaveToJson() (result bool) {
	jsonObjBytes, err := json.MarshalIndent(n, "", "\t")

	if err != nil {
		fmt.Println("Unable to convert to valid JSON object.")
		result = false
	}

	err = os.WriteFile("todo.json", jsonObjBytes, 0644)
	if err != nil {
		fmt.Println("Unable to save JSON Object to the filesystem.")
		result = false
	}

	result = true

	return
}
