package models

import ("gorm.io/gorm"
)

// Book is a struct that represents a book in the database
type Book struct {
	gorm.Model
	Title  string
	Author string
PublishedDate string
	ISBN string
	Price float64
}