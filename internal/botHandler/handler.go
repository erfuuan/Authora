package botHandler

import (
	"fmt"
	"log"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/erfuuan/Authora/connection"
	"github.com/erfuuan/Authora/model"
)

func HandleStart(update tgbotapi.Update, bot *tgbotapi.BotAPI) {

	startMsg := `🤖 Welcome to Authora Bot!

I'm here to help you with authentication and verification. Choose an option below to proceed:

🔹 **📌 Sign Up** – Register your business to start using our authentication system.
🔹 **🔍 Verify Me** – Verify your identity quickly and securely.
🔹 **ℹ️ Help** – You're here! Need assistance? Just follow the instructions below.

📢 **How to Use?**
1️⃣ If you're a business, tap **📌 Sign Up** to get started.
2️⃣ If you're a user, tap **🔍 Verify Me** to authenticate.
3️⃣ Need help? Just tap **ℹ️ Help** anytime!

For more details, feel free to reach out! 🚀`

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, startMsg)
	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("📌 Sign Up", "start_signup"),
			tgbotapi.NewInlineKeyboardButtonData("🔍 Verify Me", "verify_me"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("ℹ️ Help", "show_help"),
		),
	)
	bot.Send(msg)
}

func HandleHelpButton(update tgbotapi.Update, bot *tgbotapi.BotAPI) {

	helpMessage := `ℹ️ *Authora Bot Help Guide*

	Welcome to *Authora Bot*! 🤖 I'm here to help you with authentication and verification.
	
	🔹 *Sign Up 📌* – If you're a business, tap *📌 Sign Up* to register.
	🔹 *Verify Me 🔍* – If you're a user, tap *🔍 Verify Me* to authenticate.
	🔹 *Help ℹ️* – You’re here! Need assistance? Check out the steps below.
	
	📢 *How It Works?*
	1️⃣ *Business Owners*: Use *Sign Up* to register your business.
	2️⃣ *Users*: Use *Verify Me* to confirm your identity.
	3️⃣ *Need Assistance?* Just tap *Help* anytime!
	
	For additional support, feel free to reach out! 🚀`
	SnedMsg(update.Message.Chat.ID, helpMessage)

}

func HandleButtonClick(query *tgbotapi.CallbackQuery, bot *tgbotapi.BotAPI) {
	fmt.Println(query.Data)
	switch query.Data {
	case "start_signup":
		HandleStartSignupButton(tgbotapi.Update{Message: query.Message}, bot)
	case "verify_me":
		HandleStartVerifyButton(tgbotapi.Update{Message: query.Message}, bot)
	case "show_help":
		HandleHelpButton(tgbotapi.Update{Message: query.Message}, bot)

	default:
		SnedMsg(query.Message.Chat.ID, "not supported")
	}
}

func HandleStartSignupButton(update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "💼 Please enter your Business Name:")
	msg.ReplyToMessageID = update.Message.MessageID
	err := connection.RedisClient.Set(connection.Ctx, fmt.Sprintf("%s%d", "status_", update.Message.Chat.ID), "wait_for_business_name", 24*time.Hour).Err()
	if err != nil {
		fmt.Println("❌ Error saving to Redis:", err)
	}
	bot.Send(msg)
}

func HandleFinishSignup(update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	var business model.Business
	result := connection.DB.Where("name = ?", update.Message.Text).First(&business)
	if result.Error == gorm.ErrRecordNotFound {
		fmt.Println("Business not found, lets create.")
		token := uuid.NewString()
		business := model.Business{
			Token: token,
			Name:  update.Message.Text,
		}
		err := connection.DB.Create(&business).Error
		if err != nil {
			fmt.Println("failed to create : ", err)
			bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "failed to create , please try again later"))
		}
		bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("Successfully signed up! Your token is: %s", token)))
		connection.RedisClient.Del(connection.Ctx, fmt.Sprintf("%s%d", "status_", update.Message.Chat.ID))
	} else {
		SnedMsg(update.Message.Chat.ID, "already exist")
	}

}

func HandleStartVerifyButton(update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	var user model.User
	checkExist := connection.DB.Where("chat_id = ?", update.Message.Chat.ID).First(&user)
	if checkExist.Error == gorm.ErrRecordNotFound {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Please enter your token : ")
		msg.ReplyToMessageID = update.Message.MessageID
		err := connection.RedisClient.Set(connection.Ctx, fmt.Sprintf("%s%d", "status_", update.Message.Chat.ID), "wait_for_verify_token", 24*time.Hour).Err()
		if err != nil {
			fmt.Println("❌ Error saving to Redis:", err)
		}
		bot.Send(msg)
	} else {
		SnedMsg(update.Message.Chat.ID, "you are already verified")
		connection.RedisClient.Del(connection.Ctx, fmt.Sprintf("%s%d", "status_", update.Message.Chat.ID))

	}
}

func SnedMsg(chatId int64, msg string) error {

	_, err := bot.Send(tgbotapi.NewMessage(chatId, msg))
	if err != nil {
		log.Fatal("err for send msg", err)
	}
	return err
}

func HandleFinishVerifyMe(update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	var user model.User

	checkExist := connection.DB.Where("chat_id = ?", update.Message.Chat.ID).First(&user)
	if checkExist.Error == gorm.ErrRecordNotFound {
		value, _ := connection.RedisClient.Get(connection.Ctx, "verify_token_"+update.Message.Text).Result()
		if value != "" {
			user = model.User{
				UserId: value,
				ChatId: update.Message.Chat.ID,
			}
			err := connection.DB.Create(&user).Error
			if err != nil {
				fmt.Println("failed to create : ", err)
				SnedMsg(update.Message.Chat.ID, "failed to create , please try again later")
			}
			SnedMsg(update.Message.Chat.ID, "successfull verified")

		} else {
			SnedMsg(update.Message.Chat.ID, "token is invalid")
		}
	} else {
		SnedMsg(update.Message.Chat.ID, "you are already verified")
		connection.RedisClient.Del(connection.Ctx, fmt.Sprintf("%s%d", "status_", update.Message.Chat.ID))
	}
}
