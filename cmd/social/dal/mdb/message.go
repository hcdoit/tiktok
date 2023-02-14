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
	IsRead     bool
	ToUserID   int64
	FromUserID int64
	Content    string
}

func GetMessages(ctx context.Context, fromUserID int64, toUserID int64) ([]*Message, error) {
	collection := fmt.Sprintf("chat:%d-%d", fromUserID, toUserID)
	if toUserID < fromUserID {
		collection = fmt.Sprintf("chat:%d-%d", toUserID, fromUserID)
	}
	klog.Info(collection)
	coll := MDB.Database(consts.MongoDataBase).Collection(collection)
	cursor, err := coll.Find(ctx, bson.M{"$and": []bson.M{{"isread": false}, {"touserid": fromUserID}}})
	if err != nil {
		return nil, err
	}
	messages := make([]*Message, 0)
	err = cursor.All(ctx, &messages)
	klog.Info(len(messages))
	if err != nil {
		return nil, err
	}
	many, err := coll.UpdateMany(ctx, bson.M{"$and": []bson.M{{"isread": false}, {"touserid": fromUserID}}}, bson.M{"$set": bson.M{"isread": true}})
	klog.Info(many.UpsertedCount)
	if err != nil {
		return nil, err
	}
	return messages, nil

}

func InsertMessage(ctx context.Context, fromUserID int64, toUserID int64, content string) error {
	collection := fmt.Sprintf("chat:%d-%d", fromUserID, toUserID)
	if toUserID < fromUserID {
		collection = fmt.Sprintf("chat:%d-%d", toUserID, fromUserID)
	}
	coll := MDB.Database(consts.MongoDataBase).Collection(collection)
	_, err := coll.InsertOne(ctx, &Message{
		CreateTime: time.Now(),
		IsRead:     false,
		ToUserID:   toUserID,
		FromUserID: fromUserID,
		Content:    content,
	})

	return err
}
