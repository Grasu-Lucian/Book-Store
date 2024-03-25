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
// It defines a GET route on "/ping" that responds with a JSON message "pong".
// Finally, it starts the server and listens on port 3000.
func main() {
	// Initialize a default Gin router
	r := gin.Default()

	r.POST("/register", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middleware.RequireAuth, controllers.ValidateToken)
	r.POST("/book", middleware.RequireAuth, controllers.BookPost)
	r.GET("/book/:id", controllers.BookGet)
	r.GET("/books", controllers.BookGetAll)
	r.PUT("/book/:id", middleware.RequireAuth, controllers.BookUpdate)
	// Start the HTTP server and listen on port 3000
	r.Run()
}
