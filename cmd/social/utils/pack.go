package utils

import (
	"context"
	"errors"
	"github.com/hcdoit/tiktok/cmd/social/rpc"
	"github.com/hcdoit/tiktok/kitex_gen/user"
	"github.com/hcdoit/tiktok/pkg/errno"
	"github.com/hcdoit/tiktok/pkg/jwt"
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

func BuildUsers(ids []int64, myID int64, ctx context.Context) []*user.User {
	users := make([]*user.User, 0)
	token, err := jwt.CreateToken(jwt.CustomClaims{
		Id: int64(myID),
	})
	if err != nil {
		return nil
	}
	for _, id := range ids {
		user, err := rpc.GetUser(ctx, &user.GetUserRequest{
			UserId: id,
			Token:  token,
		})
		if user != nil && err == nil {
			users = append(users, user)
		}

	}
	return users
}
