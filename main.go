package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"
)

type Status int

const (
	StatusTodo Status = iota
	StatusInProgress
	StatusDone
)

type Task struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	Status      Status    `json:"status"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

const tasksFile = "tasks.json"
const IDsFile = "ids.json"

func fileExists(f string) (bool, error) {
	if _, err := os.Stat(f); errors.Is(err, os.ErrNotExist) {
		return false, nil
	} else if err != nil {
		return false, err
	} else {
		return true, nil
	}
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("usage: task-cli <command> [arguments]")
	}

	command := os.Args[1]

	switch command {
	case "add":
		addHandler := NewAddTaskHandler()
		addHandler.Handle()

	case "update":

	case "delete":

	case "mark-in-progress":

	case "mark-done":

	case "list":

	default:
		fmt.Println("Unknown command:", command)
	}
}
