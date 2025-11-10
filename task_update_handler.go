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
	if len(os.Args) < 4 {
		log.Fatal("usage: task-cli update <task id> <new task description>")
	}

	taskID := os.Args[2]
	newDesc := os.Args[3]

	tasks, err := loadTasks()
	if err != nil {
		fmt.Println("error while loading tasks: ", err)
		return
	}

	task, exists := tasks[taskID]
	if !exists {
		fmt.Println("requested task does not exist")
		return
	}

	task.SetDescription(newDesc)
	tasks[taskID] = task

	if err := saveTasks(tasks); err != nil {
		fmt.Println("error while saving tasks: ", err)
		return
	}

	fmt.Printf("task %s updated succesfully", taskID)
}
