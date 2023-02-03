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

func (s *VideoServiceImpl) GetVideo(ctx context.Context, req *video.VideoRequest) (resp *video.VideoResponse, err error) {
	// TODO: Your code here...
	resp = new(video.VideoResponse)
	if err = req.IsValid(); err != nil {
		resp.StatusCode, resp.StatusMsg = utils.BuildStatus(err)
		return resp, nil
	}
	getVideo, err := service.NewGetVideoService(ctx).GetVideo(req)
	if err != nil {
		resp.StatusCode, resp.StatusMsg = utils.BuildStatus(err)
		return resp, nil
	}
	resp.StatusCode, resp.StatusMsg = utils.BuildStatus(errno.Success)
	resp.Video = getVideo

	return resp, nil
}

// GetFeed implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) GetFeed(ctx context.Context, req *video.FeedRequest) (resp *video.FeedResponse, err error) {
	// TODO: Your code here...
	resp = new(video.FeedResponse)
	if err = req.IsValid(); err != nil {
		resp.StatusCode, resp.StatusMsg = utils.BuildStatus(err)
		return resp, nil
	}
	myID := int64(0)
	if len(req.Token) != 0 {
		claim, err := jwt.ParseToken(req.Token)
		if err != nil {
			resp.StatusCode, resp.StatusMsg = utils.BuildStatus(err)
			return resp, nil
		}
		myID = claim.Id
	}
	latestTime := time.Now().Unix()
	if req.LatestTime != 0 {
		latestTime = req.GetLatestTime()
	}
	videos, nextTime, err := service.NewGetFeedService(ctx).GetFeed(myID, latestTime)
	if err != nil {
		resp.StatusCode, resp.StatusMsg = utils.BuildStatus(err)
		return resp, nil
	}
	resp.StatusCode, resp.StatusMsg = utils.BuildStatus(errno.Success)
	resp.VideoList = videos
	resp.NextTime = nextTime

	return resp, nil
}

// PublishAction implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) PublishAction(ctx context.Context, req *video.PublishActionRequest) (resp *video.PublishActionResponse, err error) {
	// TODO: Your code here...
	resp = new(video.PublishActionResponse)
	claim, err := jwt.ParseToken(req.Token)
	if err != nil {
		resp.StatusCode, resp.StatusMsg = utils.BuildStatus(err)
		return resp, nil
	}
	//判断参数长度
	if err = req.IsValid(); err != nil {
		resp.StatusCode, resp.StatusMsg = utils.BuildStatus(err)
		return resp, nil
	}
	err = service.NewPublishService(ctx).PublishAction(req, claim.Id)
	if err != nil {
		resp.StatusCode, resp.StatusMsg = utils.BuildStatus(err)
		return resp, nil
	}
	resp.StatusCode, resp.StatusMsg = utils.BuildStatus(errno.Success)
	return resp, nil
}

// GetPublishList implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) GetPublishList(ctx context.Context, req *video.PublishListRequest) (resp *video.PublishListResponse, err error) {
	// TODO: Your code here...
	resp = new(video.PublishListResponse)
	if err = req.IsValid(); err != nil {
		resp.StatusCode, resp.StatusMsg = utils.BuildStatus(err)
		return resp, nil
	}
	myID := int64(0)
	if len(req.Token) != 0 {
		claim, err := jwt.ParseToken(req.Token)
		if err != nil {
			resp.StatusCode, resp.StatusMsg = utils.BuildStatus(err)
			return resp, nil
		}
		myID = claim.Id
	}

	videos, err := service.NewPublishService(ctx).GetPublishList(myID, req.UserId)
	resp.VideoList = videos
	resp.StatusCode, resp.StatusMsg = utils.BuildStatus(errno.Success)
	return resp, nil
}
