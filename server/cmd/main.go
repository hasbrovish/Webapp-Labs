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
	// Apply CORS middleware
	r.Use(corsMiddleware)
	r.Methods(http.MethodOptions).HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:4200")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.WriteHeader(http.StatusOK)
	})
	handlers.RegisterUserHandler(r, d, store)
	http.Handle("/", r)
	fmt.Println("Starting server at 9090")
	log.Fatal(http.ListenAndServe(":9090", r))
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:4200")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Handle preflight request
		// if r.Method == "OPTIONS" {
		// 	w.WriteHeader(http.StatusOK)
		// 	return
		// }

		next.ServeHTTP(w, r)
	})
}
