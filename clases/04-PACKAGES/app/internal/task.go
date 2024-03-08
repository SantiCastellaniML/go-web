package internal

import "errors"

type Task struct {
	ID          int
	Title       string
	Description string
	Done        bool
}

var (
	// errors that my interface could return (related to a repository)
	ErrTaskNotFound     = errors.New("task not found")
	ErrTaskDuplicated   = errors.New("task already exists")
	ErrTaskConflict     = errors.New("task can't be processed")
	ErrTaskInvalidField = errors.New("task is invalid")
)

// TaskRepository is an interface that represents a task repository
type TaskRepository interface {
	// Create creates a new task
	Save(task *Task) (err error)
}

var (
	//errors related to a service.
	ErrTaskSerivce = errors.New("task service can't be processed")
)

type TaskService interface {
	Save(Task *Task) (err error)
}
