package routes

import (
	"app.yourcompany.com/handlers"

	"github.com/gofiber/fiber/v2"
)

func RegisterWeb(web fiber.Router) {

	web.Get("/", handlers.HandleLandingIndex())

	// Handle 404
	web.Get("/*", func(c *fiber.Ctx) error {
		return c.SendStatus(404)
	})
}
