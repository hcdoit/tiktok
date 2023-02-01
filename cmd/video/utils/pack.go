package utils

import (
	"context"
	"errors"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/hcdoit/tiktok/cmd/video/dal/db"
	"github.com/hcdoit/tiktok/cmd/video/rpc"
	"github.com/hcdoit/tiktok/kitex_gen/user"
	"github.com/hcdoit/tiktok/kitex_gen/video"
	"github.com/hcdoit/tiktok/pkg/errno"
)

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

func BuildVideo(v *db.Video, ctx context.Context, myID int64) (*video.Video, error) {
	if v == nil {
		return nil, nil
	}
	resp, err := rpc.GetUser(ctx, &user.GetUserRequest{
		UserId: v.AuthorID,
		Token:  "",
	})
	if err != nil {
		return nil, err
	}
	return &video.Video{Id: int64(v.ID), Author: resp, PlayUrl: v.PlayURL, CoverUrl: v.CoverURL, Title: v.Title}, nil
}

func BuildVideos(vs []*db.Video, ctx context.Context, myID int64) []*video.Video {
	videos := make([]*video.Video, 0)
	for _, v := range vs {
		klog.Info(v.AuthorID)
		if temp, err := BuildVideo(v, ctx, myID); temp != nil && err == nil {
			videos = append(videos, temp)
		}
	}
	return videos
}
