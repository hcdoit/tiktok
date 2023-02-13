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

// GetMessageChat implements the SocialServiceImpl interface.
func (s *SocialServiceImpl) GetMessageChat(ctx context.Context, req *social.MessageChatRequest) (resp *social.MessageChatResponse, err error) {
	resp = new(social.MessageChatResponse)

	// 校验参数
	if err = req.IsValid(); err != nil {
		resp.StatusCode, resp.StatusMsg = utils.BuildStatus(err)
		return resp, nil
	}

	// 解析Token
	claim, err := jwt.ParseToken(req.Token)
	if err != nil {
		resp.StatusCode, resp.StatusMsg = utils.BuildStatus(err)
		return resp, nil
	}

	// 调用service层
	messages, err := service.NewMessageService(ctx).GetMessageChat(req, claim.Id)
	if err != nil {
		resp.StatusCode, resp.StatusMsg = utils.BuildStatus(err)
		return resp, nil
	}

	// 包装正常响应
	resp.MessageList = messages
	resp.StatusCode, resp.StatusMsg = utils.BuildStatus(errno.Success)
	return resp, nil

}

// MessageAction implements the SocialServiceImpl interface.
func (s *SocialServiceImpl) MessageAction(ctx context.Context, req *social.MessageActionRequest) (resp *social.MessageActionResponse, err error) {
	resp = new(social.MessageActionResponse)

	// 校验参数
	if err = req.IsValid(); err != nil {
		resp.StatusCode, resp.StatusMsg = utils.BuildStatus(err)
		return resp, nil
	}

	// 解析Token
	claim, err := jwt.ParseToken(req.Token)
	if err != nil {
		resp.StatusCode, resp.StatusMsg = utils.BuildStatus(err)
		return resp, nil
	}

	// 调用service层
	err = service.NewMessageService(ctx).MessageAction(req, claim.Id)
	if err != nil {
		resp.StatusCode, resp.StatusMsg = utils.BuildStatus(err)
		return resp, nil
	}

	// 包装正常响应
	resp.StatusCode, resp.StatusMsg = utils.BuildStatus(errno.Success)
	return resp, nil
}

// RelationAction implements the SocialServiceImpl interface.
func (s *SocialServiceImpl) RelationAction(ctx context.Context, req *social.RelationActionRequest) (resp *social.RelationActionResponse, err error) {
	resp = new(social.RelationActionResponse)

	// 校验参数
	if err = req.IsValid(); err != nil {
		resp.StatusCode, resp.StatusMsg = utils.BuildStatus(err)
		return resp, nil
	}

	// 解析Token
	claim, err := jwt.ParseToken(req.Token)
	if err != nil {
		resp.StatusCode, resp.StatusMsg = utils.BuildStatus(err)
		return resp, nil
	}

	// 调用service层
	err = service.NewRelationService(ctx).RelationAction(req, claim.Id)
	if err != nil {
		resp.StatusCode, resp.StatusMsg = utils.BuildStatus(err)
		return resp, nil
	}

	// 包装正常响应
	resp.StatusCode, resp.StatusMsg = utils.BuildStatus(errno.Success)
	return resp, nil
}

// GetFollowList implements the SocialServiceImpl interface.
func (s *SocialServiceImpl) GetFollowList(ctx context.Context, req *social.RelationListRequest) (resp *social.RelationListResponse, err error) {
	resp = new(social.RelationListResponse)

	// 校验参数
	if err = req.IsValid(); err != nil {
		resp.StatusCode, resp.StatusMsg = utils.BuildStatus(err)
		return resp, nil
	}

	// 解析Token
	claim, err := jwt.ParseToken(req.Token)
	if err != nil {
		resp.StatusCode, resp.StatusMsg = utils.BuildStatus(err)
		return resp, nil
	}

	// 调用service层
	users, err := service.NewRelationService(ctx).RelationList(req, claim.Id, service.Follow)
	if err != nil {
		resp.StatusCode, resp.StatusMsg = utils.BuildStatus(err)
		return resp, nil
	}

	// 包装正常响应
	resp.StatusCode, resp.StatusMsg = utils.BuildStatus(errno.Success)
	resp.UserList = users
	return resp, nil
}

// GetFollowerList implements the SocialServiceImpl interface.
func (s *SocialServiceImpl) GetFollowerList(ctx context.Context, req *social.RelationListRequest) (resp *social.RelationListResponse, err error) {
	resp = new(social.RelationListResponse)

	// 校验参数
	if err = req.IsValid(); err != nil {
		resp.StatusCode, resp.StatusMsg = utils.BuildStatus(err)
		return resp, nil
	}

	// 解析Token
	claim, err := jwt.ParseToken(req.Token)
	if err != nil {
		resp.StatusCode, resp.StatusMsg = utils.BuildStatus(err)
		return resp, nil
	}

	// 调用service层
	users, err := service.NewRelationService(ctx).RelationList(req, claim.Id, service.Follower)
	if err != nil {
		resp.StatusCode, resp.StatusMsg = utils.BuildStatus(err)
		return resp, nil
	}

	// 包装正常响应
	resp.StatusCode, resp.StatusMsg = utils.BuildStatus(errno.Success)
	resp.UserList = users
	return resp, nil
}

// GetFriendList implements the SocialServiceImpl interface.
func (s *SocialServiceImpl) GetFriendList(ctx context.Context, req *social.RelationListRequest) (resp *social.RelationListResponse, err error) {
	resp = new(social.RelationListResponse)

	// 校验参数
	if err = req.IsValid(); err != nil {
		resp.StatusCode, resp.StatusMsg = utils.BuildStatus(err)
		return resp, nil
	}

	// 解析Token
	claim, err := jwt.ParseToken(req.Token)
	if err != nil {
		resp.StatusCode, resp.StatusMsg = utils.BuildStatus(err)
		return resp, nil
	}

	// 调用service层
	users, err := service.NewRelationService(ctx).RelationList(req, claim.Id, service.Follow)
	if err != nil {
		resp.StatusCode, resp.StatusMsg = utils.BuildStatus(err)
		return resp, nil
	}

	// 包装正常响应
	resp.StatusCode, resp.StatusMsg = utils.BuildStatus(errno.Success)
	resp.UserList = users
	return resp, nil
}

// GetRelationInfo implements the SocialServiceImpl interface.
func (s *SocialServiceImpl) GetRelationInfo(ctx context.Context, req *social.RelationInfoRequest) (resp *social.RelationInfoResponse, err error) {
	resp = new(social.RelationInfoResponse)

	// 校验参数
	if err = req.IsValid(); err != nil {
		resp.StatusCode, resp.StatusMsg = utils.BuildStatus(err)
		return resp, nil
	}

	// 调用service层
	followCount, followerCount, isFollow, err := service.NewRelationService(ctx).GetRelationInfo(req)
	if err != nil {
		resp.StatusCode, resp.StatusMsg = utils.BuildStatus(err)
		return resp, nil
	}

	// 包装正常响应
	resp.StatusCode, resp.StatusMsg = utils.BuildStatus(errno.Success)
	resp.FollowCount, resp.FollowerCount, resp.IsFollow = followCount, followerCount, isFollow
	return resp, nil
}
