package handlers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/hasbrovish/Webapp-Labs/internal/services"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB
var store = sessions.NewCookieStore([]byte("your-secret-key"))

var RegisterUserHandler = func(router *mux.Router, db *gorm.DB, store *sessions.CookieStore) {
	router.HandleFunc("/register", services.CreateUser).Methods("POST")
	//router.HandleFunc("/user/{userId}", services.GetUser).Methods("GET")
	router.HandleFunc("/login", services.LoginService(db, store)).Methods("POST") // Match password + session creation
	//router.HandleFunc("/logout", services.LogoutService).Methods("GET")
	router.Handle("/dashboard", AuthMiddleware(store)(http.HandlerFunc(DashboardHandler(db)))).Methods("GET")
	router.HandleFunc("/logout", services.LogoutService(store)).Methods("GET")

}

func DashboardHandler(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to the dashboard!")
	}
}
