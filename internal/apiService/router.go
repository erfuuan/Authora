package apiService

import (
	"github.com/gofiber/fiber/v3"
)

func Router(router fiber.Router) {
	router.Get("/ping", Ping)
	router.Post("/send-otp", SendOtp)

	router.Post("/user-verify", userIdVerify)

}
