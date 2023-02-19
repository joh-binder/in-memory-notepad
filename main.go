package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("Enter the maximum number of notes:\n")
	var capacity int
	_, err := fmt.Scan(&capacity)
	if err != nil {
		fmt.Print("[Error] Could not parse input as an integer\n")
		return
	}

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines) // makes scanner read until new line

	notepad := make([]string, capacity)
	noteCounter := 0

	mainLoop := true
	for mainLoop {
		// prompt user for input
		fmt.Print("Enter a command and data: ")
		scanner.Scan()
		line := scanner.Text()
		lineAsWords := strings.Fields(line)
		command := lineAsWords[0] // separate command from remaining input
		inputData := lineAsWords[1:]

		switch command {
		case "create":
			if noteCounter >= capacity {
				fmt.Print("[Error] Notepad is full\n")
				continue
			}
			if len(inputData) == 0 {
				fmt.Print("[Error] Missing note argument\n")
				continue
			}
			currentNote := strings.Join(inputData, " ") // turn slice of strings to a single string
			notepad[noteCounter] = currentNote          // write currentNote to notepad
			noteCounter += 1
			fmt.Print("[OK] The note was successfully created\n")
		case "update":
			// check if there are enough arguments
			if len(inputData) == 0 {
				fmt.Print("[Error] Missing position argument\n")
				continue
			}
			if len(inputData) == 1 {
				fmt.Print("[Error] Missing note argument\n")
				continue
			}

			position, err := strconv.Atoi(inputData[0])

			// check if position is valid
			if err != nil {
				fmt.Printf("[Error] Invalid position: %s\n", inputData[0])
				continue
			}
			if position < 1 || position > capacity {
				fmt.Printf("[Error] Position %d is out of the boundaries [1, %d]\n", position, capacity)
				continue
			}
			if position > noteCounter {
				fmt.Print("[Error] There is nothing to update\n")
				continue
			}

			newNote := strings.Join(inputData[1:], " ") // combine input data to a single string (except position)
			notepad[position-1] = newNote               // replace old note
			fmt.Printf("[OK] The note at position %d was successfully updated\n", position)
		case "delete":
			if len(inputData) == 0 {
				fmt.Print("[Error] Missing position argument\n")
				continue
			}

			position, err := strconv.Atoi(inputData[0])

			// check if position is valid
			if err != nil {
				fmt.Printf("[Error] Invalid position: %s\n", inputData[0])
				continue
			}
			if position < 1 || position > capacity {
				fmt.Printf("[Error] Position %d is out of the boundaries [1, %d]\n", position, capacity)
				continue
			}
			if position > noteCounter {
				fmt.Print("[Error] There is nothing to delete\n")
				continue
			}

			for i := position; i < noteCounter; i++ {
				notepad[i-1] = notepad[i] // for all notes *after* position: shift 1 to the left
				/* loop starts at i == position (and not at position + 1) because the position number begins
				at 1 instead of 0, so the +1 is automatically included */
			}
			notepad[noteCounter-1] = "" // "delete" the last note by overwriting with an empty string
			noteCounter -= 1
			fmt.Printf("[OK] The note at position %d was successfully deleted\n", position)
		case "list":
			if noteCounter == 0 {
				fmt.Print("[Info] Notepad is empty\n")
			}
			// prints all notes that currently exist
			for i := 0; i < noteCounter; i++ {
				fmt.Printf("[Info] %d: %s\n", i+1, notepad[i])
			}
		case "clear":
			notepad = make([]string, capacity) // assign notepad to a new, empty slice of the same capacity
			noteCounter = 0                    // resets counter
			fmt.Print("[OK] All notes were successfully deleted\n")
		case "exit":
			mainLoop = false
		default:
			fmt.Print("[Error] Unknown command\n")
		}
	}

	fmt.Print("[Info] Bye!\n")
}
