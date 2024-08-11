package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/hasbrovish/Webapp-Labs/config"
)

var db *sql.DB

func initDB() {
	cfg := config.NewConfig()
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected", db)
}
