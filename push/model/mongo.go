package model

import (
	"context"
	"gmimo/push/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	MongoPool       *mongo.Client
	MongoDB         *mongo.Database
	MongoCollection *mongo.Collection
)

// 创建Mongo Client
func InitMongoClient(cf config.ConfMongo) (err error) {
	// 加载连接Mongo配置信息
	opts := options.Client().ApplyURI(cf.Uri)
	opts.SetConnectTimeout(cf.ConnTimeout).
		SetMaxPoolSize(cf.MaxPoolSize)

	// 创建连接
	ctx, _ := context.WithTimeout(context.Background(), cf.ConnTimeout)
	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		return err
	}
	// 验证连接是否可用
	if err = client.Ping(ctx, nil); err != nil {
		return err
	}
	MongoPool = client
	MongoDB = client.Database(cf.Dbname)
	MongoCollection = client.Database(cf.Dbname).Collection(cf.Collection)
	return
}
