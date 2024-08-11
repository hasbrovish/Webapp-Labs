package config

import (
	"os"
)

// TODO create your own
type Config struct {
	User     string
	Password string
	Net      string
	Host     string
	Port     string
	DBName   string
}

func NewConfig() Config {
	return Config{
		User:     os.Getenv("DBUSER"),
		Password: os.Getenv("DBPASS"),
		Net:      "tcp",
		Host:     os.Getenv("DBHOST"),
		Port:     os.Getenv("DBPORT"),
		DBName:   os.Getenv("DBNAME"),
	}
}
