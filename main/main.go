package main

import (
	"fmt"
	"go-practice/note"
	"go-practice/util"
)

func main() {
	title, content := util.GetNoteData()

	userNote, err := note.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}
	userNote.Display()
	err = userNote.Save()

	if err != nil {
		fmt.Println("Saving the note failed.")
		return
	}

	fmt.Println("Saving the note successfully.")
}
