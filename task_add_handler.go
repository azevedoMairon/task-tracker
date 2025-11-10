package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

type TaskAddHandler struct {
	loader ILoader
	saver  ISaver
}

func NewTaskAddHandler(loader ILoader, saver ISaver) *TaskAddHandler {
	return &TaskAddHandler{
		loader: loader,
		saver:  saver,
	}
}

func (h TaskAddHandler) Handle() {
	if len(os.Args) < 3 {
		log.Fatal("usage: task-cli add <task description>")
	}
	description := os.Args[2]

	tasks, err := h.loader.Load()
	if err != nil {
		fmt.Println("error while loading tasks: ", err)
		return
	}

	nextID := h.getNextID(tasks)
	newTask := NewTask(nextID, description)

	tasks[strconv.Itoa(nextID)] = newTask

	if err := h.saver.Save(tasks); err != nil {
		fmt.Println("error while saving tasks: ", err)
		return
	}

	fmt.Printf("Task added succesfully (ID: %d)", nextID)
}

func (h TaskAddHandler) getNextID(tasks TaskMap) int {
	maxID := 0
	for idStr := range tasks {
		if id, err := strconv.Atoi(idStr); err == nil && id > maxID {
			maxID = id
		}
	}
	return maxID + 1
}
