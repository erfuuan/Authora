package apiRoute

import (
	"github.com/erfuuan/Authora/handler"
	"github.com/gofiber/fiber/v3"
)

func Router(app *fiber.App) {

	app.Get("/ping", handler.Ping)
}
