package utils

import (
	"github.com/hcdoit/tiktok/pkg/consts"
	"github.com/hcdoit/tiktok/pkg/jwt"
)

var (
	Jwt *jwt.JWT
)

func JwtInit() {
	Jwt = jwt.NewJWT([]byte(consts.SecretKey))
}
