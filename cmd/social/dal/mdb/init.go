package mdb

import (
	"context"
	"github.com/hcdoit/tiktok/pkg/consts"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MDB *mongo.Client

func Init() {
	clientOptions := options.Client().ApplyURI(consts.MongoUri)

	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		panic(err)
	}
	if err = client.Ping(context.Background(), nil); err != nil {
		panic(err)
	}
	MDB = client
}
