package config

import "github.com/spf13/viper"

type config struct {
	databaseConfig DatabaseConfig
	serverConfig   ServerConfig
	logConfig      LogConfig
}

var appConfig config

func Load() error {
	viper.AutomaticEnv()
	viper.SetConfigName("application")
	viper.AddConfigPath("./")
	viper.AddConfigPath("../")
	viper.AddConfigPath("../../")
	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	appConfig = config{
		databaseConfig: newDatabaseConfig(),
		serverConfig:   newSeverConfig(),
		logConfig:      newLogConfig(),
	}

	return nil
}

func GetDatabaseConfig() DatabaseConfig {
	return appConfig.databaseConfig
}

func GetServerConfig() ServerConfig {
	return appConfig.serverConfig
}

func GetLogConfig() LogConfig {
	return appConfig.logConfig
}
