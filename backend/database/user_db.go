package database

import (
	"errors"

	"github.com/Oozaku/bug-bug-tracker/backend/database/models"
	"github.com/Oozaku/bug-bug-tracker/backend/validator"
	"gorm.io/gorm"
)

func GetUser(username string) (models.User, error) {
	if err := validator.ValidateUsersUsername(username); err != nil {
		return models.User{}, err
	}

	var user models.User

	err := DBConnection.Model(&models.User{}).First(
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

func GetUsers(usernames []string) ([]models.User, error) {
	var users []models.User

	if usernames == nil {
		return users, nil
	}

	for _, username := range usernames {

		user, err := GetUser(username)

		if err != nil {
			return users, err
		}

		users = append(users, user)
	}

	return users, nil
}
