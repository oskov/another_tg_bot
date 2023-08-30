package config

import (
	"fmt"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	DbUser string `envconfig:"MYSQL_USER"`
	DbPass string `envconfig:"MYSQL_PASSWORD"`
	DbName string `envconfig:"MYSQL_DATABASE"`
	DbAddr string `envconfig:"DB_ADDR"`

	DbMaxIdleConnections             int `envconfig:"DB_MAX_IDLE_CONNECTIONS" default:"1"`
	DbMaxOpenConnections             int `envconfig:"DB_MAX_OPEN_CONNECTIONS" default:"10"`
	DbConnectionMaxIdleTimeInSeconds int `envconfig:"DB_CONNECTION_MAX_IDLE_TIME_IN_SECONDS" default:"60"`
	DbConnectionMaxLifeTimeInSeconds int `envconfig:"DB_CONNECTION_MAX_LIFE_TIME_IN_SECONDS" default:"3600"`

	BotToken string `envconfig:"BOT_TOKEN"`
}

var conf Config
var emptyConf Config

func GetConfig() (Config, error) {
	if conf != emptyConf {
		return conf, nil
	}

	err := envconfig.Process("", &conf)
	if err != nil {
		return conf, fmt.Errorf("can't process the config: %w", err)
	}

	return conf, nil
}
