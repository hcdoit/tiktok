package service

import (
	"context"
	"github.com/hcdoit/tiktok/cmd/interact/dal/db"
	"github.com/hcdoit/tiktok/cmd/interact/dal/rdb"
	"github.com/hcdoit/tiktok/kitex_gen/interact"
)

type GetVideoInteractService struct {
	ctx context.Context
}

func NewGetVideoInteractService(ctx context.Context) *GetVideoInteractService {
	return &GetVideoInteractService{
		ctx: ctx,
	}
}

// GetVideoInteract 获取视频的点赞数、评论数
func (s *GetVideoInteractService) GetVideoInteract(req *interact.VideoInteractRequest) (favoriteCount int, commentCount int, isFavorite bool, err error) {
	favoriteCount, err = rdb.GetFavoriteCountByVideoID(s.ctx, req.VideoId)
	if err != nil {
		return 0, 0, false, err
	}
	commentCount, err = rdb.GetCommentCountByVideoID(s.ctx, req.VideoId)
	if err != nil {
		return 0, 0, false, err
	}
	favorite, err := db.QueryFavoriteByUserIDAndVideoID(s.ctx, req.UserId, req.VideoId)
	if err != nil {
		return 0, 0, false, err
	}
	if len(favorite) == 0 {
		return favoriteCount, commentCount, false, nil
	}
	return favoriteCount, commentCount, true, nil

}
