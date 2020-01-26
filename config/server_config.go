package config

import "fmt"

type ServerConfig struct {
	host     string
	port     int
	protocol string
}

func newSeverConfig() ServerConfig {
	return ServerConfig{
		host:     viperString("APP_HOST"),
		port:     viperInt("APP_PORT"),
		protocol: viperString("APP_PROTOCOL"),
	}
}

func (sc ServerConfig) Address() string {
	return fmt.Sprintf(":%d", sc.port)
}

func (sc ServerConfig) Protocol() string {
	return sc.protocol
}

