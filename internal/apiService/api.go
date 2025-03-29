package apiService

import (
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/logger"

	"github.com/erfuuan/Authora/conf"
	"github.com/erfuuan/Authora/middlewares"
)

func Init(cfg *conf.Config) {
	app := fiber.New(fiber.Config{})

	app.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST"},
		AllowHeaders: []string{"*"},
	}))

	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))

	api := app.Group("/api/v1/authora", middlewares.AuthApi)
	Router(api)

	port := cfg.Port
	err := app.Listen(port)
	if err != nil {
		log.Fatal("Error starting server: ", err)
	}
	log.Println("✅ Application started successfully! 🚀")
}
