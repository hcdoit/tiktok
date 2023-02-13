package utils

import (
	"context"
	"errors"
	"github.com/hcdoit/tiktok/cmd/user/dal/db"
	"github.com/hcdoit/tiktok/cmd/user/rpc"
	"github.com/hcdoit/tiktok/kitex_gen/social"
	"github.com/hcdoit/tiktok/kitex_gen/user"
	"github.com/hcdoit/tiktok/pkg/errno"
)

// BuildStatus 将error转换为StatusCode和StatusMsg
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

// BuildUser pack user info
func BuildUser(u *db.User, myID int64, ctx context.Context) *user.User {
	buildUser := &user.User{Id: int64(u.ID), Name: u.Username}
	relationInfo, err := rpc.GetRelationInfo(ctx, &social.RelationInfoRequest{
		UserId: int64(u.ID),
		MyId:   myID,
	})
	if err == nil {
		buildUser.FollowCount, buildUser.FollowerCount, buildUser.IsFollow = relationInfo.FollowCount, relationInfo.FollowerCount, relationInfo.IsFollow
	}
	if buildUser.Id == myID {
		buildUser.IsFollow = true
	}

	return buildUser
}
