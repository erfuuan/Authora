package botHandler

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"golang.org/x/net/proxy"

	"github.com/erfuuan/Authora/conf"
	"github.com/erfuuan/Authora/connection"
)

var bot *tgbotapi.BotAPI

func Init(cfg *conf.Config) {
	socks5Proxy := "socks5://127.0.0.1:25344"
	proxyURL, err := url.Parse(socks5Proxy)
	if err != nil {
		log.Fatal("Failed to parse proxy URL:", err)
	}

	dialer, err := proxy.SOCKS5("tcp", proxyURL.Host, nil, proxy.Direct)
	if err != nil {
		log.Fatal("Failed to create SOCKS5 dialer:", err)
	}

	transport := &http.Transport{
		Dial: func(network, addr string) (net.Conn, error) {
			return dialer.Dial(network, addr)
		},
	}

	httpClient := &http.Client{Transport: transport}

	bot, err = tgbotapi.NewBotAPIWithClient(cfg.BotToken, tgbotapi.APIEndpoint, httpClient)

	if err != nil {
		log.Fatal("Failed to create bot:", err)
	}

	bot.Debug = cfg.DebugMode

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for {
		for update := range updates {
			fmt.Println("====================================")
			// jsonData, err := json.MarshalIndent(update, "", "  ")
			// if err != nil {
			// 	log.Println("Error formatting JSON:", err)
			// }
			// fmt.Println(string(jsonData))
			fmt.Println("====================================")

			if update.Message != nil && update.Message.Chat != nil {
				if strings.HasPrefix(update.Message.Text, "/start") {
					HandleStart(update, bot)
					continue
				}

				value, _ := connection.RedisClient.Get(connection.Ctx, fmt.Sprintf("%s%d", "status_", update.Message.Chat.ID)).Result()
				if value != "" {
					switch value {
					case "wait_for_business_name":
						HandleFinishSignup(update, bot)
					case "wait_for_verify_token":
						HandleFinishVerifyMe(update, bot)
					default:
						SnedMsg(update.CallbackQuery.Message.Chat.ID, "not support yet")
					}
					continue
				}

				SnedMsg(update.Message.Chat.ID, "مثل آدم بنویس")
			}

			if update.CallbackQuery != nil {
				value, _ := connection.RedisClient.Get(connection.Ctx, fmt.Sprintf("%s%d", "status_", update.CallbackQuery.Message.Chat.ID)).Result()
				if value != "" && update.Message != nil && update.Message.Text != "" {
					switch value {
					case "wait_for_business_name":
						fmt.Printf("call  HandleFinishSignup")
						HandleFinishSignup(update, bot)
					case "wait_for_verify_token":
						HandleFinishVerifyMe(update, bot)
					default:
						fmt.Println("line default")
						SnedMsg(update.CallbackQuery.Message.Chat.ID, "not support yet")
					}
				} else {
					HandleButtonClick(update.CallbackQuery, bot)
				}
			}
		}
		time.Sleep(60 * time.Second)
	}
}
