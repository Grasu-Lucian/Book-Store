package initializers

import "BOOK-STORE/models"
// SyncDatabase is a function that syncs the database
func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}