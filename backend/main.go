package main

import (
	"github.com/Oozaku/bug-bug-tracker/backend/database"
	"github.com/Oozaku/bug-bug-tracker/backend/router"
	"github.com/gofiber/fiber/v2"
)

func main() {
	database.OpenDatabaseConnection()

	app := fiber.New()
	router.SetupRoutes(app)

	app.Listen(":3000")
}
