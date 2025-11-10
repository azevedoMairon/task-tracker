package main

import (
	"fmt"
	"log"
	"os"
)

type TaskListHandler struct {
	loader ILoader
}

func NewTaskListHandler(loader ILoader) *TaskListHandler {
	return &TaskListHandler{
		loader: loader,
	}
}

func (h TaskListHandler) Handle() {
	if len(os.Args) < 2 {
		log.Fatal("usage: task-cli list <status>")
	}

	if len(os.Args) < 3 {
		h.handleListAll()
	}

	h.handleListByStatus()
}

func (h TaskListHandler) handleListAll() {
	tasks, err := h.loader.Load()
	if err != nil {
		fmt.Println("error while loading tasks:", err)
		return
	}

	if len(tasks) == 0 {
		fmt.Println("no tasks found")
		return
	}

	for id, task := range tasks {
		fmt.Printf("Task %s: %s [%s] | %s\n", id, task.Description, task.Status.String(), task.CreatedAt.Format("2006-01-02"))
	}
}

func (h TaskListHandler) handleListByStatus() {
	statusStr := os.Args[2]
	status, err := parseStatus(statusStr)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	tasks, err := h.loader.Load()
	if err != nil {
		fmt.Println("error while loading tasks:", err)
		return
	}

	found := false
	for id, task := range tasks {
		if task.Status == status {
			found = true
			fmt.Printf("Task %s: %s [%s] | %s\n", id, task.Description, task.Status.String(), task.CreatedAt.Format("2006-01-02"))
		}
	}

	if !found {
		fmt.Println("no tasks found with this status")
	}
}

func parseStatus(statusStr string) (Status, error) {
	switch statusStr {
	case "todo":
		return StatusTodo, nil
	case "in-progress":
		return StatusInProgress, nil
	case "done":
		return StatusDone, nil
	default:
		return -1, fmt.Errorf("invalid status: %s (use: todo, in-progress, done)", statusStr)
	}
}
