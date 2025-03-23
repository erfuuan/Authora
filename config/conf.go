package conf

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerPort string
	BotToken   string
	DBPort     string
	DBHost     string
	DBPassword string
	DBName     string
	DBUser     string
	SslMode    string
}

func loadConf() *Config {

	if err := godotenv.Load(".env"); err != nil {
		log.Println("Warning: No .env file found, using system environment variables.")
	}

	return &Config{
		ServerPort: os.Getenv("PORT"),
		BotToken:   os.Getenv("BOT_TOKEN"),
		DBPort:     os.Getenv("DB_PORT"),
		DBHost:     os.Getenv("DB_HOST"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
		DBUser:     os.Getenv("DB_USER"),
		SslMode:    os.Getenv("SSL_MODE"),
	}

}
