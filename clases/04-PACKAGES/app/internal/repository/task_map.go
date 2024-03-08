package repository

import "go-packages/internal"

type TaskMap struct {
	// db is a map that stores the tasks
	// -key: id
	// -value: task
	db map[int]internal.Task

	lastID int
}

func NewTaskMap(db map[int]internal.Task, lastID int) *TaskMap {
	if db == nil {
		db = make(map[int]internal.Task)
	}

	if lastID == 0 {
		lastID = 0
	}

	return &TaskMap{
		db:     db,
		lastID: lastID,
	}
}
