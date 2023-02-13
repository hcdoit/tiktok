package main

import (
	"context"
	"github.com/hcdoit/tiktok/cmd/interact/service"
	"github.com/hcdoit/tiktok/cmd/interact/utils"
	interact "github.com/hcdoit/tiktok/kitex_gen/interact"
	"github.com/hcdoit/tiktok/pkg/errno"
	"github.com/hcdoit/tiktok/pkg/jwt"
)

// InteractServiceImpl implements the last service interface defined in the IDL.
type InteractServiceImpl struct{}

// FavoriteAction implements the InteractServiceImpl interface.
func (s *InteractServiceImpl) FavoriteAction(ctx context.Context, req *interact.FavoriteActionRequest) (resp *interact.FavoriteActionResponse, err error) {
	resp = new(interact.FavoriteActionResponse)

	// 校验参数
	if err = req.IsValid(); err != nil {
		resp.StatusCode, resp.StatusMsg = utils.BuildStatus(err)
		return resp, nil
	}

	// 解析Token
	claim, err := jwt.ParseToken(req.Token)
	if err != nil {
		resp.StatusCode, resp.StatusMsg = utils.BuildStatus(err)
		return resp, nil
	}

	// 调用service层
	err = service.NewFavoriteService(ctx).FavoriteAction(req, claim.Id)
	if err != nil {
		resp.StatusCode, resp.StatusMsg = utils.BuildStatus(err)
		return resp, nil
	}

	// 包装正常响应
	resp.StatusCode, resp.StatusMsg = utils.BuildStatus(errno.Success)
	return resp, nil
}

// GetFavoriteList implements the InteractServiceImpl interface.
func (s *InteractServiceImpl) GetFavoriteList(ctx context.Context, req *interact.FavoriteListRequest) (resp *interact.FavoriteListResponse, err error) {
	resp = new(interact.FavoriteListResponse)

	// 校验参数
	if err = req.IsValid(); err != nil {
		resp.StatusCode, resp.StatusMsg = utils.BuildStatus(err)
		return resp, nil
	}

	// 解析Token
	claim, err := jwt.ParseToken(req.Token)
	if err != nil {
		resp.StatusCode, resp.StatusMsg = utils.BuildStatus(err)
		return resp, nil
	}

	// 调用service层
	videos, err := service.NewFavoriteService(ctx).GetFavoriteList(req.UserId, claim.Id)
	if err != nil {
		resp.StatusCode, resp.StatusMsg = utils.BuildStatus(err)
		return resp, nil
	}

	// 包装正常响应
	resp.StatusCode, resp.StatusMsg = utils.BuildStatus(errno.Success)
	resp.VideoList = videos

	return resp, nil
}

// CommentAction implements the InteractServiceImpl interface.
func (s *InteractServiceImpl) CommentAction(ctx context.Context, req *interact.CommentActionRequest) (resp *interact.CommentActionResponse, err error) {
	resp = new(interact.CommentActionResponse)

	// 校验参数
	if err = req.IsValid(); err != nil {
		resp.StatusCode, resp.StatusMsg = utils.BuildStatus(err)
		return resp, nil
	}

	// 解析Token
	claim, err := jwt.ParseToken(req.Token)
	if err != nil {
		resp.StatusCode, resp.StatusMsg = utils.BuildStatus(err)
		return resp, nil
	}

	// 调用service层
	err = service.NewCommentService(ctx).CommentAction(req, claim.Id)
	if err != nil {
		resp.StatusCode, resp.StatusMsg = utils.BuildStatus(err)
		return resp, nil
	}

	// 包装正常响应
	resp.StatusCode, resp.StatusMsg = utils.BuildStatus(errno.Success)
	return resp, nil
}

// GetCommentList implements the InteractServiceImpl interface.
func (s *InteractServiceImpl) GetCommentList(ctx context.Context, req *interact.CommentListRequest) (resp *interact.CommentListResponse, err error) {
	resp = new(interact.CommentListResponse)

	// 校验参数
	if err = req.IsValid(); err != nil {
		resp.StatusCode, resp.StatusMsg = utils.BuildStatus(err)
		return resp, nil
	}

	// 解析Token获取id，若空默认为0
	myID := int64(0)
	if len(req.Token) != 0 {
		claim, err := jwt.ParseToken(req.Token)
		if err != nil {
			resp.StatusCode, resp.StatusMsg = utils.BuildStatus(err)
			return resp, nil
		}
		myID = claim.Id
	}
	if err != nil {
		resp.StatusCode, resp.StatusMsg = utils.BuildStatus(err)
		return resp, nil
	}

	// 调用service层
	comments, err := service.NewCommentService(ctx).GetCommentList(req.VideoId, myID)
	if err != nil {
		resp.StatusCode, resp.StatusMsg = utils.BuildStatus(err)
		return resp, nil
	}

	// 包装正常响应
	resp.StatusCode, resp.StatusMsg = utils.BuildStatus(errno.Success)
	resp.CommentList = comments

	return resp, nil
}

// GetVideoInteract implements the InteractServiceImpl interface.
func (s *InteractServiceImpl) GetVideoInteract(ctx context.Context, req *interact.VideoInteractRequest) (resp *interact.VideoInteractResponse, err error) {
	resp = new(interact.VideoInteractResponse)

	// 校验参数
	if err = req.IsValid(); err != nil {
		resp.StatusCode, resp.StatusMsg = utils.BuildStatus(err)
		return resp, nil
	}

	// 调用service层
	favoriteCount, commentCount, isFavorite, err := service.NewGetVideoInteractService(ctx).GetVideoInteract(req)
	if err != nil {
		resp.StatusCode, resp.StatusMsg = utils.BuildStatus(err)
		return resp, nil
	}

	// 包装正常响应
	resp.StatusCode, resp.StatusMsg = utils.BuildStatus(errno.Success)
	resp.FavoriteCount = int64(favoriteCount)
	resp.CommentCount = int64(commentCount)
	resp.IsFavorite = isFavorite

	return resp, nil
}
