package store

import (
	"github.com/jmoiron/sqlx"
	"todolist/util"
)

type Store struct {
	TaskStore
}

func NewStore(db *sqlx.DB) Store {
	util.DebugLog("[Store] [NewStore]", db)
	return Store{
		TaskStore: NewTaskStore(db),
	}
}

func (s Store) GetTodoListStore() TaskStore {
	util.DebugLog("[Store] [GetTodoListStore]")
	return s.TaskStore
}
