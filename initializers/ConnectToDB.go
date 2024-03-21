package initializers

// Import necessary packages: "gorm.io/driver/postgres" for PostgreSQL driver and "gorm.io/gorm" for the GORM ORM.
import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

// Declare a global variable to hold the GORM database object.
var DB *gorm.DB

// ConnectToDB initializes a connection to a PostgreSQL database using GORM.
func ConnectToDB() {
	var err error
	//postgres://:@/xpsoxguy
	// Define the Data Source Name (DSN) for connecting to the PostgreSQL database.
	dsn := os.Getenv("DB")

	// Open a connection to the PostgreSQL database using GORM.
	// The gorm.Open() function takes the PostgreSQL driver and DSN as parameters.
	// It returns a GORM database object (db) and an error (err).
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	// If an error occurs during the database connection, panic with an error message.
	if err != nil {
		panic("failed to connect to db")
	} 
}
