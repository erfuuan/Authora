package conf

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config struct to hold configuration values
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
}

// LoadConfig loads configuration from environment variables or .env file
func LoadConf() *Config {
	err := godotenv.Load() // Load the .env file if available
	if err != nil {
		log.Println("Warning: No .env file found, using system environment variables")
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
	}
}
