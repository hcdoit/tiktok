package service

import (
	"context"
	"crypto/md5"
	"fmt"
	"github.com/hcdoit/tiktok/cmd/user/dal/db"
	"github.com/hcdoit/tiktok/cmd/user/utils"
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

// Login check user info
func (s *LoginService) Login(req *user.UserLoginRequest) (int64, string, error) {
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
		return 0, "", errno.AuthorizationFailedErr
	}
	u := users[0]
	if u.Password != passWord {
		return 0, "", errno.AuthorizationFailedErr
	}
	token, err := utils.Jwt.CreateToken(jwt.CustomClaims{
		Id: int64(u.ID),
	})
	if err != nil {
		return 0, "", err
	}
	return int64(u.ID), token, nil
}
