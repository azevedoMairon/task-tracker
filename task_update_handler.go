package main

import (
	"fmt"
	"log"
	"os"
)

type TaskUpdateHandler struct {
}

func NewTaskUpdateHandler() *TaskUpdateHandler {
	return &TaskUpdateHandler{}
}

func (h TaskUpdateHandler) Handle() {
	if len(os.Args) < 3 {
		log.Fatal("usage: task-cli update <task id> <new task description>")
		return
	}
	description := os.Args[2]

	tasks, err := loadTasks()
	if err != nil {
		fmt.Println("Error while loading tasks: ", err)
		return
	}

	nextID := getNextID(tasks)

	newTask := NewTask(nextID, description)

	tasks = append(tasks, newTask)

	if err := saveTasks(tasks); err != nil {
		fmt.Println("Error while saving tasks: ", err)
		return
	}

	fmt.Printf("Task added succesfully (ID: %d)", nextID)
}
