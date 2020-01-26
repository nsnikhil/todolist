package store

import (
	"errors"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/magiconair/properties/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"todolist/config"
	"todolist/domain"
)

func cleanDB(db *sqlx.DB, t *testing.T) {
	_, err := db.Exec("TRUNCATE TABLE todolist")
	require.NoError(t, err)
}

func TestCreateNewTaskStore(t *testing.T) {
	err := config.Load()
	require.NoError(t, err)

	testCases := []struct {
		name                  string
		actualTodoListStore   TaskStore
		expectedTodoListStore TaskStore
	}{
		{
			name:                  "test create new todo list store",
			actualTodoListStore:   NewTaskStore(&sqlx.DB{}),
			expectedTodoListStore: DefaultTaskStore{db: &sqlx.DB{}},
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, testCase.expectedTodoListStore, testCase.actualTodoListStore)
		})
	}
}

func TestAddTaskInDB(t *testing.T) {
	err := config.Load()
	require.NoError(t, err)

	testCases := []struct {
		name          string
		actualTasks   func() []domain.Task
		expectedTasks func() []domain.Task
	}{
		{
			name: "store buy groceries task in db",
			actualTasks: func() []domain.Task {
				dbHandle := NewDBHandler(config.GetDatabaseConfig())
				db, err := dbHandle.GetDB()
				require.NoError(t, err)
				cleanDB(db, t)

				taskStore := NewTaskStore(db)

				task, err := domain.NewTaskFactory().Create("buy groceries", "", false)
				require.NoError(t, err)

				err = taskStore.Add(task)
				require.NoError(t, err)

				var defaultTasks []*domain.DefaultTask

				err = db.Select(&defaultTasks, "SELECT * FROM todolist")
				require.NoError(t, err)

				tasks := make([]domain.Task, 0)
				for _, df := range defaultTasks {
					tasks = append(tasks, df)
				}

				cleanDB(db, t)
				return tasks
			},
			expectedTasks: func() []domain.Task {
				task, err := domain.NewTaskFactory().Create("buy groceries", "", false)
				require.NoError(t, err)

				return []domain.Task{task}
			},
		},
		{
			name: "store buy groceries and read xyz task in db",
			actualTasks: func() []domain.Task {
				dbHandle := NewDBHandler(config.GetDatabaseConfig())
				db, err := dbHandle.GetDB()
				require.NoError(t, err)
				taskStore := NewTaskStore(db)
				cleanDB(db, t)

				buyGroceries, err := domain.NewTaskFactory().Create("buy groceries", "", false)
				require.NoError(t, err)

				readXyz, err := domain.NewTaskFactory().Create("read xyz", "", false)
				require.NoError(t, err)

				err = taskStore.Add(buyGroceries)
				require.NoError(t, err)

				err = taskStore.Add(readXyz)
				require.NoError(t, err)

				var defaultTasks []*domain.DefaultTask

				err = db.Select(&defaultTasks, "SELECT * FROM todolist")
				require.NoError(t, err)

				tasks := make([]domain.Task, 0)
				for _, df := range defaultTasks {
					tasks = append(tasks, df)
				}

				cleanDB(db, t)
				return tasks
			},
			expectedTasks: func() []domain.Task {
				buyGroceries, err := domain.NewTaskFactory().Create("buy groceries", "", false)
				require.NoError(t, err)

				readXyz, err := domain.NewTaskFactory().Create("read xyz", "", false)
				require.NoError(t, err)

				return []domain.Task{buyGroceries, readXyz}
			},
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			expectedTasks := testCase.expectedTasks()
			actualTasks := testCase.actualTasks()
			assert.Equal(t, len(expectedTasks), len(actualTasks))
			l := len(expectedTasks)
			for i := 0; i < l; i++ {
				assert.Equal(t, expectedTasks[i].GetTitle(), actualTasks[i].GetTitle())
				assert.Equal(t, expectedTasks[i].GetDescription(), actualTasks[i].GetDescription())
				assert.Equal(t, expectedTasks[i].GetStatus(), actualTasks[i].GetStatus())
				assert.Equal(t, expectedTasks[i].GetTags(), actualTasks[i].GetTags())
			}
		})
	}
}

func TestRemoveTaskFromDB(t *testing.T) {
	err := config.Load()
	require.NoError(t, err)

	testCases := []struct {
		name          string
		actualTasks   func() []domain.Task
		expectedTasks func() []domain.Task
	}{
		{
			name: "remove buy groceries task from db",
			actualTasks: func() []domain.Task {
				dbHandle := NewDBHandler(config.GetDatabaseConfig())
				db, err := dbHandle.GetDB()
				require.NoError(t, err)

				taskStore := NewTaskStore(db)
				cleanDB(db, t)

				task, err := domain.NewTaskFactory().Create("buy groceries", "", false)
				require.NoError(t, err)

				err = taskStore.Add(task)
				require.NoError(t, err)

				err = taskStore.Remove(task.GetID())
				require.NoError(t, err)

				var defaultTasks []*domain.DefaultTask

				err = db.Select(&defaultTasks, "SELECT * FROM todolist")
				require.NoError(t, err)

				tasks := make([]domain.Task, 0)
				for _, df := range defaultTasks {
					tasks = append(tasks, df)
				}

				cleanDB(db, t)
				return tasks
			},
			expectedTasks: func() []domain.Task {
				return []domain.Task{}
			},
		},
		{
			name: "remove buy groceries and read xyz task from db",
			actualTasks: func() []domain.Task {
				dbHandle := NewDBHandler(config.GetDatabaseConfig())
				db, err := dbHandle.GetDB()
				require.NoError(t, err)
				taskStore := NewTaskStore(db)
				cleanDB(db, t)

				buyGroceries, err := domain.NewTaskFactory().Create("buy groceries", "", false)
				require.NoError(t, err)

				readXyz, err := domain.NewTaskFactory().Create("read xyz", "", false)
				require.NoError(t, err)

				err = taskStore.Add(buyGroceries)
				require.NoError(t, err)

				err = taskStore.Add(readXyz)
				require.NoError(t, err)

				err = taskStore.Remove(buyGroceries.GetID())
				require.NoError(t, err)

				err = taskStore.Remove(readXyz.GetID())
				require.NoError(t, err)

				var defaultTasks []*domain.DefaultTask

				err = db.Select(&defaultTasks, "SELECT * FROM todolist")
				require.NoError(t, err)

				tasks := make([]domain.Task, 0)
				for _, df := range defaultTasks {
					tasks = append(tasks, df)
				}

				cleanDB(db, t)
				return tasks
			},
			expectedTasks: func() []domain.Task {
				return []domain.Task{}
			},
		},
		{
			name: "remove buy groceries and read xyz tasks from db",
			actualTasks: func() []domain.Task {
				dbHandle := NewDBHandler(config.GetDatabaseConfig())
				db, err := dbHandle.GetDB()
				require.NoError(t, err)
				taskStore := NewTaskStore(db)
				cleanDB(db, t)

				buyGroceries, err := domain.NewTaskFactory().Create("buy groceries", "", false)
				require.NoError(t, err)

				readXyz, err := domain.NewTaskFactory().Create("read xyz", "", false)
				require.NoError(t, err)

				callPerson, err := domain.NewTaskFactory().Create("call person", "", false)
				require.NoError(t, err)

				err = taskStore.Add(buyGroceries)
				require.NoError(t, err)

				err = taskStore.Add(readXyz)
				require.NoError(t, err)

				err = taskStore.Add(callPerson)
				require.NoError(t, err)

				err = taskStore.Remove(buyGroceries.GetID())
				require.NoError(t, err)

				err = taskStore.Remove(readXyz.GetID())
				require.NoError(t, err)

				var defaultTasks []*domain.DefaultTask

				err = db.Select(&defaultTasks, "SELECT * FROM todolist")
				require.NoError(t, err)

				tasks := make([]domain.Task, 0)
				for _, df := range defaultTasks {
					tasks = append(tasks, df)
				}

				cleanDB(db, t)
				return tasks
			},
			expectedTasks: func() []domain.Task {
				callPerson, err := domain.NewTaskFactory().Create("call person", "", false)
				require.NoError(t, err)

				return []domain.Task{callPerson}
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			expectedTasks := testCase.expectedTasks()
			actualTasks := testCase.actualTasks()
			assert.Equal(t, len(expectedTasks), len(actualTasks))
			l := len(expectedTasks)
			for i := 0; i < l; i++ {
				assert.Equal(t, expectedTasks[i].GetTitle(), actualTasks[i].GetTitle())
				assert.Equal(t, expectedTasks[i].GetDescription(), actualTasks[i].GetDescription())
				assert.Equal(t, expectedTasks[i].GetStatus(), actualTasks[i].GetStatus())
				assert.Equal(t, expectedTasks[i].GetTags(), actualTasks[i].GetTags())
			}
		})
	}
}

func TestUpdateTaskInDB(t *testing.T) {
	err := config.Load()
	require.NoError(t, err)

	testCases := []struct {
		name          string
		actualTasks   func() []domain.Task
		expectedTasks func() []domain.Task
	}{
		{
			name: "update buy bread to buy groceries",
			actualTasks: func() []domain.Task {
				dbHandle := NewDBHandler(config.GetDatabaseConfig())
				db, err := dbHandle.GetDB()
				require.NoError(t, err)
				taskStore := NewTaskStore(db)
				cleanDB(db, t)

				task, err := domain.NewTaskFactory().Create("buy bread", "", false)
				require.NoError(t, err)

				err = taskStore.Add(task)
				require.NoError(t, err)

				task.UpdateTitle("buy groceries")

				err = taskStore.Update(task)
				require.NoError(t, err)

				var defaultTasks []*domain.DefaultTask

				err = db.Select(&defaultTasks, "SELECT * FROM todolist")
				require.NoError(t, err)

				tasks := make([]domain.Task, 0)
				for _, df := range defaultTasks {
					tasks = append(tasks, df)
				}

				cleanDB(db, t)
				return tasks
			},
			expectedTasks: func() []domain.Task {
				task, err := domain.NewTaskFactory().Create("buy groceries", "", false)
				require.NoError(t, err)

				return []domain.Task{task}
			},
		},
		{
			name: "update buy bread to buy groceries and call other to call xyz",
			actualTasks: func() []domain.Task {
				dbHandle := NewDBHandler(config.GetDatabaseConfig())
				db, err := dbHandle.GetDB()
				require.NoError(t, err)
				taskStore := NewTaskStore(db)
				cleanDB(db, t)

				task, err := domain.NewTaskFactory().Create("buy bread", "", false)
				require.NoError(t, err)

				callOther, err := domain.NewTaskFactory().Create("call other", "", false)
				require.NoError(t, err)

				err = taskStore.Add(task)
				require.NoError(t, err)

				err = taskStore.Add(callOther)
				require.NoError(t, err)

				task.UpdateTitle("buy groceries")
				callOther.UpdateTitle("call xyz")

				err = taskStore.Update(task)
				require.NoError(t, err)

				err = taskStore.Update(callOther)
				require.NoError(t, err)

				var defaultTasks []*domain.DefaultTask

				err = db.Select(&defaultTasks, "SELECT * FROM todolist")
				require.NoError(t, err)

				tasks := make([]domain.Task, 0)
				for _, df := range defaultTasks {
					tasks = append(tasks, df)
				}

				cleanDB(db, t)
				return tasks
			},
			expectedTasks: func() []domain.Task {
				task, err := domain.NewTaskFactory().Create("buy groceries", "", false)
				require.NoError(t, err)

				callXyz, err := domain.NewTaskFactory().Create("call xyz", "", false)
				require.NoError(t, err)

				return []domain.Task{task, callXyz}
			},
		},
	}
	for _, testCase := range testCases {
		expectedTasks := testCase.expectedTasks()
		actualTasks := testCase.actualTasks()
		assert.Equal(t, len(expectedTasks), len(actualTasks))
		l := len(expectedTasks)
		for i := 0; i < l; i++ {
			assert.Equal(t, expectedTasks[i].GetTitle(), actualTasks[i].GetTitle())
			assert.Equal(t, expectedTasks[i].GetDescription(), actualTasks[i].GetDescription())
			assert.Equal(t, expectedTasks[i].GetStatus(), actualTasks[i].GetStatus())
			assert.Equal(t, expectedTasks[i].GetTags(), actualTasks[i].GetTags())
		}
	}
}

func TestGetTasks(t *testing.T) {
	err := config.Load()
	require.NoError(t, err)
	id := uuid.New().String()

	testCases := []struct {
		name          string
		actualTask    func() ([]domain.Task, error)
		expectedTask  func() []domain.Task
		expectedError error
	}{
		{
			name: "server buy groceries task",
			actualTask: func() ([]domain.Task, error) {
				dbHandle := NewDBHandler(config.GetDatabaseConfig())
				db, err := dbHandle.GetDB()
				require.NoError(t, err)
				taskStore := NewTaskStore(db)
				cleanDB(db, t)

				buyGroceries, err := domain.NewTaskFactory().Create("buy groceries", "", false)
				require.NoError(t, err)

				err = taskStore.Add(buyGroceries)
				require.NoError(t, err)

				list, err := taskStore.GetTasks(buyGroceries.GetID())

				cleanDB(db, t)
				return list, err
			},
			expectedTask: func() []domain.Task {
				buyGroceries, err := domain.NewTaskFactory().Create("buy groceries", "", false)
				require.NoError(t, err)

				return []domain.Task{buyGroceries}
			},
		},
		{
			name: "server all tasks",
			actualTask: func() ([]domain.Task, error) {
				dbHandle := NewDBHandler(config.GetDatabaseConfig())
				db, err := dbHandle.GetDB()
				require.NoError(t, err)
				taskStore := NewTaskStore(db)
				cleanDB(db, t)

				buyGroceries, err := domain.NewTaskFactory().Create("buy groceries", "", false)
				require.NoError(t, err)

				readXyz, err := domain.NewTaskFactory().Create("readXyz", "", false)
				require.NoError(t, err)

				err = taskStore.Add(buyGroceries)
				require.NoError(t, err)

				err = taskStore.Add(readXyz)
				require.NoError(t, err)

				list, err := taskStore.GetTasks()

				cleanDB(db, t)
				return list, err
			},
			expectedTask: func() []domain.Task {
				buyGroceries, err := domain.NewTaskFactory().Create("buy groceries", "", false)
				require.NoError(t, err)

				readXyz, err := domain.NewTaskFactory().Create("readXyz", "", false)
				require.NoError(t, err)

				return []domain.Task{buyGroceries, readXyz}
			},
		},
		{
			name: "fail to retrieve task from database",
			actualTask: func() ([]domain.Task, error) {
				dbHandle := NewDBHandler(config.GetDatabaseConfig())
				db, err := dbHandle.GetDB()
				require.NoError(t, err)
				taskStore := NewTaskStore(db)
				cleanDB(db, t)

				list, err := taskStore.GetTasks(id)

				cleanDB(db, t)
				return list, err
			},
			expectedError: errors.New("no task found for : [" + id + "]"),
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			actualTasks, err := testCase.actualTask()
			assert.Equal(t, testCase.expectedError, err)

			if err == nil {
				expectedTasks := testCase.expectedTask()
				assert.Equal(t, len(expectedTasks), len(actualTasks))
				l := len(expectedTasks)
				for i := 0; i < l; i++ {
					assert.Equal(t, expectedTasks[i].GetTitle(), actualTasks[i].GetTitle())
					assert.Equal(t, expectedTasks[i].GetDescription(), actualTasks[i].GetDescription())
					assert.Equal(t, expectedTasks[i].GetStatus(), actualTasks[i].GetStatus())
					assert.Equal(t, expectedTasks[i].GetTags(), actualTasks[i].GetTags())
				}
			}

		})
	}
}
