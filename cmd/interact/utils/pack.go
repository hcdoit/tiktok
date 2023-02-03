package utils

import (
	"context"
	"errors"
	"github.com/hcdoit/tiktok/cmd/interact/dal/db"
	"github.com/hcdoit/tiktok/cmd/interact/rpc"
	"github.com/hcdoit/tiktok/kitex_gen/interact"
	"github.com/hcdoit/tiktok/kitex_gen/user"
	"github.com/hcdoit/tiktok/pkg/errno"
	"github.com/hcdoit/tiktok/pkg/jwt"
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

func BuildComment(c *db.Comment, ctx context.Context, myID int64) (*interact.Comment, error) {
	if c == nil {
		return nil, nil
	}
	token, err := jwt.CreateToken(jwt.CustomClaims{
		Id: int64(myID),
	})
	if err != nil {
		return nil, err
	}
	user, err := rpc.GetUser(ctx, &user.GetUserRequest{
		UserId: c.UserID,
		Token:  token,
	})
	if err != nil {
		return nil, err
	}
	return &interact.Comment{
		Id:         int64(c.ID),
		User:       user,
		Content:    c.CommentText,
		CreateDate: c.CreatedAt.Format("01-02"),
	}, nil
}

func BuildComments(cs []*db.Comment, ctx context.Context, myID int64) []*interact.Comment {
	comments := make([]*interact.Comment, 0)
	for _, c := range cs {
		if temp, err := BuildComment(c, ctx, myID); temp != nil && err == nil {
			comments = append(comments, temp)
		}
	}
	return comments
}
