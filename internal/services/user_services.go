package services

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/hasbrovish/Webapp-Labs/internal/models"
	"github.com/hasbrovish/Webapp-Labs/internal/utils"
	"golang.org/x/crypto/bcrypt"
)

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
		log.Fatal(err)
	}
	res, _ := json.Marshal(user)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//need to clear idea ofwh atis the rein  repo layer what is there in service laye
//r , i got very confused with this now
