package main

import (
	"fmt"
	"log"
	"os"

	"github.com/azevedoMairon/task-tracker/internal/file"
	"github.com/azevedoMairon/task-tracker/internal/tasks"
)

const tasksFile = "tasks.json"

var (
	taskLoader = file.NewLoader(tasksFile)
	taskSaver  = file.NewSaver(tasksFile)

	creator = tasks.NewCreator(taskLoader, taskSaver)
	reader  = tasks.NewReader(taskLoader)
	updater = tasks.NewUpdater(taskLoader, taskSaver)
	deleter = tasks.NewDeleter(taskLoader, taskSaver)
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("usage: task-cli <command> [arguments]")
	}

	command := os.Args[1]

	switch command {
	case "add":
		creator.Create()
	case "update":
		updater.Update()
	case "delete":
		deleter.Delete()
	case "mark-in-progress":
		updater.MarkInProgress()
	case "mark-done":
		updater.MarkDone()
	case "list":
		reader.Read()

	default:
		fmt.Println("Unknown command:", command)
	}
}
