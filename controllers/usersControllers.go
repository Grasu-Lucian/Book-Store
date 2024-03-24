package controllers

import (
	"BOOK-STORE/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"main/initializers"
	"net/http"
	"os"
	"time"
)

func Signup(c *gin.Context) {

	//get username and password from the request body
	var body struct {
		Username string
		Password string
	}

	//Bind the request body to the body struct
	if c.Bind(&body) != nil {
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	// Query the database for a user with the given username
	var userModel models.User
	userExists := initializers.DB.Where("username = ?", body.Username).First(&userModel).RowsAffected > 0

	// If the user exists, return an error message
	if userExists {
		c.JSON(400, gin.H{"error": "Username is already taken"})
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

func Login(c *gin.Context) {

	//get username and password from the request body
	var body struct {
		Username string
		Password string
	}
	//print("body")

	//Bind the request body to the body struct
	if c.Bind(&body) != nil {
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	//look up the user in the database
	var userModel models.User
	initializers.DB.Where("username = ?", body.Username).First(&userModel)
	if userModel.ID == 0 {
		c.JSON(400, gin.H{"error": "User not found"})
		return
	}
	//compare sent in password with the hashed password in the database
	err := bcrypt.CompareHashAndPassword([]byte(userModel.Password), []byte(body.Password))

	//if the passwords don't match, return an error
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid password"})

		return
	}

	//generate a JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userModel.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		c.JSON(500, gin.H{"error": "Error generating token"})
		return
	}

	//send the token back to the user
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)
	c.JSON(200, gin.H{"token": tokenString})
}
