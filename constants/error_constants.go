package constants

const (
	InvalidTaskErrorMessage = "%s : task cannot have empty description"

	InvalidTaskUUIDErrorMessage = "%s : invalid uuid %s for the task"

	ErrorTaskValidationFailed = "%s : failed to validate task %v"

	TaskRemoveFailedErrorMessage = "%s : task %v not found in the todo list"

	ErrorFailedToInsertTask = "%s : failed to insert task %v into database : %v"

	ErrorFailedToAddTask = "%s : failed to insert task %v : %v"

	ErrorRowsAffected = "%s : failed to insert task %v into database : %v"

	ErrorDatabaseFailedToLoad = "%s : failed to load database for %s : %v"

	ErrorDatabasePingFailed = "%s : failed to ping database : %v"

	ErrorFailedToDeleteTask = "%s : failed to delete task %v from db : %v"

	ErrorFailedToRemoveTask = "%s : failed to remove task %v : %v"

	ErrorRowsAffectedDelete = "%s : failed to delete task %v from db : %v"

	ErrorFailedToUpdateTask = "%s : failed to update task %v from db : %v"

	ErrorTaskUpdateFailed = "%s : failed to update task %v : %v"

	ErrorRowsAffectedUpdate = "%s : failed to update task %v from db : %v"

	ErrorTaskNotPresentInDatabase = "%s : failed to find task %v in database : %v"

	ErrorTaskWithIDNotPresentInDatabase = "%s : failed to find task with id %s in database : %v"

	ErrorTaskNotFound = "%s : failed to find task with id %s : %v"

	ErrorFailedToGetAllTasksFromDatabase = "%s : failed to get all tasks from db : %v"

	ErrorFailedToGetAllTasks = "%s : failed to get all tasks : %v"

	ErrorFailedToWriteAPIResponse = "%s : failed to write response: %v"

	ErrorFailedToReadRequestBody = "%s : failed to read request body : %v"

	ErrorFailedToUnMarshalRequestBody = "%s : failed to unmarshal request body %v : %v"

	ErrorRequestBodyValidationFailed = "%s : request body validation failed"

	ErrorFailedToLoadConfig = "%s : failed to load config : %v"

	ErrorFailedToShutdownServer = "%s : failed to shutdown server : %v"

	ErrorFailedToDisplayConfig = "%s : failed to display config"

	ErrorFailedToGetDatabaseDriver = "%s : failed to get database driver : %v"

	ErrorFailedToGetMigrate = "%s : failed to get migrate : %v"

	ErrorMigrationFailed = "%s : failed to run migrations : %v"

	ErrorRollbackFailed = "%s : failed to rollback migrations : %v"
)
