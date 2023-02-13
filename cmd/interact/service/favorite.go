package service

import (
	"context"
	"github.com/hcdoit/tiktok/cmd/interact/dal/db"
	"github.com/hcdoit/tiktok/cmd/interact/dal/rdb"
	"github.com/hcdoit/tiktok/cmd/interact/rpc"
	"github.com/hcdoit/tiktok/kitex_gen/interact"
	"github.com/hcdoit/tiktok/kitex_gen/video"
	"github.com/hcdoit/tiktok/pkg/errno"
)

type FavoriteService struct {
	ctx context.Context
}

func NewFavoriteService(ctx context.Context) *FavoriteService {
	return &FavoriteService{
		ctx: ctx,
	}
}

// FavoriteAction 点赞或取消赞
func (s FavoriteService) FavoriteAction(req *interact.FavoriteActionRequest, userID int64) (err error) {
	if req.ActionType == 1 {
		err = db.CreateFavorite(s.ctx, []*db.Favorite{
			{
				UserID:  userID,
				VideoID: req.VideoId,
			},
		})
	}
	if req.ActionType == 2 {
		err = db.CancelFavorite(s.ctx, userID, req.VideoId)
	}
	if err != nil {
		return err
	}
	if req.ActionType == 1 {
		err = rdb.AddFavoriteCountByVideoID(s.ctx, req.VideoId)
	}
	if req.ActionType == 2 {
		err = rdb.MinusFavoriteCountByVideoID(s.ctx, req.VideoId)
	}

	if err != nil {
		return err
	}

	return nil
}

// GetFavoriteList 获取用户所有点赞视频
func (s FavoriteService) GetFavoriteList(userID int64, myID int64) (videos []*video.Video, err error) {
	favorites, err := db.QueryFavoriteByUserID(s.ctx, userID)
	if err != nil {
		return nil, err
	}
	videos = make([]*video.Video, 0)
	for _, v := range favorites {
		if resp, err := rpc.GetVideo(s.ctx, &video.VideoRequest{
			UserId:  myID,
			VideoId: v.VideoID,
		}); resp != nil && resp.StatusCode == errno.Success.ErrCode && err == nil {
			videos = append(videos, resp.Video)
		}
		if err != nil {
			return nil, err
		}
	}
	return videos, nil
}
