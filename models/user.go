package models

import (
	"Golang/config"
	"errors"

	"gorm.io/gorm"
)

// User represents a user in the database
type User struct {
	ID       uint   `gorm:"primaryKey"`
	Email    string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	Name     string `gorm:"not null"` 
	Todos    []Todo // One-to-many relationship
}

// CreateUser creates a new user in the database
func CreateUser(user *User) error {

	db := config.InitDB()
	if err := db.Create(user).Error; err != nil {
		return err
	}

	return nil
}

// GetUserByEmail retrieves a user by their email from the database
func GetUserByEmail(email string) (*User, error) {

	db := config.InitDB()
	var user User

	if err := db.Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil 
		}
		return nil, err
	}
	
	return &user, nil
}
