package main

import (
	"github.com/gin-gonic/gin"
	"main/initializers"
)

func init() {
	initializers.LoadEnvVariables()
}

// This function initializes a Gin router with default middleware.
// It defines a GET route on "/ping" that responds with a JSON message "pong".
// Finally, it starts the server and listens on port 3000.
func main() {

	// Initialize a default Gin router
	r := gin.Default()

	// Define a route for GET requests on "/ping"
	r.GET("/ping", func(c *gin.Context) {

		// Respond with JSON containing a message
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// Start the HTTP server and listen on port 3000
	r.Run()
}
