package handlers

import (
	"github.com/gorilla/mux"
	"github.com/hasbrovish/Webapp-Labs/internal/services"
)

var RegisterUserHandler = func(router *mux.Router) {
	router.HandleFunc("/user", services.CreateUser).Methods("POST")
	//router.HandleFunc("/user/{userId}", services.GetUser).Methods("GET")
}
