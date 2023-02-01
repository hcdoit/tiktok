package service

import (
	"context"
	"fmt"
	"github.com/hcdoit/tiktok/cmd/video/dal/db"
	"github.com/hcdoit/tiktok/cmd/video/dal/minio"
	"github.com/hcdoit/tiktok/cmd/video/utils"
	"github.com/hcdoit/tiktok/kitex_gen/video"
	"time"
)

type PublishActionService struct {
	ctx context.Context
}

// NewPublishActionService new PublishActionService
func NewPublishActionService(ctx context.Context) *PublishActionService {
	return &PublishActionService{ctx: ctx}
}

func (s *PublishActionService) PublishAction(req *video.PublishActionRequest, authorID int64) error {

	publishTime := time.Now().Unix()
	videoName := fmt.Sprintf("%d_%d.mp4", authorID, publishTime)
	localURL, url, err := minio.SaveVideo(s.ctx, videoName, req.Data)
	if err != nil {
		return err
	}
	coverData, err := utils.GetCover(localURL)
	if err != nil {
		return err
	}
	coverName := fmt.Sprintf("%d_%d.jpg", authorID, publishTime)
	covorURL, err := minio.SaveCover(s.ctx, coverName, coverData)
	if err != nil {
		return err
	}
	return db.CreateVideo(s.ctx, []*db.Video{{
		AuthorID:  authorID,
		PublishAt: publishTime,
		PlayURL:   url,
		CoverURL:  covorURL,
		Title:     req.Title,
	}})
}
