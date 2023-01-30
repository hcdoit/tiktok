package utils

import (
	"errors"
	"github.com/hcdoit/tiktok/cmd/user/dal/db"
	"github.com/hcdoit/tiktok/kitex_gen/user"
	"github.com/hcdoit/tiktok/pkg/errno"
)

func BuildStatus(err error) (int32, string) {
	if err == nil {
		return errno.Success.ErrCode, errno.Success.ErrMsg
	}
	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return e.ErrCode, e.ErrMsg
	}
	s := errno.ServiceErr.WithMessage(err.Error())
	return s.ErrCode, s.ErrMsg
}

// User pack user info
func BuildUser(u *db.User) *user.User {
	if u == nil {
		return nil
	}

	return &user.User{Id: int64(u.ID), Name: u.Username}
}
