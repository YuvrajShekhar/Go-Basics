package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note/note"
	"example.com/note/todo"
)

func main() {
	title, content := getNoteData()
	todoText := getUserInput("Enter todo item : ")

	userNote, _ := note.New(title, content)
	userTodo, _ := todo.New(todoText)

	userNote.DisplayNote()
	userNote.SaveNote()
	userTodo.DisplayTodo()
	userTodo.SaveTodo()
	fmt.Println("Note Saved succefully")
}

func getNoteData() (string, string) {
	title := getUserInput("Enter Note Title : ")
	content := getUserInput("Enter Note Title : ")
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
