package rdb

import (
	"context"
	"github.com/hcdoit/tiktok/pkg/consts"
	"time"
)

func SaveToken(ctx context.Context, token string) error {
	key := token
	expire := time.Now().Add(consts.TokenExpireDuration)
	value := expire.Format(consts.TokenExpireFormat)
	err := RDB.Set(context.Background(), key, value, 0).Err()
	return err
}
