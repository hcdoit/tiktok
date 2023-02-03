package service

import (
	"context"
	"github.com/hcdoit/tiktok/cmd/video/dal/db"
	"github.com/hcdoit/tiktok/cmd/video/utils"
	"github.com/hcdoit/tiktok/kitex_gen/video"
	"github.com/hcdoit/tiktok/pkg/errno"
)

type GetVideoService struct {
	ctx context.Context
}

func NewGetVideoService(ctx context.Context) *GetVideoService {
	return &GetVideoService{ctx: ctx}
}

func (s *GetVideoService) GetVideo(req *video.VideoRequest) (*video.Video, error) {

	videos, err := db.QueryVideoByID(s.ctx, req.VideoId)
	if err != nil {
		return nil, err
	}
	if len(videos) != 1 {
		return nil, errno.ResourceNotFound
	}
	buildVideo, err := utils.BuildVideo(videos[0], s.ctx, req.UserId)
	if err != nil {
		return nil, err
	}
	return buildVideo, nil
}
