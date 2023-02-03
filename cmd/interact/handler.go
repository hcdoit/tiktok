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
	// TODO: Your code here...
	resp = new(interact.FavoriteActionResponse)
	claim, err := jwt.ParseToken(req.Token)
	if err != nil {
		resp.StatusCode, resp.StatusMsg = utils.BuildStatus(err)
		return resp, nil
	}
	if err = req.IsValid(); err != nil {
		resp.StatusCode, resp.StatusMsg = utils.BuildStatus(err)
		return resp, nil
	}
	err = service.NewFavoriteService(ctx).FavoriteAction(req, claim.Id)
	if err != nil {
		resp.StatusCode, resp.StatusMsg = utils.BuildStatus(err)
		return resp, nil
	}
	resp.StatusCode, resp.StatusMsg = utils.BuildStatus(errno.Success)
	return resp, nil
}

// GetFavoriteList implements the InteractServiceImpl interface.
func (s *InteractServiceImpl) GetFavoriteList(ctx context.Context, req *interact.FavoriteListRequest) (resp *interact.FavoriteListResponse, err error) {
	// TODO: Your code here...
	resp = new(interact.FavoriteListResponse)
	claim, err := jwt.ParseToken(req.Token)
	if err != nil {
		resp.StatusCode, resp.StatusMsg = utils.BuildStatus(err)
		return resp, nil
	}
	if err = req.IsValid(); err != nil {
		resp.StatusCode, resp.StatusMsg = utils.BuildStatus(err)
		return resp, nil
	}
	videos, err := service.NewFavoriteService(ctx).GetFavoriteList(req.UserId, claim.Id)
	if err != nil {
		resp.StatusCode, resp.StatusMsg = utils.BuildStatus(err)
		return resp, nil
	}
	resp.StatusCode, resp.StatusMsg = utils.BuildStatus(errno.Success)
	resp.VideoList = videos

	return resp, nil
}

// CommentAction implements the InteractServiceImpl interface.
func (s *InteractServiceImpl) CommentAction(ctx context.Context, req *interact.CommentActionRequest) (resp *interact.CommentActionResponse, err error) {
	// TODO: Your code here...
	resp = new(interact.CommentActionResponse)
	claim, err := jwt.ParseToken(req.Token)
	if err != nil {
		resp.StatusCode, resp.StatusMsg = utils.BuildStatus(err)
		return resp, nil
	}
	if err = req.IsValid(); err != nil {
		resp.StatusCode, resp.StatusMsg = utils.BuildStatus(err)
		return resp, nil
	}
	err = service.NewCommentService(ctx).CommentAction(req, claim.Id)
	if err != nil {
		resp.StatusCode, resp.StatusMsg = utils.BuildStatus(err)
		return resp, nil
	}
	resp.StatusCode, resp.StatusMsg = utils.BuildStatus(errno.Success)
	return resp, nil
}

// GetCommentList implements the InteractServiceImpl interface.
func (s *InteractServiceImpl) GetCommentList(ctx context.Context, req *interact.CommentListRequest) (resp *interact.CommentListResponse, err error) {
	// TODO: Your code here...
	resp = new(interact.CommentListResponse)
	if err = req.IsValid(); err != nil {
		resp.StatusCode, resp.StatusMsg = utils.BuildStatus(err)
		return resp, nil
	}
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
	if err = req.IsValid(); err != nil {
		resp.StatusCode, resp.StatusMsg = utils.BuildStatus(err)
		return resp, nil
	}
	comments, err := service.NewCommentService(ctx).GetCommentList(req.VideoId, myID)
	if err != nil {
		resp.StatusCode, resp.StatusMsg = utils.BuildStatus(err)
		return resp, nil
	}
	resp.StatusCode, resp.StatusMsg = utils.BuildStatus(errno.Success)
	resp.CommentList = comments

	return resp, nil
}

// GetVideoInteract implements the InteractServiceImpl interface.
func (s *InteractServiceImpl) GetVideoInteract(ctx context.Context, req *interact.VideoInteractRequest) (resp *interact.VideoInteractResponse, err error) {
	// TODO: Your code here...
	resp = new(interact.VideoInteractResponse)
	if err = req.IsValid(); err != nil {
		resp.StatusCode, resp.StatusMsg = utils.BuildStatus(err)
		return resp, nil
	}
	favoriteCount, commentCount, isFavorite, err := service.NewGetVideoInteractService(ctx).GetVideoInteract(req)
	if err != nil {
		resp.StatusCode, resp.StatusMsg = utils.BuildStatus(err)
		return resp, nil
	}
	resp.StatusCode, resp.StatusMsg = utils.BuildStatus(errno.Success)
	resp.FavoriteCount = int64(favoriteCount)
	resp.CommentCount = int64(commentCount)
	resp.IsFavorite = isFavorite

	return resp, nil
}
