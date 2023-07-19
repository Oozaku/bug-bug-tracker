package handler

import (
	"github.com/Oozaku/bug-bug-tracker/backend/database"
	"github.com/Oozaku/bug-bug-tracker/backend/database/models"
	"github.com/gofiber/fiber/v2"
)

func GetAllUsers(c *fiber.Ctx) error {
	var users []models.UserPublic

	result := database.DBConnection.Model(&models.User{}).Find(&users)

	if result.Error != nil {
		c.Status(500)
		return nil
	}

	c.JSON(users)
	return nil
}

func GetSingleUser(c *fiber.Ctx) error {
	// TODO: get username
	// FIXME: protect me against SQL injection
	//        https://gorm.io/docs/security.html
	username := c.Params("username")

	var user models.UserPublic

	result := database.DBConnection.Model(&models.User{}).Find(&user, username)

	if result.Error != nil {
		c.Status(500)
		return nil
	}

	c.JSON(user)
	return nil
}

func GetUserHimself(c *fiber.Ctx) error {
	// TODO: get username
	// FIXME: protect me against SQL injection
	//        https://gorm.io/docs/security.html
	username := "admin"

	var user models.UserPrivate

	result := database.DBConnection.Model(&models.User{}).Find(&user, username)

	if result.Error != nil {
		c.Status(500)
		return nil
	}

	c.JSON(user)
	return nil
}
