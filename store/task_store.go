package store

import (
	"todolist/domain"
)

type TaskStore interface {
	Add(task domain.Task) (string, error)
	Remove(id string, ids ...string) (int64, error)
	Update(task domain.Task) (int64, error)
	GetTasks(ids ...string) ([]domain.Task, error)
}
