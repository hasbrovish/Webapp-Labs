package services

import (
	"fmt"
	"net/http"

	"github.com/jinzhu/gorm"
)

func DashboardHandler(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to the dashboard!")
	}
}
