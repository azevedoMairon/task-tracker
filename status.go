package main

type Status int

const (
	StatusTodo Status = iota
	StatusInProgress
	StatusDone
)

func (s Status) String() string {
	switch s {
	case StatusTodo:
		return "todo"
	case StatusInProgress:
		return "in-progress"
	case StatusDone:
		return "done"
	default:
		return "unknown"
	}
}
