package main

import (
	"fmt"
	"log"
	"os"
)

const tasksFile = "tasks.json"
const IDsFile = "id_tracker.json"

func main() {
	if len(os.Args) < 2 {
		log.Fatal("usage: task-cli <command> [arguments]")
	}

	command := os.Args[1]

	switch command {
	case "add":
		addHandler := NewTaskAddHandler()
		addHandler.Handle()
	case "update":
		updateHandler := NewTaskUpdateHandler()
		updateHandler.Handle()
	case "delete":
		deleteHandler := NewTaskDeleteHandler()
		deleteHandler.Handle()
	case "mark-in-progress":

	case "mark-done":

	case "list":

	default:
		fmt.Println("Unknown command:", command)
	}
}
