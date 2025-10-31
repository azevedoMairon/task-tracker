package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
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
		if len(os.Args) < 3 {
			log.Fatal("usage: task-cli add <task name>")
		}

		if exists, err := fileExists(IDsFile); !exists {
			f, err := os.Create(IDsFile)
			if err != nil {
				log.Fatal("error creating IDs file: ", err)
			}
			defer f.Close()
		} else if err != nil {
			log.Fatal("error checking ids file existence: ", err)
		}

		b, err := os.ReadFile(IDsFile)
		if err != nil {
			log.Fatal("error reading IDs file: ", err)
		}

		id, err := strconv.Atoi(string(b))
		if err != nil {
			log.Fatal("error converting ID bytes to int: ", err)
		}
		id = id + 1

		// _, err = f.Write([]byte(strconv.Itoa(id)))
		// if err != nil {
		// 	log.Fatal("error writing ids file")
		// }

		f, err := os.Create(tasksFile)
		if err != nil {
			log.Fatal("error creating task file: ", err)
		}
		defer f.Close()

		t := Task{
			ID:          id,
			Description: os.Args[2],
			Status:      StatusTodo,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}

		task, err := json.Marshal(t)
		if err != nil {
			fmt.Printf("error marshaling task struct to json: %v", err)
		}
		_, err = f.Write(task)
		if err != nil {
			fmt.Printf("error while writing task into %s file: %v", tasksFile, err)
		}

		fmt.Printf("Task added succesfully (ID: %d)", t.ID)

	case "update":

	case "delete":

	case "mark-in-progress":

	case "mark-done":

	case "list":

	default:
		fmt.Println("Unknown command:", command)
	}
}
