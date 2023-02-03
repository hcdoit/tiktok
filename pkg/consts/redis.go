package consts

import "time"

const (
	RedisAddr             = "localhost:6379"
	RedisPsw              = "123456"
	RedisDB               = 0
	TokenExpireFormat     = time.RFC3339
	TokenExpireDuration   = time.Duration(24 * time.Hour)
	ForeverExpireDuration = -1
)
