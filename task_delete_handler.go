package main

import (
	"fmt"
	"log"
	"os"
)

type TaskDeleteHandler struct {
}

func NewTaskDeleteHandler() *TaskDeleteHandler {
	return &TaskDeleteHandler{}
}

func (h TaskDeleteHandler) Handle() {
	if len(os.Args) < 3 {
		log.Fatal("usage: task-cli delete <task id>")
	}
	taskID := os.Args[2]

	tasks, err := loadTasks()
	if err != nil {
		fmt.Println("error while loading tasks: ", err)
		return
	}

	_, exists := tasks[taskID]
	if !exists {
		fmt.Println("requested task does not exist")
		return
	}

	delete(tasks, taskID)

	if err := saveTasks(tasks); err != nil {
		fmt.Println("error while saving tasks: ", err)
		return
	}

	fmt.Printf("task with id %s deleted succesfully", taskID)
}
