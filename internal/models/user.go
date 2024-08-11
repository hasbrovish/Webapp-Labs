package models

import (
	"errors"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/hasbrovish/Webapp-Labs/pkg/db"
	"github.com/jinzhu/gorm"
)

type User struct {
	ID        uint      `gorm:"unique;primaryKey;autoIncrement"`
	Firstname string    `gorm:"" json:"firstname"`
	Lastname  string    `gorm:"" json:"lastname"`
	Email     string    `gorm:"" json:"email"`
	Password  string    `gorm:"" json:"-"` // this `-` is used to skip this field from response
	Gender    string    `gorm:"" json:"gender"`
	Phone     int       `gorm:"" json:"phone"`
	CreatedAt time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"updated_at"`
}

var d *gorm.DB

var (
	ErrUserExists = errors.New("user with this email already exists, please try logging in")
)

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
		return nil, ErrUserExists
	}

	// Attempt to create the new user record
	if err := d.Create(u).Error; err != nil {
		// Check for unique constraint violation (duplicate entry)
		if gorm.IsRecordNotFoundError(err) || isDuplicateEntryError(err) {
			return nil, ErrUserExists
		}
		return nil, err
	}

	// Return the created user and no error
	return u, nil
}

func isDuplicateEntryError(err error) bool {
	// Check if the error is a MySQL error
	if mysqlErr, ok := err.(*mysql.MySQLError); ok {
		// Error 1062 is the MySQL error code for a duplicate entry
		return mysqlErr.Number == 1062
	}
	return false
}
