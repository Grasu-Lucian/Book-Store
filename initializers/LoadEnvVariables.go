package initializers

import (
	"log"

	"github.com/joho/godotenv"
)

//Extracts the information from the  env file
func LoadEnvVariables() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
