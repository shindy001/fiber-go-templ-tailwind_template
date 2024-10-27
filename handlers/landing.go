package handlers

import (
	"app.yourcompany.com/views/landing"

	"github.com/gofiber/fiber/v2"
)

func HandleLandingIndex() fiber.Handler {
	return templHandler(landing.Index())
}

