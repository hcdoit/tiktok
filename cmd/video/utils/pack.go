package utils

import (
	"context"
	"errors"
	"github.com/hcdoit/tiktok/cmd/video/dal/db"
	"github.com/hcdoit/tiktok/cmd/video/rpc"
	"github.com/hcdoit/tiktok/kitex_gen/interact"
	"github.com/hcdoit/tiktok/kitex_gen/user"
	"github.com/hcdoit/tiktok/kitex_gen/video"
	"github.com/hcdoit/tiktok/pkg/errno"
	"github.com/hcdoit/tiktok/pkg/jwt"
)

// BuildStatus 将error转换为StatusCode和StatusMsg
func BuildStatus(err error) (int32, string) {
	if err == nil {
		return errno.Success.ErrCode, errno.Success.ErrMsg
	}
	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return e.ErrCode, e.ErrMsg
	}
	s := errno.ServiceErr.WithMessage(err.Error())
	return s.ErrCode, s.ErrMsg
}

// BuildVideo 打包单个视频
func BuildVideo(v *db.Video, ctx context.Context, myID int64) (*video.Video, error) {
	if v == nil {
		return nil, nil
	}
	token, err := jwt.CreateToken(jwt.CustomClaims{
		Id: int64(myID),
	})
	if err != nil {
		return nil, err
	}
	author, err := rpc.GetUser(ctx, &user.GetUserRequest{
		UserId: v.AuthorID,
		Token:  token,
	})
	if err != nil {
		return nil, err
	}
	info, err := rpc.GetInteract(ctx, &interact.VideoInteractRequest{
		UserId:  myID,
		VideoId: int64(v.ID),
	})
	if err != nil {
		return nil, err
	}
	return &video.Video{
		Id:            int64(v.ID),
		Author:        author,
		PlayUrl:       v.PlayURL,
		CoverUrl:      v.CoverURL,
		Title:         v.Title,
		CommentCount:  info.CommentCount,
		FavoriteCount: info.FavoriteCount,
		IsFavorite:    info.IsFavorite,
	}, nil
}

// BuildVideos 打包视频列表
func BuildVideos(vs []*db.Video, ctx context.Context, myID int64) []*video.Video {
	videos := make([]*video.Video, 0)
	for _, v := range vs {
		if temp, err := BuildVideo(v, ctx, myID); temp != nil && err == nil {
			videos = append(videos, temp)
		}
	}
	return videos
}
