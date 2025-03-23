package botHandler

import (
	"fmt"
	"log"
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

func SnedMsg(chatId int64, msg string) error {

	_, err := bot.Send(tgbotapi.NewMessage(chatId, msg))
	if err != nil {
		log.Fatal("err for send msg", err)
	}
	return err
}

func HandleUserVerify(msg tgbotapi.Update, bot *tgbotapi.BotAPI) {

	token := strings.TrimSpace(strings.TrimPrefix(msg.Message.Text, "/user-verify"))
	if token == "" {
		bot.Send(tgbotapi.NewMessage(msg.Message.Chat.ID, "Please send valid token : //user-verify <token>"))
		return
	}

	fmt.Println(token)
	value, _ := connection.RedisClient.Get(connection.Ctx, token).Result()

	if value == "" {
		bot.Send(tgbotapi.NewMessage(msg.Message.Chat.ID, "hey im joke for you ? "))
		return
	}

	user := model.User{
		UserId: value,
		ChatId: msg.Message.Chat.ID,
	}
	err := connection.DB.Create(&user).Error
	if err != nil {
		fmt.Println("failed to create : ", err)
		bot.Send(tgbotapi.NewMessage(msg.Message.Chat.ID, "failed to create , please try again later"))
		return
	}
	bot.Send(tgbotapi.NewMessage(msg.Message.Chat.ID, "Successfully user verified"))

}
