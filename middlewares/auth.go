package middlewares

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"

	"github.com/erfuuan/Authora/connection"
	"github.com/erfuuan/Authora/model"
)

func AuthApi(c fiber.Ctx) error {
	token := c.Get("Authorization")

	if c.Path() == "/api/v1/authora/ping" && c.Method() == "GET" {
		return c.Next()
	}

	if token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Unauthorized", "statusCode": 401})
	}

	value, _ := connection.RedisClient.Get(connection.Ctx, token).Result()
	if value == "" {
		fmt.Println("⚠️ Token not found, searching in DB...")
		var business model.Business
		result := connection.DB.Where("token = ?", token).First(&business)
		if result.Error != nil {
			if result.Error == gorm.ErrRecordNotFound {
				fmt.Println("❌ Token not found in DB")
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Unauthorized", "statusCode": 401})
			}
		}

		jsonData, err := json.Marshal(business)
		if err != nil {
			fmt.Println("❌ Error marshalling JSON:", err)
		}

		err = connection.RedisClient.Set(connection.Ctx, token, jsonData, 24*time.Hour).Err()
		if err != nil {
			fmt.Println("❌ Error saving to Redis:", err)
			return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"message": "Sorry, something went wrong!", "statusCode": 500})
		}
		fmt.Println("✅ Token successfully stored in Redis!")
	}
	return c.Next()
}
