package connection

import (
	"context"
	"fmt"

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
		return fmt.Errorf("could not connect to Redis: %v", err)
	} else {
		fmt.Println("redis connected successfully")
	}

	RedisClient = rdb
	Ctx = ctx

	return nil
}
