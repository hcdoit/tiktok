package main

import (
	"context"
	"github.com/hcdoit/tiktok/cmd/social/service"
	"github.com/hcdoit/tiktok/cmd/social/utils"
	social "github.com/hcdoit/tiktok/kitex_gen/social"
	"github.com/hcdoit/tiktok/pkg/errno"
	"github.com/hcdoit/tiktok/pkg/jwt"
)

// SocialServiceImpl implements the last service interface defined in the IDL.
type SocialServiceImpl struct{}

// RelationAction implements the SocialServiceImpl interface.
func (s *SocialServiceImpl) RelationAction(ctx context.Context, req *social.RelationActionRequest) (resp *social.RelationActionResponse, err error) {
	// TODO: Your code here...
	resp = new(social.RelationActionResponse)

	claim, err := jwt.ParseToken(req.Token)
	if err != nil {
		resp.StatusCode, resp.StatusMsg = utils.BuildStatus(err)
		return resp, nil
	}
	if err = req.IsValid(); err != nil {
		resp.StatusCode, resp.StatusMsg = utils.BuildStatus(err)
		return resp, nil
	}

	err = service.NewRelationService(ctx).RelationAction(req, claim.Id)
	if err != nil {
		resp.StatusCode, resp.StatusMsg = utils.BuildStatus(err)
		return resp, nil
	}

	resp.StatusCode, resp.StatusMsg = utils.BuildStatus(errno.Success)
	return resp, nil
}

// GetFollowList implements the SocialServiceImpl interface.
func (s *SocialServiceImpl) GetFollowList(ctx context.Context, req *social.RelationListRequest) (resp *social.RelationListResponse, err error) {
	// TODO: Your code here...
	resp = new(social.RelationListResponse)

	claim, err := jwt.ParseToken(req.Token)
	if err != nil {
		resp.StatusCode, resp.StatusMsg = utils.BuildStatus(err)
		return resp, nil
	}
	if err = req.IsValid(); err != nil {
		resp.StatusCode, resp.StatusMsg = utils.BuildStatus(err)
		return resp, nil
	}
	users, err := service.NewRelationService(ctx).RelationList(req, claim.Id, service.Follow)
	if err != nil {
		resp.StatusCode, resp.StatusMsg = utils.BuildStatus(err)
		return resp, nil
	}

	resp.StatusCode, resp.StatusMsg = utils.BuildStatus(errno.Success)
	resp.UserList = users
	return resp, nil
}

// GetFollowerList implements the SocialServiceImpl interface.
func (s *SocialServiceImpl) GetFollowerList(ctx context.Context, req *social.RelationListRequest) (resp *social.RelationListResponse, err error) {
	// TODO: Your code here...
	resp = new(social.RelationListResponse)

	claim, err := jwt.ParseToken(req.Token)
	if err != nil {
		resp.StatusCode, resp.StatusMsg = utils.BuildStatus(err)
		return resp, nil
	}
	if err = req.IsValid(); err != nil {
		resp.StatusCode, resp.StatusMsg = utils.BuildStatus(err)
		return resp, nil
	}
	users, err := service.NewRelationService(ctx).RelationList(req, claim.Id, service.Follower)
	if err != nil {
		resp.StatusCode, resp.StatusMsg = utils.BuildStatus(err)
		return resp, nil
	}

	resp.StatusCode, resp.StatusMsg = utils.BuildStatus(errno.Success)
	resp.UserList = users
	return resp, nil
}

// GetFriendList implements the SocialServiceImpl interface.
func (s *SocialServiceImpl) GetFriendList(ctx context.Context, req *social.RelationListRequest) (resp *social.RelationListResponse, err error) {
	// TODO: Your code here...
	resp = new(social.RelationListResponse)

	claim, err := jwt.ParseToken(req.Token)
	if err != nil {
		resp.StatusCode, resp.StatusMsg = utils.BuildStatus(err)
		return resp, nil
	}
	if err = req.IsValid(); err != nil {
		resp.StatusCode, resp.StatusMsg = utils.BuildStatus(err)
		return resp, nil
	}
	users, err := service.NewRelationService(ctx).RelationList(req, claim.Id, service.Follow)
	if err != nil {
		resp.StatusCode, resp.StatusMsg = utils.BuildStatus(err)
		return resp, nil
	}

	resp.StatusCode, resp.StatusMsg = utils.BuildStatus(errno.Success)
	resp.UserList = users
	return resp, nil
}

// GetRelationInfo implements the SocialServiceImpl interface.
func (s *SocialServiceImpl) GetRelationInfo(ctx context.Context, req *social.RelationInfoRequest) (resp *social.RelationInfoResponse, err error) {
	// TODO: Your code here...
	resp = new(social.RelationInfoResponse)
	if err = req.IsValid(); err != nil {
		resp.StatusCode, resp.StatusMsg = utils.BuildStatus(err)
		return resp, nil
	}
	followCount, followerCount, isFollow, err := service.NewRelationService(ctx).GetRelationInfo(req)
	if err != nil {
		resp.StatusCode, resp.StatusMsg = utils.BuildStatus(err)
		return resp, nil
	}

	resp.StatusCode, resp.StatusMsg = utils.BuildStatus(errno.Success)
	resp.FollowCount, resp.FollowerCount, resp.IsFollow = followCount, followerCount, isFollow
	return resp, nil
}
