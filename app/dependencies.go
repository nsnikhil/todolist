package app

import (
	"todolist/domain"
	"todolist/service"
)

type Dependencies struct {
	service.Service
	domain.TaskFactory
}
