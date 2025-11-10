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

	updateHandler := NewTaskUpdateHandler(taskLoader, taskSaver)
	addHandler := NewTaskAddHandler(taskLoader, taskSaver)
	deleteHandler := NewTaskDeleteHandler(taskLoader, taskSaver)
	listHandler := NewTaskListHandler(taskLoader)

	switch command {
	case "add":
		addHandler.Handle()
	case "update":
		updateHandler.HandleUpdateDesc()
	case "delete":
		deleteHandler.Handle()
	case "mark-in-progress":
		updateHandler.HandleMarkInProgress()
	case "mark-done":
		updateHandler.HandleMarkDone()
	case "list":
		listHandler.Handle()

	default:
		fmt.Println("Unknown command:", command)
	}
}
