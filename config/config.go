package config

import (
	"github.com/spf13/viper"
	"os"
)

type Config struct {
	DB    DB    `json:"db"`
	Redis Redis `json:"redis"`
}

type DB struct {
	Host     string `json:"host"`
	User     string `json:"user"`
	Password string `json:"password"`
	Database string `json:"database"`
}

type Redis struct {
	Addr     string `json:"addr"`
	Password string `json:"password"`
	Db       int    `json:"db"`
}

var Cfg *Config

func init() {
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	config := viper.New()

	config.AddConfigPath(path)
	config.SetConfigName("./config/config")
	config.SetConfigType("yaml")

	if err := config.ReadInConfig(); err != nil {
		panic(err)
	}

	err = config.Unmarshal(&Cfg)
	if err != nil {
		panic(err)
	}
}
