package main

import (
	"github.com/Oozaku/bug-bug-tracker/backend/database"
	"github.com/Oozaku/bug-bug-tracker/backend/database/models"
)

func main() {
	database.OpenDatabaseConnection()

	database.DBConnection.Create(&models.User{
		Username: "admin",
		Email:    "admin@email.com",
		Name:     "Bob Smith",
		Role:     "admin",
	})

	database.DBConnection.Create(&models.User{
		Username: "user",
		Email:    "user@email.com",
		Name:     "Ryan Lewis",
		Role:     "normal",
	})
}
