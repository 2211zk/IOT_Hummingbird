package data

import (
	"context"

	"IOT_Hummingbird_back_end/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type greeterRepo struct {
	data *Data
	log  *log.Helper
}

// NewGreeterRepo .
func NewGreeterRepo(data *Data, logger log.Logger) biz.GreeterRepo {
	return &greeterRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *greeterRepo) Save(ctx context.Context, g *biz.Greeter) (*biz.Greeter, error) {
	// MySQL 示例：执行一条原生 SQL 查询
	var version string
	err := r.data.MySQL.Raw("SELECT VERSION() as version").Scan(&version).Error
	if err != nil {
		r.log.Errorf("MySQL 查询失败: %v", err)
	} else {
		r.log.Infof("MySQL 连接成功，数据库版本: %s", version)
	}

	// Redis 示例：写入一条数据
	err = r.data.Redis.Set(ctx, "greeter:last_hello", g.Hello, 0).Err()
	if err != nil {
		r.log.Errorf("Redis 写入失败: %v", err)
	} else {
		r.log.Infof("Redis 写入成功: greeter:last_hello = %s", g.Hello)
	}

	// MongoDB 示例：获取数据库列表
	dbs, err := r.data.Mongo.ListDatabaseNames(ctx, nil)
	if err != nil {
		r.log.Errorf("MongoDB 查询失败: %v", err)
	} else {
		r.log.Infof("MongoDB 连接成功，数据库列表: %v", dbs)
	}

	return g, nil
}

func (r *greeterRepo) Update(ctx context.Context, g *biz.Greeter) (*biz.Greeter, error) {
	return g, nil
}

func (r *greeterRepo) FindByID(context.Context, int64) (*biz.Greeter, error) {
	return nil, nil
}

func (r *greeterRepo) ListByHello(context.Context, string) ([]*biz.Greeter, error) {
	return nil, nil
}

func (r *greeterRepo) ListAll(context.Context) ([]*biz.Greeter, error) {
	return nil, nil
}
