package file

import (
	"encoding/json"
	"os"

	"github.com/azevedoMairon/task-tracker/internal/models"
)

type Saver struct {
	file string
}

func NewSaver(file string) *Saver {
	return &Saver{
		file: file,
	}
}

func (s Saver) Save(tasks models.TaskMap) error {
	file, err := os.Create(s.file)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(tasks)
}
