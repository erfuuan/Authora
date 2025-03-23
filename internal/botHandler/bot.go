package botHandler

import (
	"encoding/json"
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
)

func Init(cfg *conf.Config) {
	// Set up the SOCKS5 proxy
	socks5Proxy := "socks5://127.0.0.1:25344" // Change this to your proxy address
	proxyURL, err := url.Parse(socks5Proxy)
	if err != nil {
		log.Fatal("Failed to parse proxy URL:", err)
	}

	// Create a SOCKS5 dialer
	dialer, err := proxy.SOCKS5("tcp", proxyURL.Host, nil, proxy.Direct)
	if err != nil {
		log.Fatal("Failed to create SOCKS5 dialer:", err)
	}

	// Create a transport with the SOCKS5 dialer
	transport := &http.Transport{
		Dial: func(network, addr string) (net.Conn, error) {
			return dialer.Dial(network, addr)
		},
	}

	// Create an HTTP client with the transport
	httpClient := &http.Client{Transport: transport}

	// Initialize the Telegram bot with the proxy-configured HTTP client
	bot, err := tgbotapi.NewBotAPIWithClient(cfg.BotToken, tgbotapi.APIEndpoint, httpClient)
	if err != nil {
		log.Fatal("Failed to create bot:", err)
	}

	// Set bot debug mode
	// bot.Debug = true

	//?
	log.Printf("Authorized on account %s", bot.Self.UserName)

	// Start handling updates
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for {
		for update := range updates {
			fmt.Println(update)
			jsonData, err := json.MarshalIndent(update, "", "  ")
			if err != nil {
				log.Println("Error formatting JSON:", err)
			} else {
				fmt.Println(string(jsonData))
			}

			if strings.HasPrefix(update.Message.Text, "/signup") {
				HandleSignUp(update, bot)
			}

			if strings.HasPrefix(update.Message.Text, "/start") {
				bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "welcome to bot"))
			}

			if strings.HasPrefix(update.Message.Text, "/start") {
				bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "welcome to bot"))
			}

			if strings.HasPrefix(update.Message.Text, "/verifyOtpRequest") {
				verifyOtpRequest(update, bot)

			}

		}
		time.Sleep(3 * time.Second)
	}
}
