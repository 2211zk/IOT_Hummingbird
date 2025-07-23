package data

import (
	"IOT_Hummingbird_back_end/internal/conf"

	"context"
	"fmt"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewGreeterRepo, NewUserRepo)

// Data .
type Data struct {
	MySQL *gorm.DB
	Redis *redis.Client
	Mongo *mongo.Client
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	helper := log.NewHelper(logger)
	// 1. 使用 viper 解析 config.yaml
	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath("./configs")
	if err := v.ReadInConfig(); err != nil {
		helper.Errorf("读取配置文件失败: %v", err)
		return nil, nil, err
	}
	// 2. 解析 MySQL 配置
	mysqlCfg := c.GetMysql()
	mysqlDsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=True&loc=Local", mysqlCfg.User, mysqlCfg.Password, mysqlCfg.Host, mysqlCfg.Port, mysqlCfg.Database)
	mysqlDB, err := gorm.Open(mysql.Open(mysqlDsn), &gorm.Config{})
	if err != nil {
		helper.Errorf("MySQL 连接失败: %v", err)
		return nil, nil, err
	} else {
		helper.Infof("MySQL 连接成功")
	}
	// 3. 解析 Redis 配置
	redisCfg := c.GetRedis()
	redisClient := redis.NewClient(&redis.Options{
		Addr:     redisCfg.Addr,
		Password: redisCfg.Password,
		DB:       int(redisCfg.Db),
	})
	if err := redisClient.Ping(context.Background()).Err(); err != nil {
		helper.Errorf("Redis 连接失败: %v", err)
		return nil, nil, err
	} else {
		helper.Infof("Redis 连接成功，地址: %s", redisCfg.Addr)
	}
	// 4. 解析 MongoDB 配置
	mongoCfg := c.GetMongodb()
	mongoUri := fmt.Sprintf("mongodb://%s:%s@%s:%d/%s", mongoCfg.User, mongoCfg.Password, mongoCfg.Host, mongoCfg.Port, mongoCfg.Database)
	mongoClient, err := mongo.Connect(context.Background(), options.Client().ApplyURI(mongoUri))
	if err != nil {
		helper.Errorf("MongoDB 连接失败: %v", err)
		return nil, nil, err
	} else {
		helper.Infof("MongoDB 连接成功，地址: %s:%d", mongoCfg.Host, mongoCfg.Port)
	}
	cleanup := func() {
		helper.Info("closing the data resources")
		_ = mongoClient.Disconnect(context.Background())
		_ = redisClient.Close()
		// gorm 无需手动关闭
	}
	return &Data{
		MySQL: mysqlDB,
		Redis: redisClient,
		Mongo: mongoClient,
	}, cleanup, nil
}
