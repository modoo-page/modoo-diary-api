package server

import "github.com/gofiber/fiber/v2"

func Create() *fiber.App {
	app := fiber.New()
	return app
}
