package domain

import (
	"fmt"
	"todolist/apperror"
	"todolist/applogger"
	"todolist/constants"
)

type TodoListInterface interface {
	Add(task *Task)
	Remove(task *Task) error
}

type TodoList struct {
	Tasks []*Task
}

func NewTodoList(task *Task, tasks ...*Task) *TodoList {
	todoListTasks := append(tasks, task)
	applogger.Infof(constants.TodoListCreationSuccessMessage, fmt.Sprint("[TodoList] [NewTodoList]"), todoListTasks)
	return &TodoList{Tasks: todoListTasks}
}

func (tl *TodoList) Add(task *Task) {
	tl.Tasks = append(tl.Tasks, task)
	applogger.Infof(constants.TaskAddSuccessMessage, fmt.Sprint("[TodoList] [Add]"), task.Description)
}

func (tl *TodoList) Remove(task *Task) error {
	for i, tk := range tl.Tasks {
		if tk == task {
			tl.Tasks = append(tl.Tasks[:i], tl.Tasks[i+1:]...)
			applogger.Infof(constants.TaskRemoveSuccessMessage, fmt.Sprint("[TodoList] [Remove]"), task.Description)
			return nil
		}
	}
	applogger.Errorf(constants.TaskRemoveFailedErrorMessage, fmt.Sprint("[TodoList] [Remove]"), task.Description)
	return apperror.NewTaskRemoveFailedError(task.Description)
}
