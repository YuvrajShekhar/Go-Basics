package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Todo struct {
	Content string
}

func (todo Todo) Display() {
	fmt.Printf("The content of your todo is as follows:\n%v", todo.Content)
}

func (todo Todo) Save() error {
	filename := "todo.json"

	json, err := json.Marshal(todo)

	if err != nil {
		return err
	}

	return os.WriteFile(filename, json, 0644)
}

func New(content string) (Todo, error) {
	if content == "" {
		return Todo{}, errors.New("invalid Input")
	}

	return Todo{
		Content: content,
	}, nil
}
