package db

import (
	"fmt"
	"log"

	"github.com/hasbrovish/Webapp-Labs/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// var db *sql.DB
var db *gorm.DB
var cfg config.Config

func InitDB() {

	cfg = config.NewConfig()
	fmt.Println(cfg)
	databaseURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DBName)
	fmt.Println(databaseURI)
	var err error
	db, err = gorm.Open("mysql", databaseURI)
	if err != nil {
		log.Fatal(err)
	}
	d := db
	pingErr := d.DB().Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")
}

func GetDB() *gorm.DB {
	return db
}
