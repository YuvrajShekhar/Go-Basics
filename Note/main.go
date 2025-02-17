package main

import (
	"fmt"

	"example.com/note/note"
)

func main() {
	title, content := getNoteData()

	userNote, err := note.New(title, content)
}

func getNoteData() (string, string) {
	title := getUserInput("Enter Note Title : ")
	content := getUserInput("Enter Note Title : ")
	return title, content
}

func getUserInput(promt string) string {
	fmt.Print(promt)
	var value string
	fmt.Scanln(&value)
	return value
}
