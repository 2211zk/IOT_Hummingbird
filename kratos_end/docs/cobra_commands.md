# Cobra 脚本中心命令文档

本项目已集成 [Cobra](https://github.com/spf13/cobra) 命令行脚本中心，支持多命令扩展，方便服务启动、数据迁移、批量任务等运维操作。

## 1. 启动服务

```sh
# 启动 HTTP/gRPC 服务
 go run ./cmd/IOT_Hummingbird_back_end/main.go serve
```

## 2. 查看所有命令

```sh
go run ./cmd/IOT_Hummingbird_back_end/main.go --help
```

输出示例：
```
Hummingbird 脚本中心，支持服务启动、数据迁移、批量任务等自定义命令。

Usage:
  hummingbird [command]

Available Commands:
  serve       启动服务
  help        Help about any command

Flags:
  -h, --help   help for hummingbird

Use "hummingbird [command] --help" for more information about a command.
```

## 3. 扩展自定义命令

你可以在 `cmd/IOT_Hummingbird_back_end/root.go` 文件中添加自定义命令，例如：

```go
var migrateCmd = &cobra.Command{
    Use:   "migrate",
    Short: "数据库迁移",
    Run: func(cmd *cobra.Command, args []string) {
        // 你的迁移逻辑
    },
}

func init() {
    rootCmd.AddCommand(migrateCmd)
}
```

然后即可通过如下方式调用：
```sh
go run ./cmd/IOT_Hummingbird_back_end/main.go migrate
```

## 4. 常见用法
- `serve`：启动主服务
- `migrate`：数据库迁移（需自行实现）
- `task`：批量任务（可自定义）

## 5. 参考资料
- [Cobra 官方文档](https://cobra.dev/)
- [Kratos 官方文档](https://go-kratos.dev/)

---
如需添加更多命令或有脚本中心扩展需求，请联系开发者。 