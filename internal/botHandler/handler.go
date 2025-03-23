package botHandler

import (
	"fmt"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/google/uuid"

	"github.com/erfuuan/Authora/connection"
	"github.com/erfuuan/Authora/model"
)

func HandleSignUp(update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	name := strings.TrimSpace(strings.TrimPrefix(update.Message.Text, "/signup"))
	if name == "" {
		bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Please provide a business name. Usage: /signup <BusinessName>"))
		return
	}
	token := uuid.NewString()
	business := model.Business{
		Token: token,
		Name:  name,
	}
	err := connection.DB.Create(&business).Error
	if err != nil {
		fmt.Println("failed to create : ", err)
		bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "failed to create , please try again later"))
	}
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("Successfully signed up! Your token is: %s", token))
	bot.Send(msg)
}

func verifyOtpRequest(update tgbotapi.Update, bot *tgbotapi.BotAPI) {

}
