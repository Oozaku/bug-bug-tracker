package database

import (
	"fmt"
	"log"

	"github.com/Oozaku/bug-bug-tracker/backend/database/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DBConnection *gorm.DB

// TODO: receive connection values from environment
func OpenDatabaseConnection() {
	if DBConnection != nil {
		return
	}

	host := "localhost"
	port := "5432"
	username := "postgres"
	password := "postgres"
	databaseName := "postgres"
	timezone := "America/Sao_Paulo"

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s",
		host,
		username,
		password,
		databaseName,
		port,
		timezone,
	)

	var err error

	DBConnection, err = gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})

	if err != nil {
		log.Fatal("Could not connect with database: ", err)
	}

	RunMigration(DBConnection)
}

func RunMigration(dbConnection *gorm.DB) {
	error := dbConnection.AutoMigrate(&models.Issue{}, &models.User{}, &models.History{})

	if error != nil {
		log.Fatal("Could not run migration: ", error)
	}
}
