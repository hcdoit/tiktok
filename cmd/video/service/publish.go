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

type PublishService struct {
	ctx context.Context
}

func NewPublishService(ctx context.Context) *PublishService {
	return &PublishService{ctx: ctx}
}

// GetPublishList 获取用户上传的所有视频
func (s *PublishService) GetPublishList(myID int64, ID int64) ([]*video.Video, error) {
	modelVideos, err := db.QueryVideoByAuthorID(s.ctx, ID)
	if err != nil {
		return nil, err
	}
	if len(modelVideos) == 0 {
		return nil, nil
	}
	return utils.BuildVideos(modelVideos, s.ctx, myID), nil
}

// PublishAction 上传视频
func (s *PublishService) PublishAction(req *video.PublishActionRequest, authorID int64) error {

	publishTime := time.Now().Unix()
	videoName := fmt.Sprintf("%d_%d.mp4", authorID, publishTime)
	// 保存视频
	localURL, url, err := minio.SaveVideo(s.ctx, videoName, req.Data)
	if err != nil {
		return err
	}
	// 获取封面，顺带检查保存视频的URL是否有效
	coverData, err := utils.GetCover(localURL)
	if err != nil {
		return err
	}
	coverName := fmt.Sprintf("%d_%d.jpg", authorID, publishTime)
	// 保存封面
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
