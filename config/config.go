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
	if *config != nil {
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
	return config, err
}

func GetConfig() (*Config) {
	if config == nil {
		c, _ := New("")
		return c
	}
	return config
}

func ReNew(env string) (*Config, error) {
	config = nil
	return New(env)
}
