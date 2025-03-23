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

	if token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Unauthorized", "statusCode": 401})
	}
	fmt.Println("🔹 Received Token:", token)

	value, _ := connection.RedisClient.Get(connection.Ctx, token).Result()
	if value == "" {
		fmt.Println("⚠️ Token not found, searching in DB...")
		var business model.Business
		result := connection.DB.Where("token = ?", token).First(&business)
		if result.Error != nil {
			fmt.Println("3")
			if result.Error == gorm.ErrRecordNotFound {
				fmt.Println("❌ Token not found in DB")
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Unauthorized", "statusCode": 401})
			}
		}

		jsonData, err := json.Marshal(business)
		if err != nil {
			fmt.Println("❌ Error marshalling JSON:", err)
			return fmt.Errorf("error marshalling JSON: %v", err)
		}

		// Save to Redis with expiration
		err = connection.RedisClient.Set(connection.Ctx, token, jsonData, 24*time.Hour).Err()
		if err != nil {
			fmt.Println("❌ Error saving to Redis:", err)
			return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"message": "Sorry, something went wrong!", "statusCode": 500})
		}
		fmt.Println("✅ Token successfully stored in Redis!")
	}

	return c.Next()

}

// if err != nil {
// 	fmt.Println("Token not found , lets find from db...")
// 	return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"messsage": "Sorry something went wrong!", "statusCode": "500"})
// }

// return c.Next()

// func IsLogin(c fiber.Ctx) error {
// 	tokenString := c.Get("Authorization")

// 	if tokenString == "" {
// 		return c.Status(fiber.StatusUnauthorized).JSON(responses.Error{
// 			Message: "Unauthorized",
// 		})
// 	}

// 	if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
// 		tokenString = tokenString[7:]
// 	}

// 	claims := jwt.MapClaims{}
// 	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
// 		return []byte(configs.SecretKey), nil
// 	})

// 	if err != nil || !token.Valid {
// 		return c.Status(fiber.StatusUnauthorized).JSON(responses.Error{
// 			Message: "Invalid token",
// 		})
// 	}

// 	exp, ok := claims["exp"].(float64)
// 	if !ok {
// 		return c.Status(fiber.StatusUnauthorized).JSON(responses.Error{
// 			Message: "Invalid token format",
// 		})
// 	}

// 	if time.Now().Unix() > int64(exp) {
// 		return c.Status(fiber.StatusUnauthorized).JSON(responses.Error{
// 			Message: "Token expired",
// 		})
// 	}

// 	var userId = claims["user_id"]

// 	c.Locals("user_id", userId)
// 	return c.Next()
// }

// func IsAdmin(c fiber.Ctx) error {
// 	tokenString := c.Get("Authorization")

// 	if tokenString == "" {
// 		return c.Status(fiber.StatusUnauthorized).JSON(responses.Error{
// 			Message: "Unauthorized",
// 		})
// 	}

// 	if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
// 		tokenString = tokenString[7:]
// 	}

// 	claims := jwt.MapClaims{}
// 	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
// 		return []byte(configs.SecretKey), nil
// 	})

// 	if err != nil || !token.Valid {
// 		return c.Status(fiber.StatusUnauthorized).JSON(responses.Error{
// 			Message: "Invalid token",
// 		})
// 	}

// 	exp, ok := claims["exp"].(float64)
// 	if !ok {
// 		return c.Status(fiber.StatusUnauthorized).JSON(responses.Error{
// 			Message: "Invalid token format",
// 		})
// 	}

// 	if time.Now().Unix() > int64(exp) {
// 		return c.Status(fiber.StatusUnauthorized).JSON(responses.Error{
// 			Message: "Token expired",
// 		})
// 	}

// 	var userId = claims["user_id"]

// 	var user database.User
// 	database.DB.First(&user, "id = ?", userId)

// 	if user.RoleID != 2 {
// 		return c.Status(fiber.StatusForbidden).JSON(responses.Error{
// 			Message: "forbidden — You are not admin user!",
// 		})
// 	}

// 	c.Locals("user_id", userId)
// 	return c.Next()
// }
