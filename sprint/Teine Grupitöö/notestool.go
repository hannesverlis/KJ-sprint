package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 || os.Args[1] == "help" {
		fmt.Println("Usage: ./notestool [COLLECTION_NAME]")
		return
	}

	collectionName := os.Args[1]
	dbFile := collectionName + ".txt"

	for {
		fmt.Printf("\nCollection: %s\n", collectionName)

		notes, err := readNotes(dbFile)
		if err != nil {
			fmt.Println("Error reading notes:", err)
			break
		}

		if len(notes) > 0 {
			fmt.Println("Notes:")
			for i, note := range notes {
				fmt.Printf("%d. %s\n", i+1, note)
			}
		} else {
			fmt.Println("No notes in this collection yet.")
		}

		fmt.Println("\nMenu:")
		fmt.Println("1. Add note")
		fmt.Println("2. Remove note")
		fmt.Println("3. Exit")

		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			addNote(dbFile)
		case 2:
			removeNote(dbFile, notes)
		case 3:
			return
		default:
			fmt.Println("Invalid choice.")
		}
	}
}

// ... (Functions for reading, adding, and removing notes)