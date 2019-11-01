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
		actualTodoListStore   TodoListStoreInterface
		expectedTodoListStore TodoListStoreInterface
	}{
		{
			name:                  "test create new todo list store",
			actualTodoListStore:   NewTodoListStore(&sqlx.DB{}),
			expectedTodoListStore: TodoListStore{db: &sqlx.DB{}},
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
		name             string
		actualTodoList   func() *domain.TodoList
		expectedTodoList func() *domain.TodoList
	}{
		{
			name: "store buy groceries task in db",
			actualTodoList: func() *domain.TodoList {
				dbHandle := NewDBHandle(config.GetDatabaseConfig())
				db, err := dbHandle.GetDB()
				require.NoError(t, err)
				cleanDB(db, t)

				taskStore := NewTodoListStore(db)

				task, err := domain.NewTask(uuid.New().String(), "buy groceries", false)
				require.NoError(t, err)

				err = taskStore.Add(task)
				require.NoError(t, err)

				var list domain.TodoList

				err = db.Select(&list.Tasks, "SELECT * FROM todolist")
				require.NoError(t, err)

				cleanDB(db, t)
				return &list
			},
			expectedTodoList: func() *domain.TodoList {
				task, _ := domain.NewTask(uuid.New().String(), "buy groceries", false)
				return domain.NewTodoList(task)
			},
		},
		{
			name: "store buy groceries and read xyz task in db",
			actualTodoList: func() *domain.TodoList {
				dbHandle := NewDBHandle(config.GetDatabaseConfig())
				db, err := dbHandle.GetDB()
				require.NoError(t, err)
				taskStore := NewTodoListStore(db)
				cleanDB(db, t)

				buyGroceries, err := domain.NewTask(uuid.New().String(), "buy groceries", false)
				require.NoError(t, err)

				readXyz, err := domain.NewTask(uuid.New().String(), "read xyz", false)
				require.NoError(t, err)

				err = taskStore.Add(buyGroceries)
				require.NoError(t, err)

				err = taskStore.Add(readXyz)
				require.NoError(t, err)

				var list domain.TodoList

				err = db.Select(&list.Tasks, "SELECT * FROM todolist")
				require.NoError(t, err)

				cleanDB(db, t)
				return &list
			},
			expectedTodoList: func() *domain.TodoList {
				buyGroceries, _ := domain.NewTask(uuid.New().String(), "buy groceries", false)
				readXyz, _ := domain.NewTask(uuid.New().String(), "read xyz", false)
				return domain.NewTodoList(readXyz, buyGroceries)
			},
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			expectedTodoList := testCase.expectedTodoList()
			actualTodoList := testCase.actualTodoList()
			assert.Equal(t, len(expectedTodoList.Tasks), len(actualTodoList.Tasks))
			for i := 0; i < len(expectedTodoList.Tasks); i++ {
				assert.Equal(t, expectedTodoList.Tasks[i].GetDescription(), actualTodoList.Tasks[i].GetDescription())
				assert.Equal(t, expectedTodoList.Tasks[i].GetStatus(), actualTodoList.Tasks[i].GetStatus())
			}
		})
	}
}

func TestRemoveTaskFromDB(t *testing.T) {
	err := config.Load()
	require.NoError(t, err)

	testCases := []struct {
		name             string
		actualTodoList   func() *domain.TodoList
		expectedTodoList func() *domain.TodoList
	}{
		{
			name: "remove buy groceries task from db",
			actualTodoList: func() *domain.TodoList {
				dbHandle := NewDBHandle(config.GetDatabaseConfig())
				db, err := dbHandle.GetDB()
				require.NoError(t, err)
				taskStore := NewTodoListStore(db)
				cleanDB(db, t)

				task, err := domain.NewTask(uuid.New().String(), "buy groceries", false)
				require.NoError(t, err)

				err = taskStore.Add(task)
				require.NoError(t, err)

				err = taskStore.Remove(task.GetID())
				require.NoError(t, err)

				var list domain.TodoList

				err = db.Select(&list.Tasks, "SELECT * FROM todolist")
				require.NoError(t, err)

				cleanDB(db, t)
				return &list
			},
			expectedTodoList: func() *domain.TodoList {
				return &domain.TodoList{}
			},
		},
		{
			name: "remove buy groceries and read xyz task from db",
			actualTodoList: func() *domain.TodoList {
				dbHandle := NewDBHandle(config.GetDatabaseConfig())
				db, err := dbHandle.GetDB()
				require.NoError(t, err)
				taskStore := NewTodoListStore(db)
				cleanDB(db, t)

				buyGroceries, err := domain.NewTask(uuid.New().String(), "buy groceries", false)
				require.NoError(t, err)

				readXyz, err := domain.NewTask(uuid.New().String(), "read xyz", false)
				require.NoError(t, err)

				err = taskStore.Add(buyGroceries)
				require.NoError(t, err)

				err = taskStore.Add(readXyz)
				require.NoError(t, err)

				err = taskStore.Remove(buyGroceries.GetID())
				require.NoError(t, err)

				err = taskStore.Remove(readXyz.GetID())
				require.NoError(t, err)

				var list domain.TodoList

				err = db.Select(&list.Tasks, "SELECT * FROM todolist")
				require.NoError(t, err)

				cleanDB(db, t)
				return &list
			},
			expectedTodoList: func() *domain.TodoList {
				return &domain.TodoList{}
			},
		},
		{
			name: "remove buy groceries and read xyz tasks from db",
			actualTodoList: func() *domain.TodoList {
				dbHandle := NewDBHandle(config.GetDatabaseConfig())
				db, err := dbHandle.GetDB()
				require.NoError(t, err)
				taskStore := NewTodoListStore(db)
				cleanDB(db, t)

				buyGroceries, err := domain.NewTask(uuid.New().String(), "buy groceries", false)
				require.NoError(t, err)

				readXyz, err := domain.NewTask(uuid.New().String(), "read xyz", false)
				require.NoError(t, err)

				callPerson, err := domain.NewTask(uuid.New().String(), "call person", false)
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

				var list domain.TodoList

				err = db.Select(&list.Tasks, "SELECT * FROM todolist")
				require.NoError(t, err)

				cleanDB(db, t)
				return &list
			},
			expectedTodoList: func() *domain.TodoList {
				callPerson, _ := domain.NewTask(uuid.New().String(), "call person", false)
				return domain.NewTodoList(callPerson)
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			expectedTodoList := testCase.expectedTodoList()
			actualTodoList := testCase.actualTodoList()
			assert.Equal(t, len(expectedTodoList.Tasks), len(actualTodoList.Tasks))
			for i := 0; i < len(expectedTodoList.Tasks); i++ {
				assert.Equal(t, expectedTodoList.Tasks[i].GetDescription(), actualTodoList.Tasks[i].GetDescription())
				assert.Equal(t, expectedTodoList.Tasks[i].GetStatus(), actualTodoList.Tasks[i].GetStatus())
			}
		})
	}
}

func TestUpdateTaskInDB(t *testing.T) {
	err := config.Load()
	require.NoError(t, err)

	testCases := []struct {
		name             string
		actualTodoList   func() *domain.TodoList
		expectedTodoList func() *domain.TodoList
	}{
		{
			name: "update buy bread to buy groceries",
			actualTodoList: func() *domain.TodoList {
				dbHandle := NewDBHandle(config.GetDatabaseConfig())
				db, err := dbHandle.GetDB()
				require.NoError(t, err)
				taskStore := NewTodoListStore(db)
				cleanDB(db, t)

				task, err := domain.NewTask(uuid.New().String(), "buy bread", false)
				require.NoError(t, err)

				err = taskStore.Add(task)
				require.NoError(t, err)

				task.Description = "buy groceries"

				err = taskStore.Update(task)
				require.NoError(t, err)

				var list domain.TodoList

				err = db.Select(&list.Tasks, "SELECT * FROM todolist")
				require.NoError(t, err)

				cleanDB(db, t)
				return &list
			},
			expectedTodoList: func() *domain.TodoList {
				task, _ := domain.NewTask(uuid.New().String(), "buy groceries", false)
				return domain.NewTodoList(task)
			},
		},
		{
			name: "update buy bread to buy groceries and call other to call xyz",
			actualTodoList: func() *domain.TodoList {
				dbHandle := NewDBHandle(config.GetDatabaseConfig())
				db, err := dbHandle.GetDB()
				require.NoError(t, err)
				taskStore := NewTodoListStore(db)
				cleanDB(db, t)

				task, err := domain.NewTask(uuid.New().String(), "buy bread", false)
				require.NoError(t, err)

				callOther, err := domain.NewTask(uuid.New().String(), "call other", false)
				require.NoError(t, err)

				err = taskStore.Add(task)
				require.NoError(t, err)

				err = taskStore.Add(callOther)
				require.NoError(t, err)

				task.Description = "buy groceries"
				callOther.Description = "call xyz"

				err = taskStore.Update(task)
				require.NoError(t, err)

				err = taskStore.Update(callOther)
				require.NoError(t, err)

				var list domain.TodoList

				err = db.Select(&list.Tasks, "SELECT * FROM todolist")
				require.NoError(t, err)

				cleanDB(db, t)
				return &list
			},
			expectedTodoList: func() *domain.TodoList {
				task, _ := domain.NewTask(uuid.New().String(), "buy groceries", false)
				callXyz, _ := domain.NewTask(uuid.New().String(), "call xyz", false)
				return domain.NewTodoList(callXyz, task)
			},
		},
	}
	for _, testCase := range testCases {
		expectedTodoList := testCase.expectedTodoList()
		actualTodoList := testCase.actualTodoList()
		assert.Equal(t, len(expectedTodoList.Tasks), len(actualTodoList.Tasks))
		for i := 0; i < len(expectedTodoList.Tasks); i++ {
			assert.Equal(t, expectedTodoList.Tasks[i].GetDescription(), actualTodoList.Tasks[i].GetDescription())
			assert.Equal(t, expectedTodoList.Tasks[i].GetStatus(), actualTodoList.Tasks[i].GetStatus())
		}
	}
}

func TestGetAllTask(t *testing.T) {
	err := config.Load()
	require.NoError(t, err)

	testCases := []struct {
		name             string
		actualTodoList   func() domain.TodoListInterface
		expectedTodoList func() domain.TodoListInterface
	}{
		{
			name: "test get all todolist",
			actualTodoList: func() domain.TodoListInterface {
				dbHandle := NewDBHandle(config.GetDatabaseConfig())
				db, err := dbHandle.GetDB()
				require.NoError(t, err)
				taskStore := NewTodoListStore(db)
				cleanDB(db, t)

				buyGroceries, err := domain.NewTask(uuid.New().String(), "buy groceries", false)
				require.NoError(t, err)

				readXyz, err := domain.NewTask(uuid.New().String(), "read xyz", false)
				require.NoError(t, err)

				err = taskStore.Add(buyGroceries)
				require.NoError(t, err)

				err = taskStore.Add(readXyz)
				require.NoError(t, err)

				list, err := taskStore.GetTodoList()
				require.NoError(t, err)

				cleanDB(db, t)
				return list
			},
			expectedTodoList: func() domain.TodoListInterface {
				buyGroceries, _ := domain.NewTask(uuid.New().String(), "buy groceries", false)
				readXyz, _ := domain.NewTask(uuid.New().String(), "read xyz", false)
				return domain.NewTodoList(readXyz, buyGroceries)
			},
		},
		{
			name: "test get all todolist from database",
			actualTodoList: func() domain.TodoListInterface {
				dbHandle := NewDBHandle(config.GetDatabaseConfig())
				db, err := dbHandle.GetDB()
				require.NoError(t, err)
				taskStore := NewTodoListStore(db)
				cleanDB(db, t)

				buyGroceries, err := domain.NewTask(uuid.New().String(), "buy groceries", false)
				require.NoError(t, err)

				readXyz, err := domain.NewTask(uuid.New().String(), "read xyz", false)
				require.NoError(t, err)

				callXyz, err := domain.NewTask(uuid.New().String(), "call xyz", false)
				require.NoError(t, err)

				planTheThing, err := domain.NewTask(uuid.New().String(), "plan the thing", false)
				require.NoError(t, err)

				err = taskStore.Add(buyGroceries)
				require.NoError(t, err)

				err = taskStore.Add(readXyz)
				require.NoError(t, err)

				err = taskStore.Add(callXyz)
				require.NoError(t, err)

				err = taskStore.Add(planTheThing)
				require.NoError(t, err)

				list, err := taskStore.GetTodoList()
				require.NoError(t, err)

				cleanDB(db, t)
				return list
			},
			expectedTodoList: func() domain.TodoListInterface {
				buyGroceries, _ := domain.NewTask(uuid.New().String(), "buy groceries", false)
				readXyz, _ := domain.NewTask(uuid.New().String(), "read xyz", false)
				callXyz, _ := domain.NewTask(uuid.New().String(), "call xyz", false)
				planTheThing, _ := domain.NewTask(uuid.New().String(), "plan the thing", false)
				return domain.NewTodoList(planTheThing, buyGroceries, readXyz, callXyz)
			},
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			expectedTodoList := testCase.expectedTodoList().(*domain.TodoList)
			actualTodoList := testCase.actualTodoList().(*domain.TodoList)
			assert.Equal(t, len(expectedTodoList.Tasks), len(actualTodoList.Tasks))
			for i := 0; i < len(expectedTodoList.Tasks); i++ {
				assert.Equal(t, expectedTodoList.Tasks[i].GetDescription(), actualTodoList.Tasks[i].GetDescription())
				assert.Equal(t, expectedTodoList.Tasks[i].GetStatus(), actualTodoList.Tasks[i].GetStatus())
			}
		})
	}

}

func TestGetTask(t *testing.T) {
	err := config.Load()
	require.NoError(t, err)

	testCases := []struct {
		name          string
		actualTask    func() (domain.TaskInterface, error)
		expectedTask  func() domain.TaskInterface
		expectedError error
	}{
		{
			name: "get buy groceries task",
			actualTask: func() (domain.TaskInterface, error) {
				dbHandle := NewDBHandle(config.GetDatabaseConfig())
				db, err := dbHandle.GetDB()
				require.NoError(t, err)
				taskStore := NewTodoListStore(db)
				cleanDB(db, t)

				buyGroceries, err := domain.NewTask(uuid.New().String(), "buy groceries", false)
				require.NoError(t, err)

				err = taskStore.Add(buyGroceries)
				require.NoError(t, err)

				list, err := taskStore.GetTask(buyGroceries.GetID())

				cleanDB(db, t)
				return list, err
			},
			expectedTask: func() domain.TaskInterface {
				buyGroceries, _ := domain.NewTask(uuid.New().String(), "buy groceries", false)
				return buyGroceries
			},
		},
		{
			name: "fail to retrieve task from database",
			actualTask: func() (domain.TaskInterface, error) {
				dbHandle := NewDBHandle(config.GetDatabaseConfig())
				db, err := dbHandle.GetDB()
				require.NoError(t, err)
				taskStore := NewTodoListStore(db)
				cleanDB(db, t)

				list, err := taskStore.GetTask(uuid.New().String())

				cleanDB(db, t)
				return list, err
			},
			expectedError: errors.New("sql: no rows in result set"),
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			task, err := testCase.actualTask()
			assert.Equal(t, testCase.expectedError, err)
			if err == nil {
				expectedTask := testCase.expectedTask()
				assert.Equal(t, expectedTask.GetDescription(), task.GetDescription())
				assert.Equal(t, expectedTask.GetDescription(), task.GetDescription())
			}
		})
	}
}
