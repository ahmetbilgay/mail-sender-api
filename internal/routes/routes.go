package routes

import (
	"blgy-mail-sender/internal/controllers"
	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App) {
	app.Post("/sendMail", controllers.SendMail)
}
