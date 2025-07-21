# Wanlian_resource MongoDB数据源使用说明

## 概述

本项目已经集成了Wanlian_resource MongoDB数据源，用于连接和管理MongoDB数据库。数据源名称为`Wanlian_resource`，提供了完整的CRUD操作和索引管理功能。

## 文件结构

```
server/
├── initialize/
│   ├── wanlian_mongo.go          # Wanlian_resource数据源连接配置
│   └── wanlian_init.go           # Wanlian_resource数据源初始化
├── global/
│   └── global.go                 # 全局变量定义（包含GVA_WANLIAN_MONGO）
└── README_Wanlian_Resource.md    # 使用说明文档
```

## 配置说明

### 1. 数据库连接配置

在 `initialize/wanlian_mongo.go` 文件中的 `GetWanlianResourceConfig()` 函数中配置：

```go
func GetWanlianResourceConfig() *WanlianResourceConfig {
    return &WanlianResourceConfig{
        Database:         "Wanlian_resource", // 数据源名称
        Username:         "zhangkai",          // MongoDB用户名
        Password:         "Zhangkai123",       // MongoDB密码
        AuthSource:       "admin",            // 认证数据库
        MinPoolSize:      0,                  // 最小连接池
        MaxPoolSize:      100,                // 最大连接池
        SocketTimeoutMs:  0,                  // socket超时时间
        ConnectTimeoutMs: 0,                  // 连接超时时间
        Hosts: []*WanlianMongoHost{
            {
                Host: "14.103.143.229", // MongoDB主机地址
                Port: "27017",           // MongoDB端口
            },
        },
        Options: "", // MongoDB连接选项
    }
}
```

### 2. 当前连接配置

- **Host**: 14.103.143.229
- **Port**: 27017
- **Username**: zhangkai
- **Password**: Zhangkai123
- **Database**: Wanlian_resource
- **AuthSource**: admin

### 3. 修改连接配置

如果需要修改连接配置，请在 `initialize/wanlian_mongo.go` 文件中修改以下参数：

- `Host`: MongoDB服务器地址
- `Port`: MongoDB服务器端口
- `Username`: MongoDB用户名
- `Password`: MongoDB密码
- `Database`: 数据库名称（默认为"Wanlian_resource"）

## 使用方法

### 1. 获取MongoDB客户端

```go
import "github.com/flipped-aurora/gin-vue-admin/server/global"

// 获取Wanlian_resource MongoDB客户端
client := global.GetWanlianMongoClient()
```

### 2. 获取集合对象

```go
import "github.com/flipped-aurora/gin-vue-admin/server/initialize"

// 获取指定集合
collection := initialize.GetWanlianCollection("users")
```

### 3. 基本CRUD操作

#### 创建文档

```go
ctx := context.Background()
user := &User{
    Username: "testuser",
    Email:    "test@example.com",
    CreatedAt: time.Now(),
    UpdatedAt: time.Now(),
}

err := collection.InsertOne(ctx, user)
```

#### 查询文档

```go
// 查询单个文档
var user User
err := collection.Find(ctx, bson.M{"username": "testuser"}).One(&user)

// 查询多个文档
var users []User
err := collection.Find(ctx, bson.M{"status": 1}).All(&users)
```

#### 更新文档

```go
// 更新单个文档
err := collection.UpdateOne(ctx, 
    bson.M{"username": "testuser"}, 
    bson.M{"$set": bson.M{"email": "newemail@example.com"}},
)
```

#### 删除文档

```go
// 删除单个文档
err := collection.Remove(ctx, bson.M{"username": "testuser"})

// 删除多个文档
_, err := collection.RemoveAll(ctx, bson.M{"status": 0})
```

### 4. 高级查询

#### 聚合查询

```go
pipeline := []bson.M{
    {"$group": bson.M{
        "_id": "$category",
        "count": bson.M{"$sum": 1},
        "avgPrice": bson.M{"$avg": "$price"},
    }},
    {"$sort": bson.M{"count": -1}},
}

var results []bson.M
err := collection.Aggregate(ctx, pipeline).All(&results)
```

#### 带选项的查询

```go
err := collection.Find(ctx, bson.M{"status": 1}).
    Sort("-created_at").
    Limit(10).
    All(&users)
```

### 5. 索引管理

#### 创建索引

```go
// 创建单字段索引
err := collection.CreateOneIndex(ctx, qmgo.IndexModel{
    Key: []string{"username"},
})

// 创建复合索引
err := collection.CreateOneIndex(ctx, qmgo.IndexModel{
    Key: []string{"category", "price"},
})

// 创建唯一索引
err := collection.CreateOneIndex(ctx, qmgo.IndexModel{
    Key: []string{"email"},
}, qmgo.IndexOptions{
    Unique: true,
})
```

#### 删除索引

```go
err := collection.DropIndex(ctx, "index_name")
```

## 初始化流程

1. 系统启动时自动调用 `initialize.WanlianResourceInit()`
2. 连接到MongoDB数据库 (14.103.143.229:27017)
3. 使用zhangkai用户进行认证
4. 创建必要的索引
5. 将连接存储到全局变量 `global.GVA_WANLIAN_MONGO`

## 错误处理

```go
// 检查连接是否成功
if global.GVA_WANLIAN_MONGO == nil {
    panic("Wanlian_resource MongoDB client no init")
}

// 错误处理示例
err := collection.InsertOne(ctx, user)
if err != nil {
    global.GVA_LOG.Error("创建用户失败", zap.Error(err))
    return err
}
```

## 性能优化

1. **连接池配置**: 根据实际需求调整 `MinPoolSize` 和 `MaxPoolSize`
2. **索引优化**: 为常用查询字段创建索引
3. **批量操作**: 使用 `InsertMany` 进行批量插入
4. **查询优化**: 使用投影和限制来减少数据传输

## 注意事项

1. 确保MongoDB服务器正在运行 (14.103.143.229:27017)
2. 检查网络连接和防火墙设置
3. 验证用户名和密码 (zhangkai/Zhangkai123)
4. 定期备份数据库
5. 监控连接池使用情况

## 故障排除

### 连接失败

1. 检查MongoDB服务器是否运行 (14.103.143.229:27017)
2. 验证主机地址和端口
3. 检查网络连接
4. 确认用户名和密码 (zhangkai/Zhangkai123)

### 查询性能问题

1. 检查索引是否正确创建
2. 使用 `explain()` 分析查询计划
3. 优化查询条件
4. 考虑使用聚合管道

### 内存问题

1. 调整连接池大小
2. 使用分页查询
3. 限制查询结果数量
4. 定期清理连接

## 连接测试

启动服务器后，检查日志中是否有：
```
Wanlian_resource MongoDB初始化成功!
```

如果没有看到这个日志，请检查：
1. MongoDB服务器是否运行 (14.103.143.229:27017)
2. 网络连接是否正常
3. 用户名密码是否正确 (zhangkai/Zhangkai123)
4. 防火墙是否允许27017端口连接 