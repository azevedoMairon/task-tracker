package main

import (
	"encoding/json"
	"os"
)

type ILoader interface {
	Load() (TaskMap, error)
}

type Loader struct{}

func NewLoader() ILoader {
	return &Loader{}
}

func (l Loader) Load() (TaskMap, error) {
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
