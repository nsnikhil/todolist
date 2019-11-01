package config

import "fmt"

type DatabaseConfig struct {
	host        string
	port        int
	username    string
	password    string
	name        string
	maxPoolSize int
}

func newDatabaseConfig() DatabaseConfig {
	return DatabaseConfig{
		host:        viperString("db_host", "localhost"),
		port:        viperInt("db_port"),
		name:        viperString("database", "postgres"),
		username:    viperString("db_user", "postgres"),
		password:    viperString("db_password"),
		maxPoolSize: viperInt("db_maxpoolsize", 50),
	}
}

func (dc DatabaseConfig) String() string {
	return fmt.Sprintf("dbname=%s user=%s password='%s' host=%s port=%d sslmode=disable", dc.name, dc.username, dc.password, dc.host, dc.port)
}

func (dc DatabaseConfig) GetMaxPoolSize() int {
	return dc.maxPoolSize
}
