package handlers

import (
	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
)

func templHandler(c templ.Component) func(*fiber.Ctx) error {
	return adaptor.HTTPHandler(templ.Handler(c))
}