package main

import "time"

type Status int

const (
	StatusTodo Status = iota
	StatusInProgress
	StatusDone
)

type TaskMap map[string]Task

type Task struct {
	Description string    `json:"description"`
	Status      Status    `json:"status"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func NewTask(id int, desc string) Task {
	return Task{
		Description: desc,
		Status:      StatusTodo,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}

func (t *Task) SetDescription(desc string) {
	t.Description = desc
}

func (t *Task) SetStatus(status Status) {
	t.Status = status
}
