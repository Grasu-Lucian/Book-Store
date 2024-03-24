package models

import "gorm.io/gorm"

// User is a struct that represents a user in the database
type User struct {
	gorm.Model
	Username string `gorn:"unique"`
	Password string
}
