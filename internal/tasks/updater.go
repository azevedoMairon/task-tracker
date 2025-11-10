package tasks

import (
	"fmt"
	"log"
	"os"

	"github.com/azevedoMairon/task-tracker/internal/contracts"
	"github.com/azevedoMairon/task-tracker/internal/models"
)

type Updater struct {
	loader contracts.Loader
	saver  contracts.Saver
}

func NewUpdater(loader contracts.Loader, saver contracts.Saver) *Updater {
	return &Updater{
		loader: loader,
		saver:  saver,
	}
}

func (u Updater) Update() {
	if len(os.Args) < 4 {
		log.Fatal("usage: task-cli update <task id> <new task description>")
	}

	taskID := os.Args[2]
	newDesc := os.Args[3]

	tasks, err := u.loader.Load()
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

	if err := u.saver.Save(tasks); err != nil {
		fmt.Println("error while saving tasks: ", err)
		return
	}

	fmt.Printf("task %s updated succesfully", taskID)
}

func (u Updater) MarkInProgress() {
	if len(os.Args) < 3 {
		log.Fatal("usage: task-cli mark-in-progress <task id>")
	}
	taskID := os.Args[2]

	tasks, err := u.loader.Load()
	if err != nil {
		fmt.Println("error while loading tasks: ", err)
		return
	}

	task, exists := tasks[taskID]
	if !exists {
		fmt.Println("requested task does not exist")
		return
	}

	task.SetStatus(models.StatusInProgress)
	tasks[taskID] = task

	if err := u.saver.Save(tasks); err != nil {
		fmt.Println("error while saving tasks: ", err)
		return
	}

	fmt.Printf("task %s is now in progress ;)", taskID)
}

func (u Updater) MarkDone() {
	if len(os.Args) < 3 {
		log.Fatal("usage: task-cli mark-done <task id>")
	}
	taskID := os.Args[2]

	tasks, err := u.loader.Load()
	if err != nil {
		fmt.Println("error while loading tasks: ", err)
		return
	}

	task, exists := tasks[taskID]
	if !exists {
		fmt.Println("requested task does not exist")
		return
	}

	task.SetStatus(models.StatusDone)
	tasks[taskID] = task

	if err := u.saver.Save(tasks); err != nil {
		fmt.Println("error while saving tasks: ", err)
		return
	}

	fmt.Printf("task %s is now done ;)", taskID)
}
