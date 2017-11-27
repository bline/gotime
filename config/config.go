package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config = viper.Viper

var config *Config
var Env = "production"

// Singleton
func New(env string) (*Config, error) {
	if env != "" {
		Env = env
	} else {
		env = Env
	}
	var err error
	if config != nil {
		return config, err
	}
	config = viper.New()
	config.SetConfigType("yaml")
	config.SetConfigName(env)
	if env == "test" {
		config.AddConfigPath("../config/")
	} else {
		config.AddConfigPath("config/")
	}
	err = config.ReadInConfig()
	if err != nil {
		log.Fatal("error on parsing configuration file")
	}
	setDefaults(config)
	return config, err
}

func setDefaults(config *Config) {
	config.SetDefault("SQL.Driver", "mysql")
	config.SetDefault("SQL.Connection", "gotime:Passw0rd@127.0.0.1:3306/gotime")
}

func GetConfig() (*Config) {
	if config == nil {
		log.Printf("No config in GetConfig(), check main entry point")
		c, _ := New("")
		return c
	}
	return config
}

func ReNew(env string) (*Config, error) {
	config = nil
	return New(env)
}
