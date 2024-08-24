package repositories

import (
	"log"

	"github.com/hasbrovish/Webapp-Labs/internal/models"
	"github.com/hasbrovish/Webapp-Labs/internal/utils"
	"github.com/jinzhu/gorm"
)

func LoginRepo(db *gorm.DB, email string, password string) (models.User, string) {
	var user models.User
	log.Println("Login attempt with email:", email)

	// Fetch the user by email
	if err := db.Where("email = ?", email).First(&user).Error; err != nil {
		panic(err)
		return user, "Invalid email"
	}

	log.Println("Stored hashed password:", user.Password)

	// Compare the provided password with the hashed password in the database
	if !utils.CheckPasswordHash(user.Password, password) {
		log.Println("Password does not match")
		return user, "Invalid password"
	}

	log.Println("Password matches")
	return user, ""
}
