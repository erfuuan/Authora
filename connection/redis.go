package connection

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/redis/go-redis/v9"

	"github.com/erfuuan/Authora/conf"
)

var RedisClient *redis.Client
var Ctx context.Context

func InitRedis(cfg *conf.Config) error {
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.RedisAddress,
		Password: cfg.RedisPassword,
		DB:       0,
	})
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("❌ Redis connection failed: %v", err)
		os.Exit(1) // Ensures the app exits
	}

	fmt.Println("✅ Redis connected successfully")
	RedisClient = rdb
	Ctx = ctx

	return nil
}
