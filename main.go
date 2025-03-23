package main

import (
	"fmt"
	"log"

	"github.com/erfuuan/Authora/conf"
	"github.com/erfuuan/Authora/connection"
	"github.com/erfuuan/Authora/internal/apiService"
	"github.com/erfuuan/Authora/internal/botHandler"
)

func main() {

	cfg := conf.LoadConf()

	connection.InitDb(cfg)
	connection.InitRedis(cfg)

	go apiService.Init(cfg)
	botHandler.Init(cfg)

	fmt.Println(cfg)

	log.Println("✅ Application started successfully! 🚀")
}
