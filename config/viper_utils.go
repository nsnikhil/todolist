package config

import "github.com/spf13/viper"

func viperString(config string, defaultVal ...string) string {
	if len(defaultVal) > 0 {
		viper.SetDefault(config, defaultVal[0])
	}
	return viper.GetString(config)
}

func viperInt(config string, defaultVal ...int) int {
	if len(defaultVal) > 0 {
		viper.SetDefault(config, defaultVal[0])
	}
	return viper.GetInt(config)
}
