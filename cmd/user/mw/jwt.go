package mw

import (
	"github.com/hcdoit/tiktok/pkg/consts"
	"github.com/hcdoit/tiktok/pkg/jwt"
)

var (
	Jwt *jwt.JWT
)

func Init() {
	Jwt = jwt.NewJWT([]byte(consts.SecretKey))
}
