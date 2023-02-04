package service

import (
	"context"
	"github.com/hcdoit/tiktok/cmd/social/dal/mdb"
	"github.com/hcdoit/tiktok/cmd/social/utils"
	"github.com/hcdoit/tiktok/kitex_gen/social"
)

type MessageService struct {
	ctx context.Context
}

// NewMessageService new MessageService
func NewMessageService(ctx context.Context) *MessageService {
	return &MessageService{
		ctx: ctx,
	}
}

func (s *MessageService) GetMessageChat(req *social.MessageChatRequest, myID int64) ([]*social.Message, error) {
	modelMessages, err := mdb.GetMessages(s.ctx, myID, req.ToUserId)
	if err != nil {
		return nil, err
	}
	return utils.BuildMessages(modelMessages), nil
}

func (s *MessageService) MessageAction(req *social.MessageActionRequest, myID int64) error {
	return mdb.InsertMessage(s.ctx, myID, req.ToUserId, req.Content)
}
