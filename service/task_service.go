package service

import (
	"todolist/domain"
)

type TaskService interface {
	Add(task domain.Task) error
	Remove(id string, ids ...string) error
	Update(task domain.Task) error
	GetTasks(ids ...string) ([]domain.Task, error)
}
