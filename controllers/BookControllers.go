package controllers

import (
	"BOOK-STORE/models"
	"main/initializers"
	"regexp"

	"github.com/gin-gonic/gin"
)

// Book post
func BookPost(c *gin.Context) {
	//get title, author, published date, ISBN, and price from the request body
	var body struct {
		Title         string
		Author        string
		PublishedDate string
		ISBN          string
		Price         float64
	}
	//Bind the request body to the body struct
	if c.Bind(&body) != nil {
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}
	// Query the database for a book with the given ISBN
	var bookModel models.Book
	bookExists := initializers.DB.Where("ISBN = ?", body.ISBN).First(&bookModel).RowsAffected > 0
	if bookExists {
		c.JSON(400, gin.H{"error": "Book already exists"})
		return
	}
	//Make a regex to check the published date format that should be something like this 2006-01-02
	var regex = "^\\d{4}-\\d{2}-\\d{2}$"

	//Check if the published date matches regex
	match, _ := regexp.MatchString(regex, body.PublishedDate)
	//If the published date does not match the regex, return an error message
	if !match {
		c.JSON(400, gin.H{"error": "Invalid published date format"})
		return
	}

	//create a new book with the title, author, published date, ISBN, and price
	book := models.Book{Title: body.Title, Author: body.Author, PublishedDate: body.PublishedDate, ISBN: body.ISBN, Price: body.Price}
	result := initializers.DB.Create(&book)
	//Check for errors
	if result.Error != nil {
		c.JSON(500, gin.H{"error": "Error creating book"})
		return
	}
	//respond with the book
	c.JSON(200, gin.H{"book": book})
}
