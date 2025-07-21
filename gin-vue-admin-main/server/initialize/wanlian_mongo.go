package initialize

import (
	"context"
	"fmt"
	"sort"
	"strings"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/pkg/errors"
	"github.com/qiniu/qmgo"
	"github.com/qiniu/qmgo/options"
	"go.mongodb.org/mongo-driver/bson"
	option "go.mongodb.org/mongo-driver/mongo/options"
)

var WanlianMongo = new(wanlianMongo)

type (
	wanlianMongo struct{}
	WanlianIndex struct {
		V    any      `bson:"v"`
		Ns   any      `bson:"ns"`
		Key  []bson.E `bson:"key"`
		Name string   `bson:"name"`
	}
)

// WanlianResourceConfig Wanlian_resource数据源配置
type WanlianResourceConfig struct {
	Database         string
	Username         string
	Password         string
	AuthSource       string
	MinPoolSize      uint64
	MaxPoolSize      uint64
	SocketTimeoutMs  int64
	ConnectTimeoutMs int64
	Hosts            []*WanlianMongoHost
	Options          string
}

type WanlianMongoHost struct {
	Host string
	Port string
}

// GetWanlianResourceConfig 获取Wanlian_resource数据源配置
func GetWanlianResourceConfig() *WanlianResourceConfig {
	return &WanlianResourceConfig{
		Database:         "Wanlian_resource", // 数据源名称
		Username:         "zhangkai",         // MongoDB用户名
		Password:         "Zhangkai123",      // MongoDB密码
		AuthSource:       "admin",            // 认证数据库
		MinPoolSize:      0,                  // 最小连接池
		MaxPoolSize:      100,                // 最大连接池
		SocketTimeoutMs:  0,                  // socket超时时间
		ConnectTimeoutMs: 0,                  // 连接超时时间
		Hosts: []*WanlianMongoHost{
			{
				Host: "14.103.143.229", // MongoDB主机地址
				Port: "27017",          // MongoDB端口
			},
		},
		Options: "", // MongoDB连接选项
	}
}

// Uri 生成MongoDB连接URI
func (x *WanlianResourceConfig) Uri() string {
	length := len(x.Hosts)
	hosts := make([]string, 0, length)
	for i := 0; i < length; i++ {
		if x.Hosts[i].Host != "" && x.Hosts[i].Port != "" {
			hosts = append(hosts, x.Hosts[i].Host+":"+x.Hosts[i].Port)
		}
	}
	if x.Options != "" {
		return fmt.Sprintf("mongodb://%s/%s?%s", strings.Join(hosts, ","), x.Database, x.Options)
	}
	return fmt.Sprintf("mongodb://%s/%s", strings.Join(hosts, ","), x.Database)
}

// Indexes 创建Wanlian_resource数据源的索引
func (m *wanlianMongo) Indexes(ctx context.Context) error {
	// 表名:索引列表 列: "表名": [][]string{{"index1", "index2"}}
	indexMap := map[string][][]string{
		// 在这里添加Wanlian_resource数据源的索引配置
		// 例如：
		// "users": [][]string{{"email"}, {"username", "status"}},
		// "products": [][]string{{"category"}, {"price", "created_at"}},
	}

	for collection, indexes := range indexMap {
		err := m.CreateIndexes(ctx, collection, indexes)
		if err != nil {
			return err
		}
	}
	return nil
}

// Initialization 初始化Wanlian_resource数据源
func (m *wanlianMongo) Initialization() error {
	var opts []options.ClientOptions
	// 如果需要开启zap日志，可以在这里配置
	// if global.GVA_CONFIG.Mongo.IsZap {
	//     opts = internal.Mongo.GetClientOptions()
	// }

	ctx := context.Background()
	config := GetWanlianResourceConfig()

	qmgoConfig := &qmgo.Config{
		Uri:              config.Uri(),
		Coll:             "", // 这里可以为空，因为我们会为每个集合单独创建连接
		Database:         config.Database,
		MinPoolSize:      &config.MinPoolSize,
		MaxPoolSize:      &config.MaxPoolSize,
		SocketTimeoutMS:  &config.SocketTimeoutMs,
		ConnectTimeoutMS: &config.ConnectTimeoutMs,
	}

	// 如果需要认证
	if config.Username != "" && config.Password != "" {
		qmgoConfig.Auth = &qmgo.Credential{
			Username:   config.Username,
			Password:   config.Password,
			AuthSource: config.AuthSource,
		}
	}

	client, err := qmgo.Open(ctx, qmgoConfig, opts...)
	if err != nil {
		return errors.Wrap(err, "连接Wanlian_resource MongoDB数据库失败!")
	}

	// 将连接存储到全局变量中
	global.GVA_WANLIAN_MONGO = client

	// 创建索引
	err = m.Indexes(ctx)
	if err != nil {
		return err
	}

	return nil
}

// CreateIndexes 为Wanlian_resource数据源创建索引
func (m *wanlianMongo) CreateIndexes(ctx context.Context, name string, indexes [][]string) error {
	collection, err := global.GVA_WANLIAN_MONGO.Database.Collection(name).CloneCollection()
	if err != nil {
		return errors.Wrapf(err, "获取[%s]的表对象失败!", name)
	}

	list, err := collection.Indexes().List(ctx)
	if err != nil {
		return errors.Wrapf(err, "获取[%s]的索引对象失败!", name)
	}

	var entities []WanlianIndex
	err = list.All(ctx, &entities)
	if err != nil {
		return errors.Wrapf(err, "获取[%s]的索引列表失败!", name)
	}

	length := len(indexes)
	indexMap1 := make(map[string][]string, length)
	for i := 0; i < length; i++ {
		sort.Strings(indexes[i]) // 对索引key进行排序, 在使用bson.M搜索时, bson会自动按照key的字母顺序进行排序
		length1 := len(indexes[i])
		keys := make([]string, 0, length1)
		for j := 0; j < length1; j++ {
			if indexes[i][j][0] == '-' {
				keys = append(keys, indexes[i][j], "-1")
				continue
			}
			keys = append(keys, indexes[i][j], "1")
		}
		key := strings.Join(keys, "_")
		_, o1 := indexMap1[key]
		if o1 {
			return errors.Errorf("索引[%s]重复!", key)
		}
		indexMap1[key] = indexes[i]
	}

	length = len(entities)
	indexMap2 := make(map[string]map[string]string, length)
	for i := 0; i < length; i++ {
		v1, o1 := indexMap2[entities[i].Name]
		if !o1 {
			keyLength := len(entities[i].Key)
			v1 = make(map[string]string, keyLength)
			for j := 0; j < keyLength; j++ {
				v2, o2 := v1[entities[i].Key[j].Key]
				if !o2 {
					v1 = make(map[string]string)
				}
				v2 = entities[i].Key[j].Key
				v1[entities[i].Key[j].Key] = v2
				indexMap2[entities[i].Name] = v1
			}
		}
	}

	for k1, v1 := range indexMap1 {
		_, o2 := indexMap2[k1]
		if o2 {
			continue
		} // 索引存在

		if len(fmt.Sprintf("%s.%s.$%s", collection.Name(), name, v1)) > 127 {
			err = global.GVA_WANLIAN_MONGO.Database.Collection(name).CreateOneIndex(ctx, options.IndexModel{
				Key:          v1,
				IndexOptions: option.Index().SetName(utils.MD5V([]byte(k1))),
			})
			if err != nil {
				return errors.Wrapf(err, "创建索引[%s]失败!", k1)
			}
			return nil
		}

		err = global.GVA_WANLIAN_MONGO.Database.Collection(name).CreateOneIndex(ctx, options.IndexModel{
			Key:          v1,
			IndexOptions: option.Index().SetExpireAfterSeconds(86400),
		})
		if err != nil {
			return errors.Wrapf(err, "创建索引[%s]失败!", k1)
		}
	}
	return nil
}

// GetWanlianMongoClient 获取Wanlian_resource MongoDB客户端
func GetWanlianMongoClient() *qmgo.QmgoClient {
	return global.GVA_WANLIAN_MONGO
}

// GetWanlianCollection 获取Wanlian_resource数据源的集合
func GetWanlianCollection(collectionName string) *qmgo.Collection {
	return global.GVA_WANLIAN_MONGO.Database.Collection(collectionName)
}
