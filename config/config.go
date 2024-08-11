package config

import "os"

type Config struct {
	User     string
	Password string
	Net      string
	Host     string
	Port     string
	DBName   string
}

func (c *Config) NewConfig() {
	c = &Config{
		User:     os.Getenv("DBUSER"),
		Password: os.Getenv("DBPASS"),
		Net:      "tcp",
		Host:     os.Getenv("DBHOST"),
		Port:     os.Getenv("DBPORT"),
		DBName:   os.Getenv("DBNAME"),
	}
}
