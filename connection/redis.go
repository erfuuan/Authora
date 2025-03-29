package connection

import (
	"context"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"

	"github.com/erfuuan/Authora/conf"
)

var (
	RedisClient *redis.Client
	Ctx         context.Context
)

func InitRedis(cfg *conf.Config) error {
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.RedisAddress,
		Password: cfg.RedisPassword,
		DB:       0,
	})
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Panic("❌ Redis connection failed: ", err) // This should stop the app if Redis is not working
	}
	fmt.Println("✅ Redis connected successfully")
	RedisClient = rdb
	Ctx = ctx
	return nil
}
