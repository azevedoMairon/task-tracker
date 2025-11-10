package main

import (
	"encoding/json"
	"os"
)

type ISaver interface {
	Save(tasks TaskMap) error
}

type Saver struct{}

func NewSaver() ISaver {
	return &Saver{}
}

func (s Saver) Save(tasks TaskMap) error {
	file, err := os.Create(tasksFile)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(tasks)
}
