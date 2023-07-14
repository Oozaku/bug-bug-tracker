package database

import (
	"log"

	"gorm.io/gorm"
)

func RunMigration(dbConnection *gorm.DB) {
	error := dbConnection.AutoMigrate(&models.Issue{}, &models.User{}, &models.History{})

	if error != nil {
		log.Fatal("Could not run migration: ", error)
	}
}
