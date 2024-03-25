package main

import (
	"github.com/gin-gonic/gin"
	"main/controllers"
	"main/initializers"
	"main/middleware"
)

// This function is called before the main function
func init() {
	//Extracts the information from the  env file
	initializers.LoadEnvVariables()
	//Makes the connection to the database
	initializers.ConnectToDB()
	//Syncs the database
	initializers.SyncDatabase()
}

// This function initializes a Gin router with default middleware.
// It then defines the routes for the application.
func main() {
	// Initialize a default Gin router
	r := gin.Default()
	// Register the user routes
	r.POST("/register", controllers.Signup)
	// Register the login route
	r.POST("/login", controllers.Login)
	// Post a book
	r.POST("/book", middleware.RequireAuth, controllers.BookPost)
	// Get a book
	r.GET("/book/:id", controllers.BookGet)
	// Get all books
	r.GET("/books", controllers.BookGetAll)
	// Update a book
	r.PUT("/book/:id", middleware.RequireAuth, controllers.BookUpdate)
	// Delete a book
	r.DELETE("/book/:id", middleware.RequireAuth, controllers.BookDelete)
	// Start the HTTP server and listen on port 3000
	r.Run()
}
