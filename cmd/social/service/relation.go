package service

import (
	"context"
	"github.com/hcdoit/tiktok/cmd/social/dal/rdb"
	"github.com/hcdoit/tiktok/cmd/social/utils"
	"github.com/hcdoit/tiktok/kitex_gen/social"
	"github.com/hcdoit/tiktok/kitex_gen/user"
)

type RelationService struct {
	ctx context.Context
}

// NewRelationService new RelationService
func NewRelationService(ctx context.Context) *RelationService {
	return &RelationService{
		ctx: ctx,
	}
}

// RelationAction 关注或取关
func (s *RelationService) RelationAction(req *social.RelationActionRequest, myID int64) error {
	if req.ToUserId == myID {
		return nil
	}
	if req.ActionType == 1 {
		return rdb.AddFollow(s.ctx, myID, req.ToUserId)
	}
	return rdb.CancelFollow(s.ctx, myID, req.ToUserId)
}

type Method int

const (
	Friend Method = iota
	Follow
	Follower
)

// RelationList 获取不同关系的用户列表
func (s *RelationService) RelationList(req *social.RelationListRequest, myID int64, method Method) ([]*user.User, error) {
	if method == Follower {
		ids, err := rdb.GetFollowerList(s.ctx, req.UserId)
		if err != nil {
			return nil, err
		}
		return utils.BuildUsers(ids, myID, s.ctx), nil
	}
	if method == Friend {
		ids, err := rdb.GetFriendList(s.ctx, req.UserId)
		if err != nil {
			return nil, err
		}
		return utils.BuildUsers(ids, myID, s.ctx), nil
	}
	if method == Follow {
		ids, err := rdb.GetFollowList(s.ctx, req.UserId)
		if err != nil {
			return nil, err
		}
		return utils.BuildUsers(ids, myID, s.ctx), nil
	}
	return nil, nil
}

// GetRelationInfo 获取用户的关系信息
func (s *RelationService) GetRelationInfo(req *social.RelationInfoRequest) (followCount int64, followerCount int64, isFollow bool, err error) {
	followCount, err = rdb.GetFollowCount(s.ctx, req.UserId)
	if err != nil {
		return 0, 0, false, err
	}
	followerCount, err = rdb.GetFollowerCount(s.ctx, req.UserId)
	if err != nil {
		return 0, 0, false, err
	}
	isFollow, err = rdb.CheckFollow(s.ctx, req.MyId, req.UserId)
	if err != nil {
		return 0, 0, false, err
	}
	return followCount, followerCount, isFollow, nil
}
