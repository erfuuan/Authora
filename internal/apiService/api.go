package apiService

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v3/middleware/cors"

	"github.com/erfuuan/Authora/conf"
)

func Init(cfg *conf.Config) {
	app := fiber.New(fiber.Config{})

	app.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders: []string{"*"},
	}))
	port := cfg.Port
	app.Listen(port)
	log.Println("✅ Application started successfully! 🚀")

}
