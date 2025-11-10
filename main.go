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
	taskLoader := NewLoader()
	taskSaver := NewSaver()

	switch command {
	case "add":
		addHandler := NewTaskAddHandler(taskLoader, taskSaver)
		addHandler.Handle()
	case "update":
		updateHandler := NewTaskUpdateHandler(taskLoader, taskSaver)
		updateHandler.Handle()
	case "delete":
		deleteHandler := NewTaskDeleteHandler(taskLoader, taskSaver)
		deleteHandler.Handle()
	case "mark-in-progress":

	case "mark-done":

	case "list":

	default:
		fmt.Println("Unknown command:", command)
	}
}
