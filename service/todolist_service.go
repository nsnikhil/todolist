package service

import (
	"fmt"
	"todolist/applogger"
	"todolist/constants"
	"todolist/domain"
	"todolist/store"
)

type TodoListServiceInterface interface {
	Add(task domain.TaskInterface) error
	Remove(taskID string) error
	Update(task domain.TaskInterface) error
	GetTodoList() (domain.TodoListInterface, error)
	GetTask(taskID string) (domain.TaskInterface, error)
}

type TodoListService struct {
	todoListStore store.TodoListStoreInterface
}

func NewTodoListService(todoListStore store.TodoListStoreInterface) TodoListServiceInterface {
	applogger.Infof(constants.TodoListServiceNew, fmt.Sprint("[TodoListService] [NewTodoListService]"), todoListStore)
	return TodoListService{todoListStore: todoListStore}
}

func (tls TodoListService) Add(task domain.TaskInterface) error {
	if err := tls.todoListStore.Add(task); err != nil {
		applogger.Errorf(constants.ErrorFailedToAddTask, fmt.Sprint("[TodoListService] [Add]"), task, err)
		return err
	}
	applogger.Infof(constants.SuccessfulAddTask, fmt.Sprint("[TodoListService] [Add]"), task)
	return nil
}

func (tls TodoListService) Remove(taskID string) error {
	if err := tls.todoListStore.Remove(taskID); err != nil {
		applogger.Errorf(constants.ErrorFailedToRemoveTask, fmt.Sprint("[TodoListService] [Remove]"), taskID, err)
		return err
	}
	applogger.Infof(constants.SuccessfulRemoveTask, fmt.Sprint("[TodoListService] [Remove]"), taskID)
	return nil
}

func (tls TodoListService) Update(task domain.TaskInterface) error {
	if err := tls.todoListStore.Update(task); err != nil {
		applogger.Errorf(constants.ErrorTaskUpdateFailed, fmt.Sprint("[TodoListService] [Update]"), task, err)
		return err
	}
	applogger.Infof(constants.SuccessfulUpdateTask, fmt.Sprint("[TodoListService] [Update]"), task)
	return nil
}

func (tls TodoListService) GetTodoList() (domain.TodoListInterface, error) {
	todoList, err := tls.todoListStore.GetTodoList()
	if err != nil {
		applogger.Errorf(constants.ErrorFailedToGetAllTasks, fmt.Sprint("[TodoListService] [GetTodoList]"), err)
		return nil, err
	}
	applogger.Infof(constants.SuccessfulGetAllTask, fmt.Sprint("[TodoListService] [GetTodoList]"), todoList)
	return todoList, nil
}

func (tls TodoListService) GetTask(taskID string) (domain.TaskInterface, error) {
	task, err := tls.todoListStore.GetTask(taskID)
	if err != nil {
		applogger.Errorf(constants.ErrorTaskNotFound, fmt.Sprint("[TodoListService] [GetTask]"), taskID, err)
		return nil, err
	}
	applogger.Infof(constants.SuccessfulGetTask, fmt.Sprint("[TodoListService] [GetTask]"), task)
	return task, err
}
