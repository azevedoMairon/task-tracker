package main

import (
	"fmt"
	"log"
	"os"
)

type TaskUpdateHandler struct {
	loader ILoader
	saver  ISaver
}

func NewTaskUpdateHandler(loader ILoader, saver ISaver) *TaskUpdateHandler {
	return &TaskUpdateHandler{
		loader: loader,
		saver:  saver,
	}
}

func (h TaskUpdateHandler) HandleUpdateDesc() {
	if len(os.Args) < 4 {
		log.Fatal("usage: task-cli update <task id> <new task description>")
	}

	taskID := os.Args[2]
	newDesc := os.Args[3]

	tasks, err := h.loader.Load()
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

	if err := h.saver.Save(tasks); err != nil {
		fmt.Println("error while saving tasks: ", err)
		return
	}

	fmt.Printf("task %s updated succesfully", taskID)
}

func (h TaskUpdateHandler) HandleMarkInProgress() {
	if len(os.Args) < 3 {
		log.Fatal("usage: task-cli mark-in-progress <task id>")
	}
	taskID := os.Args[2]

	tasks, err := h.loader.Load()
	if err != nil {
		fmt.Println("error while loading tasks: ", err)
		return
	}

	task, exists := tasks[taskID]
	if !exists {
		fmt.Println("requested task does not exist")
		return
	}

	task.SetStatus(StatusInProgress)
	tasks[taskID] = task

	if err := h.saver.Save(tasks); err != nil {
		fmt.Println("error while saving tasks: ", err)
		return
	}

	fmt.Printf("task %s is now in progress ;)", taskID)
}

func (h TaskUpdateHandler) HandleMarkDone() {
	if len(os.Args) < 3 {
		log.Fatal("usage: task-cli mark-done <task id>")
	}
	taskID := os.Args[2]

	tasks, err := h.loader.Load()
	if err != nil {
		fmt.Println("error while loading tasks: ", err)
		return
	}

	task, exists := tasks[taskID]
	if !exists {
		fmt.Println("requested task does not exist")
		return
	}

	task.SetStatus(StatusDone)
	tasks[taskID] = task

	if err := h.saver.Save(tasks); err != nil {
		fmt.Println("error while saving tasks: ", err)
		return
	}

	fmt.Printf("task %s is now done ;)", taskID)
}
