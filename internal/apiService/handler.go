package handler

import "github.com/gofiber/fiber"

func Ping(c fiber.Ctx) error {
	return c.SendString("hello world!")
}
