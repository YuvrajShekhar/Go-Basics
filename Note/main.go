package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note/note"
	"example.com/note/todo"
)

type saver interface {
	Save() error
}

type outputable interface {
	saver
	Display()
}

func main() {
	title, content := getNoteData()
	todoText := getUserInput("Enter todo item : ")

	userNote, _ := note.New(title, content)
	userTodo, _ := todo.New(todoText)

	err := outputData(userNote)

	if err != nil {
		fmt.Println("Erorr in Note")
	}

	err = outputData(userTodo)

	if err != nil {
		fmt.Println("Erorr in Todo list")
	}

}

func getNoteData() (string, string) {
	title := getUserInput("Enter Note Title : ")
	content := getUserInput("Enter Note Content : ")
	return title, content
}

func getUserInput(promt string) string {
	fmt.Print(promt)
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		return ""
	}
	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}

func saveData(data saver) error {
	err := data.Save()

	if err != nil {
		fmt.Println("Saving note failes")
		return err
	}

	fmt.Println("Saving note successfull")
	return nil
}

func outputData(data outputable) error {
	data.Display()
	return saveData(data)
}
