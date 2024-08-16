package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	password := "test4"                                                              // the plain-text password
	hashedPassword := "$2a$10$tD27aTaP9RYfNPcrAzvJO.mUNkzNxydwFBDLcK18PzyNsQ/07CJUC" // the stored hash from DB

	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		fmt.Println("Passwords do not match!")
	} else {
		fmt.Println("Passwords match!")
	}
}
