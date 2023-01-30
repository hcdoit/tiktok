package service

import (
	"context"
	"github.com/hcdoit/tiktok/cmd/user/dal/db"
	"github.com/hcdoit/tiktok/cmd/user/utils"
	"github.com/hcdoit/tiktok/kitex_gen/user"
	"github.com/hcdoit/tiktok/pkg/errno"
)

type GetUserService struct {
	ctx context.Context
}

func NewGetUserService(ctx context.Context) *GetUserService {
	return &GetUserService{ctx: ctx}
}

func (s *GetUserService) GetUser(req *user.GetUserRequest, myID int64) (*user.User, error) {
	users, err := db.QueryUserByID(s.ctx, myID)
	if err != nil {
		return nil, err
	}
	if len(users) != 0 {
		return nil, errno.ServiceErr
	}
	return utils.BuildUser(users[0]), nil
}
