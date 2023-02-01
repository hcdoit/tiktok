package mw

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
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
		DB:       consts.RedisDB,
	})
	ctx := context.Background()
	_, err := RDB.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}

}

func CheckToken(ctx context.Context, token string) (bool, error) {
	value, err := RDB.Get(ctx, token).Result()
	if err == redis.Nil {
		hlog.Debug("no token in redis")
		return false, err
	}
	if err != nil {
		hlog.Error("redis error")
		return false, err
	}
	expire, err := time.Parse(consts.TokenExpireFormat, value)
	if err != nil {
		hlog.Error("parse token expire error")
		return false, err
	}

	return time.Now().Before(expire), nil

}

func TokenMiddlewareFunc() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		token := c.Query("token")
		if len(token) == 0 {
			token = c.PostForm("token")
		}
		if len(token) == 0 {
			hlog.Debug("no token in request")
			c.JSON(hzconsts.StatusOK, errno.AuthInvalidJwt)
			c.Abort()
		}

		ok, err := CheckToken(ctx, token)
		if !ok {
			if err == nil {
				hlog.Debug("token expired")
			}
			c.JSON(hzconsts.StatusOK, errno.AuthInvalidJwt)
			c.Abort()
		}
		c.Next(ctx)
	}
}
