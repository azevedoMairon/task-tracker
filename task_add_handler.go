package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
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
		fmt.Println("error while loading tasks: ", err)
		return
	}

	nextID := getNextID(tasks)
	newTask := NewTask(nextID, description)

	tasks[strconv.Itoa(nextID)] = newTask

	if err := saveTasks(tasks); err != nil {
		fmt.Println("error while saving tasks: ", err)
		return
	}

	fmt.Printf("Task added succesfully (ID: %d)", nextID)
}

func loadTasks() (TaskMap, error) {
	file, err := os.Open(tasksFile)
	if err != nil {
		if os.IsNotExist(err) {
			return TaskMap{}, nil
		}
		return nil, err
	}
	defer file.Close()

	var tasks TaskMap
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&tasks); err != nil {
		return nil, err
	}
	return tasks, nil
}

func saveTasks(tasks TaskMap) error {
	file, err := os.Create(tasksFile)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(tasks)
}

func getNextID(tasks TaskMap) int {
	maxID := 0
	for idStr := range tasks {
		if id, err := strconv.Atoi(idStr); err == nil && id > maxID {
			maxID = id
		}
	}
	return maxID + 1
}
