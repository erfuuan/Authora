package conf

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Port          string
	DBUser        string
	DBPassword    string
	DBName        string
	DBHost        string
	DBPort        string
	BotToken      string
	RedisAddress  string
	RedisPassword string
	SslMode       string
	DebugMode     bool
}

func LoadConf() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: No .env file found, using system environment variables")
	}

	debugModeStr := os.Getenv("DEBUG_MODE")
	debugMode, err := strconv.ParseBool(debugModeStr)
	if err != nil {
		log.Printf("Warning: Invalid boolean value for DEBUG_MODE (%s), defaulting to false", debugModeStr)
		debugMode = false
	}

	return &Config{
		Port:          os.Getenv("PORT"),
		DBUser:        os.Getenv("DB_USER"),
		DBPassword:    os.Getenv("DB_PASSWORD"),
		DBName:        os.Getenv("DB_NAME"),
		DBHost:        os.Getenv("DB_HOST"),
		DBPort:        os.Getenv("DB_PORT"),
		BotToken:      os.Getenv("BOT_TOKEN"),
		RedisAddress:  os.Getenv("REDIS_ADDRESS"),
		RedisPassword: os.Getenv("REDIS_PASSWORD"),
		SslMode:       os.Getenv("SSL_MODE"),
		DebugMode:     debugMode,
	}
}
