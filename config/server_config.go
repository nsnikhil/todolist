package config

import "fmt"

type ServerConfig struct {
	host string
	port int
}

func newSeverConfig() ServerConfig {
	return ServerConfig{
		host: viperString("APP_HOST"),
		port: viperInt("APP_PORT"),
	}
}

func (sc ServerConfig) Address() string {
	return fmt.Sprintf(":%d", sc.port)
}
