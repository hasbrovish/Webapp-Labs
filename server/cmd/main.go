package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/hasbrovish/Webapp-Labs/internal/handlers"
	"github.com/hasbrovish/Webapp-Labs/pkg/db"
)

func main() {
	db.InitDB()
	d := db.GetDB()
	store := sessions.NewCookieStore([]byte("your-secret-key"))
	r := mux.NewRouter()
	handlers.RegisterUserHandler(r, d, store)
	http.Handle("/", r)
	fmt.Println("Starting server at 9090")
	log.Fatal(http.ListenAndServe(":9090", r))
}
