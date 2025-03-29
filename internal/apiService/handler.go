package apiService

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/erfuuan/Authora/connection"
	"github.com/erfuuan/Authora/internal/botHandler"
	"github.com/erfuuan/Authora/model"
)

func Ping(c fiber.Ctx) error {
	return c.SendString("pong!")
}

func SendOtp(c fiber.Ctx) error {
	var reqBody struct {
		Message string `json:"message" validate:"require,string"`
		UserId  string `json:"userId" validate:"require,string"`
	}
	if err := c.Bind().JSON(&reqBody); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	var userData model.User
	err := connection.DB.Where("user_id = ? ", reqBody.UserId).First(&userData).Error
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "not Found !", "statusCode": 404})
	}
	botHandler.SnedMsg(userData.ChatId, reqBody.Message)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "ok", "statusCode": 200})
}

func userIdVerify(c fiber.Ctx) error {

	var reqBody struct {
		UserId string `json:"userId" validate:"require,string"`
	}

	if err := c.Bind().JSON(&reqBody); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	var userData model.User

	result := connection.DB.Where("user_id = ? ", reqBody.UserId).First(&userData)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			fmt.Println(" user not found in DB , system generate token and send with bot for verify user and get chatId")
			token := uuid.NewString()
			err := connection.RedisClient.Set(connection.Ctx, "verify_token_"+token, reqBody.UserId, 120*time.Second).Err()

			if err != nil {
				fmt.Println(" Error saving to Redis:", err)
				return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"message": "Sorry, something went wrong!", "statusCode": 500})
			}
			return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "enter this to bot for verify : " + token, "statusCode": 201})
		}
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "user already exist", "statusCode": 200})
}
