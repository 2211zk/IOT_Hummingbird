package data

import (
	"kratos/internal/conf"

	"context"
	"fmt"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewGreeterRepo)

// Data .
type Data struct {
	// TODO wrapped database client
	MongoClient *mongo.Client
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	// 构建MongoDB连接URI
	uri := fmt.Sprintf("mongodb://%s:%s@%s:%d/%s",
		c.Mongodb.User,
		c.Mongodb.Password,
		c.Mongodb.Host,
		c.Mongodb.Port,
		c.Mongodb.Database,
	)
	clientOpts := options.Client().ApplyURI(uri)
	mongoClient, err := mongo.Connect(context.Background(), clientOpts)
	if err != nil {
		return nil, nil, err
	}
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
		_ = mongoClient.Disconnect(context.Background())
	}
	return &Data{MongoClient: mongoClient}, cleanup, nil
}
