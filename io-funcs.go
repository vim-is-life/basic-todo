package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
)

// die will fatally log an error and exit
func die(msg string, err error) {
	log.Fatalln(msg+" ", err)
}

// readTodos will read the todos and store them in todoList.
// TODO better error handling
func readTodos(todoList *[]Todo) {
	inFile, err := os.Open(TODO_FILE)

	// if we weren't able to open the file, we want to just return gracefully
	if errors.Is(err, os.ErrNotExist) {
		// case where the file does not exist
		return
	} else if err != nil {
		msg := "readTodos: Something went wrong with opening the file"
		die(msg, err)
	}

	defer inFile.Close()
	r := csv.NewReader(inFile)
	for {
		var newTodoItem Todo
		record, err := r.Read()
		// we're at the end of the file here so break out
		if err == io.EOF {
			break
		}
		if err != nil {
			msg := "readTodos: Something went wrong while reading the CSV"
			die(msg, err)
		}

		// 1st field is id
		fmt.Sscan(record[0], &newTodoItem.id)

		// 2nd field is name
		newTodoItem.name = record[1]

		// 3rd field is kind
		fmt.Sscan(record[2], &newTodoItem.kind)

		// 4th field is state
		fmt.Sscan(record[3], &newTodoItem.state)
		*todoList = append(*todoList, newTodoItem)
	}
}

// saveTodos will save todoList
func saveTodos(todoList []Todo) {
	outFile, err := os.Create(TODO_FILE)
	if err != nil {
		msg := "saveTodos: Something went wrong with opening the file to write"
		die(msg, err)
	}
	defer outFile.Close()
	w := csv.NewWriter(outFile)
	for _, todoItem := range todoList {
		// doing this so we can get the enum values as their under the hood
		// integer values instead of the string conversions of them.
		// TODO test if i can just do fmt.Sprint on an enum directly without it
		// being turned into a string according to its String method
		var todoKind int
		var todoState int

		switch todoItem.kind {
		case CategoryProject:
			todoKind = 0
		case CategoryHomework:
			todoKind = 0 + 1
		case CategoryReading:
			todoKind = 0 + 2
		}

		switch todoItem.state {
		case StateTodo:
			todoState = 0
		case StateInProgress:
			todoState = 0 + 1
		case StateDone:
			todoState = 0 + 2
		}
		records := []string{fmt.Sprint(todoItem.id), todoItem.name,
			fmt.Sprint(todoKind), fmt.Sprint(todoState)}
		if err := w.Write(records); err != nil {
			die("saveTodo: error writing record to csv: ", err)
		}
	}

	// Write any buffered data to the underlying writer (standard output).
	w.Flush()
	if err := w.Error(); err != nil {
		die("saveTodos: error flushing file: ", err)
	}
}
