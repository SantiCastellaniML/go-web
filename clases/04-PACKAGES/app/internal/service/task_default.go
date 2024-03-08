package service

import "go-packages/internal"

type TaskDefault struct {
	// rp is the task repository
	rp internal.TaskRepository
	// external services / apis
	// ...
}

func NewDefaultTask(rp internal.TaskRepository) *TaskDefault {
	return &TaskDefault{
		rp: rp,
	}
}

func (t *TaskDefault) Save(task *internal.Task) (err error) {
	err = t.rp.Save(task)
	if err != nil {
		//we should check if it's a service or repository error.
		//i shouldn't identify the specific repository's error here, i shouldn't validate errors from MySQL for example. That's what generic errors are for.
	}
	return
}
