package utils

import (
	"context"
	"errors"
	"github.com/hcdoit/tiktok/cmd/social/dal/mdb"
	"github.com/hcdoit/tiktok/cmd/social/rpc"
	"github.com/hcdoit/tiktok/kitex_gen/social"
	"github.com/hcdoit/tiktok/kitex_gen/user"
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

func BuildUsers(ids []int64, myID int64, ctx context.Context) []*user.User {
	users := make([]*user.User, 0)
	token, err := jwt.CreateToken(jwt.CustomClaims{
		Id: int64(myID),
	})
	if err != nil {
		return nil
	}
	for _, id := range ids {
		user, err := rpc.GetUser(ctx, &user.GetUserRequest{
			UserId: id,
			Token:  token,
		})
		if user != nil && err == nil {
			users = append(users, user)
		}

	}
	return users
}

// BuildMessage 打包单个Message
func BuildMessage(msg *mdb.Message, mid int64) *social.Message {
	if msg == nil {
		return nil
	}
	return &social.Message{
		Id:         mid,
		ToUserId:   msg.ToUserID,
		FromUserId: msg.FromUserID,
		Content:    msg.Content,
		CreateTime: msg.CreateTime.Format("2006-01-02 15:04:05"),
	}
}

// BuildMessages 打包Message列表
func BuildMessages(msgs []*mdb.Message) []*social.Message {
	messages := make([]*social.Message, 0)
	for i, msg := range msgs {
		if temp := BuildMessage(msg, int64(i+1)); temp != nil {
			messages = append(messages, temp)
		}
	}
	return messages
}
