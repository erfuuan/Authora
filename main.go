package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2/log"

	"github.com/erfuuan/Authora/conf"
	"github.com/erfuuan/Authora/connection"
	"github.com/erfuuan/Authora/internal/apiService"
	"github.com/erfuuan/Authora/internal/botHandler"
)

func main() {

	fmt.Println("Authora started...")
	cfg := conf.LoadConf()

	connection.InitDb(cfg)
	err := connection.InitRedis(cfg)
	if err != nil {
		log.Panic("❌ Failed to initialize Redis:", err)
	}

	go apiService.Init(cfg)
	botHandler.Init(cfg)
}
