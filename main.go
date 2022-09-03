package main

import (
	"blgy-mail-sender/internal/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	routes.Routes(app)
	app.Listen(":3000")

}
