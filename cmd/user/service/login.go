package service

import (
	"context"
	"crypto/md5"
	"fmt"
	"github.com/hcdoit/tiktok/cmd/user/dal/db"
	"github.com/hcdoit/tiktok/cmd/user/dal/rdb"
	"github.com/hcdoit/tiktok/kitex_gen/user"
	"github.com/hcdoit/tiktok/pkg/errno"
	"github.com/hcdoit/tiktok/pkg/jwt"
	"io"
)

type LoginService struct {
	ctx context.Context
}

// NewLoginService new LoginService
func NewLoginService(ctx context.Context) *LoginService {
	return &LoginService{
		ctx: ctx,
	}
}

// Login 检查账号密码返回id和token
func (s *LoginService) Login(req *user.UserLoginRequest) (int64, string, error) {
	// 明文密码加密后与数据库比对
	h := md5.New()
	if _, err := io.WriteString(h, req.Password); err != nil {
		return 0, "", err
	}
	passWord := fmt.Sprintf("%x", h.Sum(nil))
	userName := req.Username
	users, err := db.QueryUserByName(s.ctx, userName)
	if err != nil {
		return 0, "", err
	}
	if len(users) == 0 {
		return 0, "", errno.AuthInvalidAccount
	}
	u := users[0]
	// 与数据库比较
	if u.Password != passWord {
		return 0, "", errno.AuthInvalidAccount
	}
	// 创建Token，并存入缓存表示登录
	token, err := jwt.CreateToken(jwt.CustomClaims{
		Id: int64(u.ID),
	})
	if err != nil {
		return 0, "", err
	}
	err = rdb.SaveToken(s.ctx, token)
	if err != nil {
		return 0, "", err
	}
	return int64(u.ID), token, nil
}
