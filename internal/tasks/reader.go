package tasks

import (
	"fmt"
	"log"
	"os"

	"github.com/azevedoMairon/task-tracker/internal/contracts"
	"github.com/azevedoMairon/task-tracker/internal/models"
)

type Reader struct {
	loader contracts.Loader
}

func NewReader(loader contracts.Loader) *Reader {
	return &Reader{
		loader: loader,
	}
}

func (r Reader) Read() {
	if len(os.Args) < 2 {
		log.Fatal("usage: task-cli list <status (optional)>")
	}

	if len(os.Args) < 3 {
		r.readAll()
	}

	r.readByStatus()
}

func (r Reader) readAll() {
	tasks, err := r.loader.Load()
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

func (r Reader) readByStatus() {
	statusStr := os.Args[2]
	status, err := parseStatus(statusStr)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	tasks, err := r.loader.Load()
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

func parseStatus(s string) (models.Status, error) {
	switch s {
	case "todo":
		return models.StatusTodo, nil
	case "in-progress":
		return models.StatusInProgress, nil
	case "done":
		return models.StatusDone, nil
	default:
		return -1, fmt.Errorf("invalid status: %s (use: todo, in-progress, done)", s)
	}
}
