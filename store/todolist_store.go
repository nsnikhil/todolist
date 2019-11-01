package store

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"todolist/applogger"
	"todolist/constants"
	"todolist/domain"
)

type TodoListStoreInterface interface {
	Add(task domain.TaskInterface) error
	Remove(taskID string) error
	Update(task domain.TaskInterface) error
	GetTodoList() (domain.TodoListInterface, error)
	GetTask(taskID string) (domain.TaskInterface, error)
}

type TodoListStore struct {
	db *sqlx.DB
}

func NewTodoListStore(db *sqlx.DB) TodoListStoreInterface {
	applogger.Infof(constants.StoreNewTaskStore, fmt.Sprint("[TodoListStore] [NewTodoListStore]"))
	return TodoListStore{db: db}
}

func (tls TodoListStore) Add(task domain.TaskInterface) error {
	result, err := tls.db.Exec(constants.InsertIntoTodoListTable, task.GetID(), task.GetDescription(), task.GetStatus())
	if err != nil {
		applogger.Errorf(constants.ErrorFailedToInsertTask, fmt.Sprint("[TodoListStore] [Add]"), task, err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if rowsAffected == 0 || err != nil {
		applogger.Errorf(constants.ErrorRowsAffected, fmt.Sprint("[TodoListStore] [Add]"), task, err)
		return err
	}

	applogger.Infof(constants.SuccessfulInsertIntoDatabase, fmt.Sprint("[TodoListStore] [Add]"), task)
	return nil
}

func (tls TodoListStore) Remove(taskID string) error {
	result, err := tls.db.Exec(constants.DeleteFromTodoListTable, taskID)
	if err != nil {
		applogger.Errorf(constants.ErrorFailedToDeleteTask, fmt.Sprint("[TodoListStore] [Remove]"), taskID, err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if rowsAffected == 0 || err != nil {
		applogger.Errorf(constants.ErrorRowsAffectedDelete, fmt.Sprint("[TodoListStore] [Remove]"), taskID, err)
		return err
	}

	applogger.Infof(constants.SuccessfulDeleteFromDatabase, fmt.Sprint("[TodoListStore] [Remove]"), taskID)
	return nil
}

func (tls TodoListStore) Update(task domain.TaskInterface) error {
	var oldTask domain.Task
	err := tls.db.Get(&oldTask, constants.FindTaskInDatabase, task.GetID())
	if err != nil {
		applogger.Errorf(constants.ErrorTaskNotPresentInDatabase, fmt.Sprint("[TodoListStore] [Update]"), task, err)
		return err
	}

	result, err := tls.db.Exec(constants.UpdateTodoListTable, task.GetDescription(), task.GetStatus(), task.GetID())
	if err != nil {
		applogger.Errorf(constants.ErrorFailedToUpdateTask, fmt.Sprint("[TodoListStore] [Update]"), task, err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if rowsAffected == 0 || err != nil {
		applogger.Errorf(constants.ErrorRowsAffectedUpdate, fmt.Sprint("[TodoListStore] [Update]"), task, err)
		return err
	}

	applogger.Infof(constants.SuccessfulUpdateTaskFromDatabase, fmt.Sprint("[TodoListStore] [Update]"), oldTask, task)
	return nil
}

func (tls TodoListStore) GetTodoList() (domain.TodoListInterface, error) {
	var todoList domain.TodoList
	err := tls.db.Select(&todoList.Tasks, constants.GetAllTasksFromDatabase)
	if err != nil {
		applogger.Errorf(constants.ErrorFailedToGetAllTasksFromDatabase, fmt.Sprint("[TodoListStore] [GetTodoList]"), err)
		return nil, err
	}
	applogger.Infof(constants.SuccessfulGetAllTaskFromDatabase, fmt.Sprint("[TodoListStore] [GetTodoList]"), len(todoList.Tasks))
	return &todoList, nil
}

func (tls TodoListStore) GetTask(taskID string) (domain.TaskInterface, error) {
	var task domain.Task
	err := tls.db.Get(&task, constants.FindTaskInDatabase, taskID)
	if err != nil {
		applogger.Errorf(constants.ErrorTaskWithIDNotPresentInDatabase, fmt.Sprint("[TodoListStore] [GetTask]"), taskID, err)
		return nil, err
	}
	applogger.Infof(constants.SuccessfulGetTaskFromDatabase, fmt.Sprint("[TodoListStore] [GetTask]"), task)
	return &task, nil
}
