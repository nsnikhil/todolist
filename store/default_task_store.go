package store

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"strings"
	"todolist/domain"
	"todolist/util"
)

const (
	insertTask = `INSERT INTO todolist (id, title, description, status, tags) values ($1, $2, $3, $4, $5);`

	deleteTask = `DELETE FROM todolist where id in ($1)`

	updateTask = `UPDATE todolist set title = $1, description = $2, status = $3, tags = $4 where id = $5`

	findTask = `SELECT id, title, description, status, tags FROM todolist where id = $1`

	getTasks = `Select id, title, description, status, tags FROM todolist where id in ($1)`

	getAllTasks = `Select id, title, description, status, tags FROM todolist`
)

type DefaultTaskStore struct {
	db *sqlx.DB
}

func NewTaskStore(db *sqlx.DB) TaskStore {
	util.DebugLog("[TaskStore] [NewTaskStore]")
	return DefaultTaskStore{db: db}
}

func (tls DefaultTaskStore) Add(task domain.Task) error {
	result, err := tls.db.Exec(insertTask, task.GetID(), task.GetTitle(), task.GetDescription(), task.GetStatus(), pq.Array(task.GetTags()))
	if err != nil {
		return util.LogAndGetError("[DefaultTaskStore] [Add] [Exec]", err)
	}

	rowsAffected, err := result.RowsAffected()
	if rowsAffected == 0 || err != nil {
		return util.LogAndGetError("[DefaultTaskStore] [Add] [RowsAffected]", err)
	}

	util.DebugLog("[DefaultTaskStore] [Add]", task)
	return nil
}

func (tls DefaultTaskStore) Remove(id string, ids ...string) error {
	result, err := tls.db.Exec(deleteTask, toArray(append(ids, id)...))
	if err != nil {
		return util.LogAndGetError("[DefaultTaskStore] [Remove] [Exec]", err)
	}

	rowsAffected, err := result.RowsAffected()
	if rowsAffected == 0 || err != nil {
		return util.LogAndGetError("[DefaultTaskStore] [Remove] [RowsAffected]", err)
	}

	util.DebugLog("[DefaultTaskStore] [Remove]", ids)
	return nil
}

func toArray(data ...string) string {
	return "{" + strings.Join(data, ",") + "}"
}

func (tls DefaultTaskStore) Update(task domain.Task) error {
	var oldTask domain.DefaultTask

	if err := tls.db.Get(&oldTask, findTask, task.GetID()); err != nil {
		return util.LogAndGetError("[DefaultTaskStore] [Update] [Get]", err)
	}

	result, err := tls.db.Exec(updateTask, task.GetTitle(), task.GetDescription(), task.GetStatus(), pq.Array(task.GetTags()), task.GetID())
	if err != nil {
		return util.LogAndGetError("[DefaultTaskStore] [Update] [Exec]", err)
	}

	rowsAffected, err := result.RowsAffected()
	if rowsAffected == 0 || err != nil {
		return util.LogAndGetError("[DefaultTaskStore] [Update] [RowsAffected]", err)
	}

	util.DebugLog("[DefaultTaskStore] [Remove]", task)
	return nil
}

func (tls DefaultTaskStore) GetTasks(ids ...string) ([]domain.Task, error) {
	var defaultTasks []*domain.DefaultTask

	if len(ids) == 0 {
		if err := tls.db.Select(&defaultTasks, getAllTasks); err != nil {
			return nil, util.LogAndGetError("[DefaultTaskStore] [GetTasks] [Select] [getAllTasks]", err)
		}
	} else {
		if err := tls.db.Select(&defaultTasks, getTasks, toArray(ids...)); err != nil {
			return nil, util.LogAndGetError("[DefaultTaskStore] [GetTasks] [Select] [getTasks]", err)
		}
	}

	if len(defaultTasks) == 0 {
		return nil, util.LogAndGetError("[DefaultTaskStore] [len]", fmt.Errorf("no task found for : %v", ids))
	}

	tasks := make([]domain.Task, 0)
	for _, df := range defaultTasks {
		tasks = append(tasks, df)
	}

	util.DebugLog("[DefaultTaskStore] [GetTasks]", ids)
	return tasks, nil
}
