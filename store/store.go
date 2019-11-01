package store

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"todolist/applogger"
	"todolist/constants"
)

type Store struct {
	todoListStore TodoListStoreInterface
}

func NewStore(db *sqlx.DB) Store {
	applogger.Infof(constants.StoreNewStore, fmt.Sprint("[Store] [NewStore]"))
	return Store{
		todoListStore: NewTodoListStore(db),
	}
}

func (s Store) GetTodoListStore() TodoListStoreInterface {
	applogger.Infof(constants.StoreGetTodoListStore, fmt.Sprint("[Store] [GetTodoListStore]"), s.todoListStore)
	return s.todoListStore
}
