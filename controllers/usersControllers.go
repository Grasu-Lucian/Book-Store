package controllers

import (
	"BOOK-STORE/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"main/initializers"
	
)

func Signup(c *gin.Context) {
	//get username and password from the request body
	var body struct {
		Username string
		Password string
	}
	// Query the database for a user with the given username
	var userModel models.User
	userExists := initializers.DB.Where("username = ?", body.Username).First(&userModel).RowsAffected > 0
	// If the user exists, return an error message
	if userExists {
		c.JSON(400, gin.H{"error": "Username is already taken"})
		return
	}

	//Bind the request body to the body struct
	if c.Bind(&body) != nil {
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	//Hash the password
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		c.JSON(500, gin.H{"error": "Error hashing password"})
		return
	}

	//Create a new user with the username and hashed password
	user := models.User{Username: body.Username, Password: string(hash)}
	result := initializers.DB.Create(&user)

	//Check for errors
	if result.Error != nil {
		c.JSON(500, gin.H{"error": "Error creating user"})
		return
	}

	//respond with the user
	c.JSON(200, gin.H{"user": user})
}
