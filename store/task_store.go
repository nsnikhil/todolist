package store

import (
	"todolist/domain"
)

type TaskStore interface {
	Add(task domain.Task) error
	Remove(id string, ids ...string) error
	Update(task domain.Task) error
	GetTasks(ids ...string) ([]domain.Task, error)
}


