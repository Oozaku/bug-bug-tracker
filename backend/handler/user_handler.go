package handler

import (
	"errors"

	"github.com/Oozaku/bug-bug-tracker/backend/database"
	"github.com/Oozaku/bug-bug-tracker/backend/database/models"
	"github.com/Oozaku/bug-bug-tracker/backend/selector"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetAllUsers(c *fiber.Ctx) error {
	var users []models.User

	err := database.DBConnection.Model(&models.User{}).Preload("Histories").Find(&users).Error

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
	// TODO: get username
	// FIXME: protect me against SQL injection
	//        https://gorm.io/docs/security.html
	username := c.Params("username")

	var user models.User

	err := database.DBConnection.Model(&models.User{}).Preload("Histories").First(&user, "username = ?", username).Error

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
	// TODO: get username
	// FIXME: protect me against SQL injection
	//        https://gorm.io/docs/security.html
	username := "admin"

	var user models.User

	err := database.DBConnection.Model(&models.User{}).Preload("Histories").Find(&user, "username = ?", username).Error

	if err != nil {
		c.Status(500)
		return nil
	}

	c.JSON(selector.GetPrivateUserData(user))
	return nil
}
