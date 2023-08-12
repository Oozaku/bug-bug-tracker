package handler

import (
	"errors"

	"github.com/Oozaku/bug-bug-tracker/backend/database"
	"github.com/Oozaku/bug-bug-tracker/backend/database/models"
	"github.com/Oozaku/bug-bug-tracker/backend/selector"
	"github.com/Oozaku/bug-bug-tracker/backend/validator"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetAllUsers(c *fiber.Ctx) error {
	var users []models.User

	err := database.DBConnection.Model(&models.User{}).Find(&users).Error

	if err != nil {
		c.Status(500)
		return nil
	}

	var publicUsers []selector.UserPublicApi

	for _, user := range users {
		publicUsers = append(publicUsers, selector.GetPublicUserData(user))
	}

	c.JSON(publicUsers)
	return nil
}

func GetSingleUser(c *fiber.Ctx) error {
	username := c.Params("username")

	if err := validator.ValidateUsersUsername(username); err != nil {
		c.Status(400).SendString("Invalid username")
		return nil
	}

	var user models.User

	err := database.DBConnection.Model(&models.User{}).First(&user, "username = ?", username).Error

	if err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.Status(404).SendString("User not found")
			return nil
		}

		c.Status(500)
		return nil
	}

	c.JSON(selector.GetPublicUserData(user))
	return nil
}

func GetUserHimself(c *fiber.Ctx) error {
	username := "admin"

	var user models.User

	err := database.DBConnection.Model(&models.User{}).Find(&user, "username = ?", username).Error

	if err != nil {
		c.Status(500)
		return nil
	}

	c.JSON(selector.GetPrivateUserData(user))
	return nil
}
