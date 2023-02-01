package service

import (
	"context"
	"github.com/hcdoit/tiktok/cmd/video/dal/db"
	"github.com/hcdoit/tiktok/cmd/video/utils"
	"github.com/hcdoit/tiktok/kitex_gen/video"
)

type GetFeedService struct {
	ctx context.Context
}

func NewGetFeedService(ctx context.Context) *GetFeedService {
	return &GetFeedService{ctx: ctx}
}

func (s *GetFeedService) GetFeed(id int64, time int64) (videos []*video.Video, nextTime int64, err error) {
	modelVideos, err := db.QueryVideBeforeTime(s.ctx, time)
	if err != nil {
		return nil, 0, err
	}
	if len(modelVideos) == 0 {
		return nil, 0, nil
	}
	return utils.BuildVideos(modelVideos, s.ctx, id), modelVideos[0].PublishAt, nil
}
