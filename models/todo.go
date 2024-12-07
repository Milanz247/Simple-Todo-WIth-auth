package models

import (
	"Golang/config"
)

type Todo struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Title    string `json:"title" gorm:"not null"`
	Complete bool   `json:"complete"`
	UserID   uint   `json:"user_id" gorm:"not null"` // Foreign key for User
}

// CreateUser creates a new user in the database
func CreateTodos(user *Todo) error {

	db := config.InitDB()
	if err := db.Create(user).Error; err != nil {
		return err
	}

	return nil
}