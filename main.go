package main

import (
	"main/controllers"
	"main/initializers"
	"main/middleware"
	"github.com/gin-gonic/gin"
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

	// Define a route for GET requests on "/ping"
	r.POST("/register", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.GET("/validate",middleware.RequireAuth  , controllers.ValidateToken)
	// Start the HTTP server and listen on port 3000
	r.Run()
}
