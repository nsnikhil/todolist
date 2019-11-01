package constants

const (
	DatabaseDriverName  = "postgres"
	DatabaseHost        = "127.0.0.1"
	DatabasePort        = "5432"
	DatabaseUser        = "postgres"
	DatabasePassword    = ""
	DatabaseMaxPoolSize = "5"

	MigrationPath   = "./store/migrations"

	RollBackStep = -1
)
