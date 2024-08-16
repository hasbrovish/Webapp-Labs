package services

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/hasbrovish/Webapp-Labs/internal/repositories"
	"github.com/jinzhu/gorm"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Message string `json:"message"`
}

func LoginService(db *gorm.DB, store *sessions.CookieStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) { //get user , check password , if successfull create session and login
		var loginReq LoginRequest

		// Decode the JSON request body into loginReq
		if err := json.NewDecoder(r.Body).Decode(&loginReq); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		// Extract email and password
		email := loginReq.Email
		password := loginReq.Password
		log.Println("Received email:", email, "and password:", password)
		user, err := repositories.LoginRepo(db, email, password)
		if err != "" {
			http.Error(w, err, http.StatusUnauthorized)
			return
		}

		session, _ := store.Get(r, "session-name")
		session.Values["user_id"] = user.ID
		session.Save(r, w)
		// Log the successful login
		log.Printf("Login successful for user ID: %d", user.ID)

		// Create a success message
		response := LoginResponse{
			Message: "Successfully logged in",
		}

		// Marshal the success message into JSON
		jsonResponse, err2 := json.Marshal(response)
		if err2 != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		// Set the content type and status code
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		// Write the JSON response
		w.Write(jsonResponse)

	}
}
