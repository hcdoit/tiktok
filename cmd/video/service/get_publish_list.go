package service

import (
	"context"
	"github.com/hcdoit/tiktok/cmd/video/dal/db"
	"github.com/hcdoit/tiktok/cmd/video/utils"
	"github.com/hcdoit/tiktok/kitex_gen/video"
)

type GetPublishListService struct {
	ctx context.Context
}

func NewGetPublishListService(ctx context.Context) *GetPublishListService {
	return &GetPublishListService{ctx: ctx}
}

func (s *GetPublishListService) GetPublishList(myID int64, ID int64) ([]*video.Video, error) {
	modelVideos, err := db.QueryVideoByAuthorID(s.ctx, ID)
	if err != nil {
		return nil, err
	}
	if len(modelVideos) == 0 {
		return nil, nil
	}
	return utils.BuildVideos(modelVideos, s.ctx, myID), nil
}
