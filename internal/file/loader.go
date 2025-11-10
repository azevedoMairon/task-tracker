package file

import (
	"encoding/json"
	"os"

	"github.com/azevedoMairon/task-tracker/internal/models"
)

type Loader struct {
	file string
}

func NewLoader(file string) *Loader {
	return &Loader{
		file: file,
	}
}

func (l Loader) Load() (models.TaskMap, error) {
	file, err := os.Open(l.file)
	if err != nil {
		if os.IsNotExist(err) {
			return models.TaskMap{}, nil
		}
		return nil, err
	}
	defer file.Close()

	var tasks models.TaskMap
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&tasks); err != nil {
		return nil, err
	}
	return tasks, nil
}
