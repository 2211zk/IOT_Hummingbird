# Cobra Script Center - 快速开始指南

## 安装和初始化

1. **构建应用程序**
```bash
cd IOT_client/kratos_end/cobra-script-center
make build
```

2. **初始化数据库**
```bash
./bin/script-center migrate
```

3. **创建管理员用户**
```bash
./bin/script-center user create --username admin --role admin --password admin123
```

## 基本使用

### 1. 创建脚本

创建一个简单的bash脚本：
```bash
./bin/script-center script create --name hello-world --language bash --description "Hello world script"
```

创建一个Python脚本：
```bash
./bin/script-center script create --name hello-python --language python --description "Python hello script" --tags demo,python
```

### 2. 查看脚本列表

```bash
./bin/script-center script list
```

按语言过滤：
```bash
./bin/script-center script list --language python
```

### 3. 执行脚本

执行脚本：
```bash
./bin/script-center script run hello-world
```

带参数执行：
```bash
./bin/script-center script run hello-world --param NAME=张三 --param MESSAGE=欢迎使用脚本中心
```

### 4. 查看执行历史

```bash
./bin/script-center execution list
```

查看特定脚本的执行历史：
```bash
./bin/script-center execution list --script hello-world
```

查看执行详情：
```bash
./bin/script-center execution show <execution-id>
```

### 5. 定时任务

创建定时任务（每小时执行一次）：
```bash
./bin/script-center schedule create hello-world "0 * * * *"
```

查看定时任务：
```bash
./bin/script-center schedule list
```

启动守护进程处理定时任务：
```bash
./bin/script-center daemon start
```

### 6. 用户管理

创建普通用户：
```bash
./bin/script-center user create --username developer --role user
```

查看用户列表：
```bash
./bin/script-center user list
```

## 支持的脚本语言

- **Bash** (.sh) - Linux/Unix shell脚本
- **Python** (.py) - Python 3脚本
- **Node.js** (.js) - JavaScript脚本
- **Go** (.go) - Go语言脚本
- **PowerShell** (.ps1) - Windows PowerShell脚本

## 配置文件

配置文件位于 `.script-center.yaml`：

```yaml
database:
  driver: sqlite3
  dsn: ./script-center.db

server:
  host: localhost
  port: 8080

security:
  jwt_secret: "your-secret-key"
  password_salt: "your-salt"
  max_executions: 10

logging:
  level: info
  format: json
  file: ""
```

## 常用命令总结

| 命令 | 描述 |
|------|------|
| `script-center migrate` | 初始化数据库 |
| `script-center user create` | 创建用户 |
| `script-center script create` | 创建脚本 |
| `script-center script list` | 列出脚本 |
| `script-center script run` | 执行脚本 |
| `script-center execution list` | 查看执行历史 |
| `script-center schedule create` | 创建定时任务 |
| `script-center daemon start` | 启动守护进程 |

## 示例脚本

项目包含了一些示例脚本在 `examples/` 目录下：

- `hello.sh` - Bash示例脚本
- `hello.py` - Python示例脚本

你可以将这些脚本的内容复制到通过CLI创建的脚本中进行测试。

## 故障排除

1. **数据库连接问题**：确保SQLite3已安装
2. **脚本执行失败**：检查脚本语言的运行时是否已安装（python3, node, go等）
3. **权限问题**：确保脚本文件有执行权限
4. **端口占用**：如果8080端口被占用，修改配置文件中的端口号

## 下一步

- 查看完整的 [README.md](README.md) 了解更多功能
- 探索 [API文档](docs/) 了解REST API接口
- 查看 [架构文档](docs/architecture.md) 了解系统设计