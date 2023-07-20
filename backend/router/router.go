package router

import (
	"github.com/Oozaku/bug-bug-tracker/backend/handler"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/users", handler.GetAllUsers)
	app.Get("/users/me", handler.GetUserHimself)
	app.Get("/users/:username", handler.GetSingleUser)
}
