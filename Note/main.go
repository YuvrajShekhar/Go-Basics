package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note/note"
)

func main() {
	title, content := getNoteData()

	userNote, _ := note.New(title, content)

	userNote.DisplayNote()
	userNote.SaveNote()
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
