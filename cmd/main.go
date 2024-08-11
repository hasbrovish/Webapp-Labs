package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/hasbrovish/Webapp-Labs/internal/handlers"
)

func main() {
	r := mux.NewRouter()
	handlers.RegisterUserHandler(r)
	http.Handle("/", r)
	fmt.Println("Starting server at 9090")
	log.Fatal(http.ListenAndServe(":9090", r))
}
