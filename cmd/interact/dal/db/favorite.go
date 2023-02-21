package db

import (
	"context"
	"github.com/hcdoit/tiktok/pkg/consts"
	"gorm.io/gorm"
)

type Favorite struct {
	gorm.Model
	UserID  int64
	VideoID int64
}

func (f Favorite) TableName() string {
	return consts.FavoriteTableName
}

func CreateFavorite(ctx context.Context, favorite []*Favorite) error {
	return DB.WithContext(ctx).Create(favorite).Error
}

func CancelFavorite(ctx context.Context, userID int64, videoID int64) error {
	return DB.WithContext(ctx).Where("user_id = ? and video_id =?", userID, videoID).Unscoped().Delete(&Comment{}).Error
}

func QueryFavoriteByUserIDAndVideoID(ctx context.Context, userID int64, videoID int64) ([]*Favorite, error) {
	res := make([]*Favorite, 0)
	if err := DB.WithContext(ctx).Where("user_id = ? and video_id =?", userID, videoID).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

func QueryFavoriteByUserID(ctx context.Context, userID int64) ([]*Favorite, error) {
	res := make([]*Favorite, 0)
	if err := DB.WithContext(ctx).Where("user_id = ?", userID).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}
func QueryFavoriteByVideoID(ctx context.Context, videoID int64) ([]*Favorite, error) {
	res := make([]*Favorite, 0)
	if err := DB.WithContext(ctx).Where("video_id = ?", videoID).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}
