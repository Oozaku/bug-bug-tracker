package database

import (
	"errors"

	"github.com/Oozaku/bug-bug-tracker/backend/database/models"
	"github.com/Oozaku/bug-bug-tracker/backend/validator"
	"gorm.io/gorm"
)

func getUser(username string) (models.User, error) {
	if err := validator.ValidateUsersUsername(username); err != nil {
		return models.User{}, err
	}

	var user models.User

	err := DBConnection.Model(&models.User{}).Preload("Histories").First(
		&user, "username = ?", username,
	).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.User{}, errors.New("User not found")
		}

		return models.User{}, errors.New("Database is down")
	}

	return user, nil
}

func getUsers(usernames []string) ([]models.User, error) {
	var users []models.User

	for _, username := range usernames {
		user, err := getUser(username)

		if err != nil {
			return users, err
		}

		users = append(users, user)
	}

	return users, nil
}
