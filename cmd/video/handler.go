package main

import (
	"context"
	"github.com/hcdoit/tiktok/cmd/video/service"
	"github.com/hcdoit/tiktok/cmd/video/utils"
	video "github.com/hcdoit/tiktok/kitex_gen/video"
	"github.com/hcdoit/tiktok/pkg/errno"
	"github.com/hcdoit/tiktok/pkg/jwt"
	"time"
)

// VideoServiceImpl implements the last service interface defined in the IDL.
type VideoServiceImpl struct{}

// GetVideo implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) GetVideo(ctx context.Context, req *video.VideoRequest) (resp *video.VideoResponse, err error) {
	resp = new(video.VideoResponse)

	// 校验参数
	if err = req.IsValid(); err != nil {
		resp.StatusCode, resp.StatusMsg = utils.BuildStatus(err)
		return resp, nil
	}

	// 调用service层
	getVideo, err := service.NewGetVideoService(ctx).GetVideo(req)
	if err != nil {
		resp.StatusCode, resp.StatusMsg = utils.BuildStatus(err)
		return resp, nil
	}

	// 包装正常响应
	resp.StatusCode, resp.StatusMsg = utils.BuildStatus(errno.Success)
	resp.Video = getVideo

	return resp, nil
}

// GetFeed implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) GetFeed(ctx context.Context, req *video.FeedRequest) (resp *video.FeedResponse, err error) {
	resp = new(video.FeedResponse)

	// 校验参数
	if err = req.IsValid(); err != nil {
		resp.StatusCode, resp.StatusMsg = utils.BuildStatus(err)
		return resp, nil
	}

	// 解析Token获取id，若空默认为0
	myID := int64(0)
	if len(req.Token) != 0 {
		claim, err := jwt.ParseToken(req.Token)
		if err != nil {
			resp.StatusCode, resp.StatusMsg = utils.BuildStatus(err)
			return resp, nil
		}
		myID = claim.Id
	}

	// 读取请求时间，若空默认为当前时间
	latestTime := time.Now().Unix()
	if req.LatestTime != 0 {
		latestTime = req.GetLatestTime()
	}

	// 调用service层
	videos, nextTime, err := service.NewGetFeedService(ctx).GetFeed(myID, latestTime)
	if err != nil {
		resp.StatusCode, resp.StatusMsg = utils.BuildStatus(err)
		return resp, nil
	}

	// 包装正常响应
	resp.StatusCode, resp.StatusMsg = utils.BuildStatus(errno.Success)
	resp.VideoList = videos
	resp.NextTime = nextTime

	return resp, nil
}

// PublishAction implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) PublishAction(ctx context.Context, req *video.PublishActionRequest) (resp *video.PublishActionResponse, err error) {
	resp = new(video.PublishActionResponse)

	// 解析Token
	claim, err := jwt.ParseToken(req.Token)
	if err != nil {
		resp.StatusCode, resp.StatusMsg = utils.BuildStatus(err)
		return resp, nil
	}

	// 校验参数
	if err = req.IsValid(); err != nil {
		resp.StatusCode, resp.StatusMsg = utils.BuildStatus(err)
		return resp, nil
	}

	// 调用service层
	err = service.NewPublishService(ctx).PublishAction(req, claim.Id)
	if err != nil {
		resp.StatusCode, resp.StatusMsg = utils.BuildStatus(err)
		return resp, nil
	}

	// 包装正常响应
	resp.StatusCode, resp.StatusMsg = utils.BuildStatus(errno.Success)
	return resp, nil
}

// GetPublishList implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) GetPublishList(ctx context.Context, req *video.PublishListRequest) (resp *video.PublishListResponse, err error) {
	resp = new(video.PublishListResponse)

	// 校验参数
	if err = req.IsValid(); err != nil {
		resp.StatusCode, resp.StatusMsg = utils.BuildStatus(err)
		return resp, nil
	}

	// 解析Token获取id，若空默认为0
	myID := int64(0)
	if len(req.Token) != 0 {
		claim, err := jwt.ParseToken(req.Token)
		if err != nil {
			resp.StatusCode, resp.StatusMsg = utils.BuildStatus(err)
			return resp, nil
		}
		myID = claim.Id
	}

	// 调用service层
	videos, err := service.NewPublishService(ctx).GetPublishList(myID, req.UserId)
	if err != nil {
		resp.StatusCode, resp.StatusMsg = utils.BuildStatus(err)
		return resp, nil
	}

	// 包装正常响应
	resp.VideoList = videos
	resp.StatusCode, resp.StatusMsg = utils.BuildStatus(errno.Success)
	return resp, nil
}
