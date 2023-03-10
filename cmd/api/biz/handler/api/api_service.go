// Code generated by hertz generator.

package api

import (
	"bytes"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/hcdoit/tiktok/cmd/api/biz/model/api"
	"github.com/hcdoit/tiktok/cmd/api/biz/rpc"
	"github.com/hcdoit/tiktok/kitex_gen/interact"
	"github.com/hcdoit/tiktok/kitex_gen/social"
	"github.com/hcdoit/tiktok/kitex_gen/user"
	"github.com/hcdoit/tiktok/kitex_gen/video"
	"github.com/hcdoit/tiktok/pkg/errno"
	"io"
)

// Register .
// @router /douyin/user/register [POST]
func Register(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.UserRegisterRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.JSON(consts.StatusOK, errno.ParamErr)
		return
	}

	resp, err := rpc.Register(ctx, &user.UserRegisterRequest{
		Username: c.Query("username"),
		Password: c.Query("password"),
	})

	if err != nil {
		c.JSON(consts.StatusOK, errno.ServiceErr)
		return
	}

	c.JSON(consts.StatusOK, resp)
}

// Login .
// @router /douyin/user/login [POST]
func Login(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.UserLoginRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.JSON(consts.StatusOK, errno.ParamErr)
		return
	}
	resp, err := rpc.Login(ctx, &user.UserLoginRequest{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		c.JSON(consts.StatusOK, errno.ServiceErr)
		return
	}

	c.JSON(consts.StatusOK, resp)
}

// GetUser .
// @router /douyin/user [POST]
func GetUser(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.GetUserRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.JSON(consts.StatusOK, errno.ParamErr)
		return
	}
	resp, err := rpc.GetUser(ctx, &user.GetUserRequest{
		UserId: req.UserID,
		Token:  req.Token,
	})
	if err != nil {
		c.JSON(consts.StatusOK, errno.ServiceErr)
		return
	}

	c.JSON(consts.StatusOK, resp)
}

// GetFeed .
// @router /douyin/feed [GET]
func GetFeed(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.FeedRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.JSON(consts.StatusOK, errno.ParamErr)
		return
	}
	resp, err := rpc.GetFeed(ctx, &video.FeedRequest{
		LatestTime: req.LatestTime,
		Token:      req.Token,
	})
	if err != nil {
		c.JSON(consts.StatusOK, errno.ServiceErr)
		return
	}

	c.JSON(consts.StatusOK, resp)
}

// PublishAction .
// @router /douyin/publish/action [POST]
func PublishAction(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.PublishActionRequest
	req.Token = c.PostForm("token")
	req.Title = c.PostForm("title")

	fileHeader, err := c.Request.FormFile("data")
	if err != nil {
		hlog.Debug("canot get video data")
		c.JSON(consts.StatusOK, errno.ParamErr)
		return
	}
	file, err := fileHeader.Open()
	if err != nil {
		hlog.Error("cannot open data")
		c.JSON(consts.StatusOK, errno.ServiceErr)
		return
	}
	defer file.Close()
	buf := bytes.NewBuffer(nil)
	_, err = io.Copy(buf, file)
	if err != nil {
		hlog.Error("cannot copy data")
		c.JSON(consts.StatusOK, errno.ServiceErr)
		return
	}
	req.Data = buf.Bytes()

	resp, err := rpc.PublishAction(ctx, &video.PublishActionRequest{
		Token: req.Token,
		Data:  req.Data,
		Title: req.Title,
	})
	if err != nil {
		c.JSON(consts.StatusOK, errno.ServiceErr)
		return
	}

	c.JSON(consts.StatusOK, resp)
}

// GetPublishList .
// @router /douyin/publish/list [GET]
func GetPublishList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.PublishListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.JSON(consts.StatusOK, errno.ParamErr)
		return
	}
	resp, err := rpc.GetPublishList(ctx, &video.PublishListRequest{
		UserId: req.UserID,
		Token:  req.Token,
	})
	if err != nil {
		c.JSON(consts.StatusOK, errno.ServiceErr)
		return
	}

	c.JSON(consts.StatusOK, resp)
}

// FavoriteAction .
// @router /douyin/favorite/action/ [POST]
func FavoriteAction(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.FavoriteActionRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.JSON(consts.StatusOK, errno.ParamErr)
		return
	}
	resp, err := rpc.FavoriteAction(ctx, &interact.FavoriteActionRequest{
		Token:      req.Token,
		VideoId:    req.VideoID,
		ActionType: req.ActionType,
	})
	if err != nil {
		c.JSON(consts.StatusOK, errno.ServiceErr)
		return
	}

	c.JSON(consts.StatusOK, resp)
}

// GetFavoriteList .
// @router /douyin/favorite/list/ [GET]
func GetFavoriteList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.FavoriteListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.JSON(consts.StatusOK, errno.ParamErr)
		return
	}
	resp, err := rpc.GetFavoriteList(ctx, &interact.FavoriteListRequest{
		Token:  req.Token,
		UserId: req.UserID,
	})
	if err != nil {
		c.JSON(consts.StatusOK, errno.ServiceErr)
		return
	}

	c.JSON(consts.StatusOK, resp)
}

// CommentAction .
// @router /douyin/comment/action/ [POST]
func CommentAction(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.CommentActionRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.JSON(consts.StatusOK, errno.ParamErr)
		return
	}
	resp, err := rpc.CommentAction(ctx, &interact.CommentActionRequest{
		Token:       req.Token,
		VideoId:     req.VideoID,
		ActionType:  req.ActionType,
		CommentText: req.CommentText,
		CommentId:   req.CommentID,
	})
	if err != nil {
		c.JSON(consts.StatusOK, errno.ServiceErr)
		return
	}

	c.JSON(consts.StatusOK, resp)
}

// GetCommentList .
// @router /douyin/comment/list/ [GET]
func GetCommentList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.CommentListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.JSON(consts.StatusOK, errno.ParamErr)
		return
	}
	resp, err := rpc.GetCommentList(ctx, &interact.CommentListRequest{
		Token:   req.Token,
		VideoId: req.VideoID,
	})
	if err != nil {
		c.JSON(consts.StatusOK, errno.ServiceErr)
		return
	}

	c.JSON(consts.StatusOK, resp)
}

// RelationAction .
// @router /douyin/relation/action/ [POST]
func RelationAction(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.RelationActionRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.JSON(consts.StatusOK, errno.ParamErr)
		return
	}

	resp, err := rpc.RelationAction(ctx, &social.RelationActionRequest{
		Token:      req.Token,
		ToUserId:   req.ToUserID,
		ActionType: req.ActionType,
	})

	if err != nil {
		c.JSON(consts.StatusOK, errno.ServiceErr)
		return
	}

	c.JSON(consts.StatusOK, resp)
}

// GetFollowList .
// @router /douyin/relation/follow/list/ [GET]
func GetFollowList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.RelationListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.JSON(consts.StatusOK, errno.ParamErr)
		return
	}

	resp, err := rpc.GetFollowList(ctx, &social.RelationListRequest{
		UserId: req.UserID,
		Token:  req.Token,
	})

	if err != nil {
		c.JSON(consts.StatusOK, errno.ServiceErr)
		return
	}

	c.JSON(consts.StatusOK, resp)
}

// GetFollowerList .
// @router /douyin/relation/follower/list/ [GET]
func GetFollowerList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.RelationListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.JSON(consts.StatusOK, errno.ParamErr)
		return
	}

	resp, err := rpc.GetFollowerList(ctx, &social.RelationListRequest{
		UserId: req.UserID,
		Token:  req.Token,
	})

	if err != nil {
		c.JSON(consts.StatusOK, errno.ServiceErr)
		return
	}

	c.JSON(consts.StatusOK, resp)
}

// GetFriendList .
// @router /douyin/relation/friend/list/ [GET]
func GetFriendList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.RelationListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.JSON(consts.StatusOK, errno.ParamErr)
		return
	}
	resp, err := rpc.GetFriendList(ctx, &social.RelationListRequest{
		UserId: req.UserID,
		Token:  req.Token,
	})

	if err != nil {
		c.JSON(consts.StatusOK, errno.ServiceErr)
		return
	}

	c.JSON(consts.StatusOK, resp)
}

// GetMessageChat .
// @router /douyin/message/chat/ [GET]
func GetMessageChat(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.MessageChatRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.JSON(consts.StatusOK, errno.ParamErr)
		return
	}

	resp, err := rpc.GetMessageChat(ctx, &social.MessageChatRequest{
		Token:    req.Token,
		ToUserId: req.ToUserID,
	})

	if err != nil {
		c.JSON(consts.StatusOK, errno.ServiceErr)
		return
	}

	c.JSON(consts.StatusOK, resp)
}

// MessageAction .
// @router /douyin/message/action/ [POST]
func MessageAction(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.MessageActionRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.JSON(consts.StatusOK, errno.ParamErr)
		return
	}

	resp, err := rpc.MessageAction(ctx, &social.MessageActionRequest{
		Token:      req.Token,
		ToUserId:   req.ToUserID,
		ActionType: req.ActionType,
		Content:    req.Content,
	})

	if err != nil {
		c.JSON(consts.StatusOK, errno.ServiceErr)
		return
	}

	c.JSON(consts.StatusOK, resp)
}
