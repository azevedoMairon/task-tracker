package main

type TaskWrapper struct {
	Tasks []Task
}

func NewTaskWrapper(tasks []Task) TaskWrapper {
	return TaskWrapper{
		Tasks: tasks,
	}
}
