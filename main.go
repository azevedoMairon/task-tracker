package main

import (
	"encoding/json"
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

func main() {
	if len(os.Args) < 2 {
		log.Fatal("usage: task-cli <command> [arguments]")
	}

	command := os.Args[1]

	switch command {
	case "add":
		if len(os.Args) < 4 {
			log.Fatal("usage: task-cli add <task name> <description> [arguments]")
		}
		taskName := os.Args[2]
		fmt.Println("Adding task:", taskName)

		f, err := os.Create(tasksFile)
		if err != nil {
			log.Fatalf("error creating task file: %w", err)
		}
		defer f.Close()

		t := Task{
			ID:          1,
			Description: os.Args[3],
			Status:      StatusTodo,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}

		task, err := json.Marshal(t)
		if err != nil {
			fmt.Println("error marshaling task struct to json: %w", err)
		}
		_, err = f.Write(task)
		if err != nil {
			fmt.Println("error while writing task into %s file: %w", tasksFile, err)
		}

	case "update":

	case "delete":

	case "mark-in-progress":

	case "mark-done":

	case "list":

	default:
		fmt.Println("Unknown command:", command)
	}
}
