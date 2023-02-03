package rdb

import (
	"context"
	"github.com/hcdoit/tiktok/pkg/consts"
	"github.com/redis/go-redis/v9"
)

var RDB *redis.Client

func Init() {
	RDB = redis.NewClient(&redis.Options{
		Addr:     consts.RedisAddr,
		Password: consts.RedisPsw,
		DB:       consts.RedisDB,
	})
	ctx := context.Background()
	_, err := RDB.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}

}
