package mw

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	hzconsts "github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/hcdoit/tiktok/pkg/consts"
	"github.com/hcdoit/tiktok/pkg/errno"
	"github.com/redis/go-redis/v9"
	"time"
)

var RDB *redis.Client

func InitCheckToken() {
	RDB = redis.NewClient(&redis.Options{
		Addr:     consts.RedisAddr,
		Password: consts.RedisPsw,
		DB:       consts.RedosDB,
	})
	ctx := context.Background()
	_, err := RDB.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}

}

func CheckToken(ctx context.Context, token string) (bool, error) {
	value, err := RDB.Get(ctx, token).Result()
	if err != nil {
		return false, err
	}
	expire, err := time.Parse(consts.TokenExpireFormat, value)
	if err != nil {
		return false, err
	}

	return time.Now().Before(expire), nil

}

func TokenMiddlewareFunc() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		ok, _ := CheckToken(ctx, c.Query("token"))
		if !ok {
			c.JSON(hzconsts.StatusUnauthorized, errno.AuthInvalidJwt)
			c.Abort()
		}
		c.Next(ctx)
	}
}
