package service

import (
	"context"
	"github.com/hcdoit/tiktok/cmd/interact/dal/db"
	"github.com/hcdoit/tiktok/cmd/interact/dal/rdb"
	"github.com/hcdoit/tiktok/cmd/interact/utils"
	"github.com/hcdoit/tiktok/kitex_gen/interact"
)

type CommentService struct {
	ctx context.Context
}

func NewCommentService(ctx context.Context) *CommentService {
	return &CommentService{
		ctx: ctx,
	}
}

// CommentAction 添加或删除评论
func (s CommentService) CommentAction(req *interact.CommentActionRequest, userID int64) (err error) {

	// 添加评论
	if req.ActionType == 1 {
		err = db.CreateComment(s.ctx, []*db.Comment{
			{
				UserID:      userID,
				VideoID:     req.VideoId,
				CommentText: req.CommentText,
			},
		})
	}
	// 删除评论
	if req.ActionType == 2 {
		err = db.CancelCommentByCommentID(s.ctx, req.CommentId)
	}
	if err != nil {
		return err
	}

	// 更新缓存
	if req.ActionType == 1 {
		err = rdb.AddCommentCountByVideoID(s.ctx, req.VideoId)
	}
	if err != nil {
		return err
	}
	if req.ActionType == 2 {
		err = rdb.MinusCommentCountByVideoID(s.ctx, req.VideoId)
	}

	if err != nil {
		return err
	}

	return nil
}

// GetCommentList 获取视频下的所有评论
func (s CommentService) GetCommentList(videoID int64, myID int64) (comments []*interact.Comment, err error) {
	modelComments, err := db.QueryCommentByVideoID(s.ctx, videoID)
	if err != nil {
		return nil, err
	}

	return utils.BuildComments(modelComments, s.ctx, myID), nil
}
