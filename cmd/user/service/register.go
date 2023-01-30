package service

import (
	"context"
	"crypto/md5"
	"fmt"
	"github.com/hcdoit/tiktok/cmd/user/dal/db"
	"github.com/hcdoit/tiktok/kitex_gen/user"
	"github.com/hcdoit/tiktok/pkg/errno"
	"io"
)

type RegisterService struct {
	ctx context.Context
}

func NewRegisterService(ctx context.Context) *RegisterService {
	return &RegisterService{ctx: ctx}
}

func (s *RegisterService) Register(req *user.UserRegisterRequest) error {
	users, err := db.QueryUserByName(s.ctx, req.Username)
	if err != nil {
		return err
	}
	if len(users) != 0 {
		return errno.UserAlreadyExistErr
	}
	h := md5.New()
	if _, err = io.WriteString(h, req.Password); err != nil {
		return err
	}
	password := fmt.Sprintf("%x", h.Sum(nil))
	return db.CreateUser(s.ctx, []*db.User{{
		Username: req.Username,
		Password: password,
	}})
}
