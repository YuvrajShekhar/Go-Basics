package main

import (
	"fmt"

	"example.com/note/note"
)

func main() {
	title, content := getNoteData()

	userNote, _ := note.New(title, content)

	userNote.DisplayNote()
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
