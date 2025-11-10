package tasks

import (
	"fmt"
	"log"
	"os"

	"github.com/azevedoMairon/task-tracker/internal/contracts"
)

type Deleter struct {
	loader contracts.Loader
	saver  contracts.Saver
}

func NewDeleter(loader contracts.Loader, saver contracts.Saver) *Deleter {
	return &Deleter{
		loader: loader,
		saver:  saver,
	}
}

func (d Deleter) Delete() {
	if len(os.Args) < 3 {
		log.Fatal("usage: task-cli delete <task id>")
	}
	taskID := os.Args[2]

	tasks, err := d.loader.Load()
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

	if err := d.saver.Save(tasks); err != nil {
		fmt.Println("error while saving tasks: ", err)
		return
	}

	fmt.Printf("task with id %s deleted succesfully", taskID)
}
