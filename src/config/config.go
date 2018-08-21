package config

import (
	"os"
	"strconv"
)

// DBConfig represents database connection configuration information.
type DBConfig struct {
	User     string
	Password string
	Host     string
	Port     int
	Name     string
}

// Config represents application configuration.
type Config struct {
	DB DBConfig
}

// NewConfig return configuration struct.
func NewConfig() (Config, error) {
	var conf Config


	port, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		return conf, err
	}
	dbConf := DBConfig{
		User: os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Host: os.Getenv("DB_HOST"),
		Port: port,
		Name: os.Getenv("DB_NAME"),
	}
	conf.DB = dbConf

	return conf, nil
}
