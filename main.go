package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/url"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"golang.org/x/net/proxy"

	"github.com/erfuuan/Authora/connection"
)

func main() {
	connection.InitDb()
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
	bot, err := tgbotapi.NewBotAPIWithClient("token", tgbotapi.APIEndpoint, httpClient)
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
			if update.Message != nil {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Hello from proxy-enabled bot!")
				bot.Send(msg)
			}
		}
		time.Sleep(3 * time.Second)
	}
}
