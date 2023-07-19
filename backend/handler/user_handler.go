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
