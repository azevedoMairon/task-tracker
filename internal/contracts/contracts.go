package contracts

import "github.com/azevedoMairon/task-tracker/internal/models"

type Loader interface {
	Load() (models.TaskMap, error)
}

type Saver interface {
	Save(tasks models.TaskMap) error
}
