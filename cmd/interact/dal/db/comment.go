package db

import (
	"context"
	"github.com/hcdoit/tiktok/pkg/consts"
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	VideoID     int64
	UserID      int64
	CommentText string
}

func (f Comment) TableName() string {
	return consts.CommentTableName
}

func CreateComment(ctx context.Context, comment []*Comment) error {
	return DB.WithContext(ctx).Create(comment).Error
}

func CancelCommentByCommentID(ctx context.Context, commentID int64) error {
	return DB.WithContext(ctx).Where("id = ?", commentID).Delete(&Comment{}).Error
}

func QueryCommentByVideoID(ctx context.Context, videoID int64) ([]*Comment, error) {
	res := make([]*Comment, 0)
	if err := DB.WithContext(ctx).Where("video_id = ?", videoID).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}
