package services

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/hasbrovish/Webapp-Labs/internal/models"
	"github.com/hasbrovish/Webapp-Labs/internal/utils"
	"golang.org/x/crypto/bcrypt"
)

type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	CreateUser := &models.User{}
	utils.ParseBody(r, CreateUser)
	hash, err := bcrypt.GenerateFromPassword([]byte(CreateUser.Password), bcrypt.DefaultCost)
	CreateUser.Password = string(hash)
	if err != nil {
		log.Fatal(err)
	}
	//CreateUser.Password, err = utils.GetHash(CreateUser.Password)
	user, err := CreateUser.CreateUser()

	if err != nil {
		// Log the error for internal tracking
		log.Printf("Failed to create user: %v", err)

		// Return a structured error response
		writeErrorResponse(w, http.StatusInternalServerError, "Failed to create user", err.Error())
		//http.Error(w, fmt.Sprintf("Failed to create user: %v", err), http.StatusInternalServerError)
		return
	}

	res, _ := json.Marshal(user)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// need to clear idea ofwh atis the rein  repo layer what is there in service laye
// r , i got very confused with this now
func writeErrorResponse(w http.ResponseWriter, statusCode int, err string, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(ErrorResponse{
		Error:   err,
		Message: message,
	})
}
