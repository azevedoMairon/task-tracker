package tasks

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/azevedoMairon/task-tracker/internal/contracts"
	"github.com/azevedoMairon/task-tracker/internal/models"
)

type Creator struct {
	loader contracts.Loader
	saver  contracts.Saver
}

func NewCreator(loader contracts.Loader, saver contracts.Saver) *Creator {
	return &Creator{
		loader: loader,
		saver:  saver,
	}
}

func (c Creator) Create() {
	if len(os.Args) < 3 {
		log.Fatal("usage: task-cli add <task description>")
	}
	description := os.Args[2]

	tasks, err := c.loader.Load()
	if err != nil {
		fmt.Println("error while loading tasks: ", err)
		return
	}

	nextID := c.getNextID(tasks)
	newTask := models.NewTask(nextID, description)

	tasks[strconv.Itoa(nextID)] = newTask

	if err := c.saver.Save(tasks); err != nil {
		fmt.Println("error while saving tasks: ", err)
		return
	}

	fmt.Printf("Task added succesfully (ID: %d)", nextID)
}

func (c Creator) getNextID(tasks models.TaskMap) int {
	maxID := 0
	for idStr := range tasks {
		if id, err := strconv.Atoi(idStr); err == nil && id > maxID {
			maxID = id
		}
	}
	return maxID + 1
}
