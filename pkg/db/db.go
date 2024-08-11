package db

import (
	"database/sql"
)

var db *sql.DB

var config Config

func initDB() {

	var err error
	db, err = sql.Open("mysql")
}
