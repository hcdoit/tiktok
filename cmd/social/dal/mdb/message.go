package mdb

import (
	"context"
	"fmt"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/hcdoit/tiktok/pkg/consts"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

type Message struct {
	CreateTime time.Time
	ToUserID   int64
	FromUserID int64
	Content    string
}

func GetMessages(ctx context.Context, fromUserID int64, toUserID int64) (messages []*Message, err error) {
	collection := fmt.Sprintf("chat:%d-%d", fromUserID, toUserID)
	if toUserID < fromUserID {
		collection = fmt.Sprintf("%d-%d", toUserID, fromUserID)
	}
	klog.Info(collection)
	coll := MDB.Database(consts.MongoDataBase).Collection(collection)
	cursor, err := coll.Find(ctx, bson.D{})
	klog.Info(err)
	if err != nil {
		return nil, err
	}
	err = cursor.All(ctx, &messages)
	if err != nil {
		return nil, err
	}
	return messages, nil

}

func InsertMessage(ctx context.Context, fromUserID int64, toUserID int64, content string) error {
	collection := fmt.Sprintf("chat:%d-%d", fromUserID, toUserID)
	if toUserID < fromUserID {
		collection = fmt.Sprintf("%d-%d", toUserID, fromUserID)
	}
	coll := MDB.Database(consts.MongoDataBase).Collection(collection)
	_, err := coll.InsertOne(ctx, &Message{
		CreateTime: time.Now(),
		ToUserID:   toUserID,
		FromUserID: fromUserID,
		Content:    content,
	})

	return err
}
