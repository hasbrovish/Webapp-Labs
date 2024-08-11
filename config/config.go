package config

import (
	"os"

	"github.com/go-sql-driver/mysql"
)

// type Config struct {
// 	User     string
// 	Password string
// 	Net      string
// 	Host     string
// 	Port     string
// 	DBName   string
// }

func NewConfig() mysql.Config {
	c := mysql.Config{
		User:   os.Getenv("DBUSER"),
		Passwd: os.Getenv("DBPASS"),
		Net:    "tcp",
		Addr:   os.Getenv("DBHOST") + ":" + os.Getenv("DBPORT"),
		DBName: os.Getenv("DBNAME"),
	}
	return c
}
