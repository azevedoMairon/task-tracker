package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type TaskAddHandler struct {
}

func NewTaskAddHandler() *TaskAddHandler {
	return &TaskAddHandler{}
}

func (h TaskAddHandler) Handle() {
	if len(os.Args) < 3 {
		log.Fatal("usage: task-cli add <task description>")
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

func loadTasks() ([]Task, error) {
	file, err := os.Open(tasksFile)
	if err != nil {
		if os.IsNotExist(err) {
			return []Task{}, nil
		}
		return nil, err
	}
	defer file.Close()

	var taskWrapper TaskWrapper
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&taskWrapper); err != nil {
		return nil, err
	}
	return taskWrapper.Tasks, nil
}

func saveTasks(tasks []Task) error {
	file, err := os.Create(tasksFile)
	if err != nil {
		return err
	}
	defer file.Close()

	taskWrapper := NewTaskWrapper(tasks)
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(taskWrapper)
}

func getNextID(tasks []Task) int {
	maxID := 0
	for _, task := range tasks {
		if task.ID > maxID {
			maxID = task.ID
		}
	}
	return maxID + 1
}
