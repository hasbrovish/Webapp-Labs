package models

import (
	"fmt"
	"time"

	"github.com/hasbrovish/Webapp-Labs/pkg/db"
	"github.com/jinzhu/gorm"
)

type User struct {
	ID        uint      `gorm:"unique;primaryKey;autoIncrement"`
	Firstname string    `gorm:"" json:"firstname"`
	Lastname  string    `gorm:"" json:"lastname"`
	Email     string    `gorm:"" json:"email"`
	Password  string    `gorm:"" json:"-"`
	Gender    string    `gorm:"" json:"gender"`
	Phone     int       `gorm:"" json:"phone"`
	CreatedAt time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"updated_at"`
}

var d *gorm.DB

func (User) TableName() string {
	return "user"
}
func init() {
	db.InitDB()
	d = db.GetDB()
	if d == nil {
		panic("DB is not initialized")
	}

	err := d.AutoMigrate(&User{}).Error
	if err != nil {
		panic("Failed to auto-migrate: " + err.Error())
	}

}

func (u *User) CreateUser() (*User, error) {
	if !d.NewRecord(u) {
		return nil, fmt.Errorf("record already exists")
	}

	// Attempt to create the new user record
	if err := d.Create(u).Error; err != nil {
		return nil, err
	}

	// Return the created user and no error
	return u, nil
}
