package config

import "fmt"

type DatabaseConfig struct {
	host        string
	port        int
	username    string
	password    string
	name        string
	maxPoolSize int
	retryCount  int
}

func newDatabaseConfig() DatabaseConfig {
	return DatabaseConfig{
		host:        viperString("DB_HOST", "localhost"),
		port:        viperInt("DB_PORT"),
		name:        viperString("database", "postgres"),
		username:    viperString("DB_USER", "postgres"),
		password:    viperString("DB_PASSWORD", "password"),
		maxPoolSize: viperInt("DB_MAXPOOLSIZE", 5),
		retryCount:  viperInt("DB_RETRY_COUNT", 10),
	}
}

func (dc DatabaseConfig) Source() string {
	return fmt.Sprintf("dbname=%s user=%s password=%s host=%s port=%d sslmode=disable", dc.name, dc.username, dc.password, dc.host, dc.port)
}

func (dc DatabaseConfig) GetMaxPoolSize() int {
	return dc.maxPoolSize
}

func (dc DatabaseConfig) GetRetryCount() int {
	return dc.retryCount
}
