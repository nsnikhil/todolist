package constants

const (
	InsertIntoTodoListTable = `INSERT INTO todolist (id, description, status) values ($1, $2, $3);`

	DeleteFromTodoListTable = `DELETE FROM todolist where id = $1`

	UpdateTodoListTable = `UPDATE todolist set description = $1, status = $2 where id = $3`

	FindTaskInDatabase = `SELECT id, description, status FROM todolist where id = $1`

	GetAllTasksFromDatabase = `Select id, description, status FROM todolist`
)
