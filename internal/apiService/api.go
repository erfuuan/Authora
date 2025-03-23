package apiService

import (
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"

	"github.com/erfuuan/Authora/conf"
	"github.com/erfuuan/Authora/middlewares"
)

func Init(cfg *conf.Config) {
	app := fiber.New(fiber.Config{})
	// Middleware: CORS configuration
	app.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders: []string{"*"},
	}))
	// Define routes
	api := app.Group("/api/v1/authora", middlewares.AuthApi)
	Router(api)
	// Get the port from the config and start the server
	port := cfg.Port
	err := app.Listen(port)
	if err != nil {
		log.Fatal("Error starting server: ", err)
	}
}
