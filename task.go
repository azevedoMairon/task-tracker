package main

import "time"

type Status int

const (
	StatusTodo Status = iota
	StatusInProgress
	StatusDone
)

type Task struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	Status      Status    `json:"status"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func NewTask(id int, desc string) Task {
	return Task{
		ID:          id,
		Description: desc,
		Status:      StatusTodo,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}
