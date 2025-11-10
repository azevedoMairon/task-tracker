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

type AddTaskHandler struct {
}

func NewAddTaskHandler() *AddTaskHandler {
	return &AddTaskHandler{}
}

func (h AddTaskHandler) Handle() {
	if len(os.Args) < 3 {
		log.Fatal("usage: task-cli add <task name>")
	}

	if exists, err := h.fileExists(IDsFile); !exists {
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
}

func (h AddTaskHandler) fileExists(f string) (bool, error) {
	if _, err := os.Stat(f); errors.Is(err, os.ErrNotExist) {
		return false, nil
	} else if err != nil {
		return false, err
	} else {
		return true, nil
	}
}
