package db

import (
	"context"
	"github.com/hcdoit/tiktok/pkg/consts"

	"gorm.io/gorm"
)

type Video struct {
	gorm.Model
	AuthorID  int64
	PublishAt int64
	PlayURL   string
	CoverURL  string
	Title     string
}

func (u *Video) TableName() string {
	return consts.VideoTableName
}

func CreateVideo(ctx context.Context, videos []*Video) error {
	return DB.WithContext(ctx).Create(videos).Error
}

func QueryVideoByID(ctx context.Context, videoID int64) ([]*Video, error) {
	res := make([]*Video, 0)
	if err := DB.WithContext(ctx).Where("id = ?", videoID).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

func QueryVideoByAuthorID(ctx context.Context, authorID int64) ([]*Video, error) {
	res := make([]*Video, 0)
	if err := DB.WithContext(ctx).Where("author_id = ?", authorID).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

func QueryVideBeforeTime(ctx context.Context, time int64) ([]*Video, error) {
	res := make([]*Video, 0)
	if err := DB.WithContext(ctx).Where("publish_at < ?", time).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}
