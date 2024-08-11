package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/hasbrovish/Webapp-Labs/pkg/db"
)

func main() {

	db.InitDB()
	// dsn := "root:Jayanti@123@tcp(localhost:3306)/webapplab?allowNativePasswords=true"
	// db, err := sql.Open("mysql", dsn)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer db.Close()

	// err = db.Ping()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// log.Println("Connected successfully!")

}
