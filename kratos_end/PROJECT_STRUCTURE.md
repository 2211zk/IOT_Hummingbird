# 项目目录结构说明

本项目为 IOT_Hummingbird_back_end，采用 Go 语言开发，目录结构清晰，便于扩展和维护。以下为各目录及主要文件的详细说明：

## 根目录

| 文件/文件夹         | 说明                                                         |
|---------------------|--------------------------------------------------------------|
| api/                | 存放 Protobuf 定义及生成的 gRPC/HTTP 代码，按业务模块划分。  |
| bin/                | 可执行文件输出目录。                                         |
| cmd/                | 主程序入口及命令行相关代码。                                 |
| configs/            | 配置文件目录。                                               |
| Dockerfile          | Docker 镜像构建文件。                                        |
| docs/               | 项目文档目录。                                               |
| go.mod/go.sum       | Go 依赖管理文件。                                            |
| internal/           | 内部应用核心代码，按领域分层。                               |
| LICENSE             | 许可证文件。                                                 |
| Makefile            | 自动化构建脚本。                                             |
| openapi.yaml        | OpenAPI (Swagger) 接口文档。                                 |
| README.md           | 项目说明文档。                                               |
| third_party/        | 第三方依赖的 proto 文件。                                    |

---

## 1. api/

- 主要用于存放 Protobuf 协议文件及其生成的代码，按业务模块（如 helloworld、user）和版本（如 v1）组织。
- 典型结构如下：

```
api/
  helloworld/
    v1/
      error_reason.pb.go
      error_reason.proto
      greeter_grpc.pb.go
      greeter_http.pb.go
      greeter.pb.go
      greeter.proto
  user/
    v1/
      user_grpc.pb.go
      user_http.pb.go
      user.pb.go
      user.proto
```

- `.proto` 文件为协议定义，`.pb.go` 为 Go 语言生成文件，`*_grpc.pb.go` 为 gRPC 相关代码，`*_http.pb.go` 为 HTTP 相关代码。

---

## 2. bin/

- 存放编译后的可执行文件，通常由 Makefile 或构建脚本自动生成。

---

## 3. cmd/

- 主程序入口，通常每个子文件夹对应一个可编译的应用。
- 例如：

```
cmd/
  IOT_Hummingbird_back_end/
    main.go         // 主入口
    root.go         // 命令行根命令
    wire.go         // 依赖注入相关
    wire_gen.go     // 依赖注入生成文件
```

---

## 4. configs/

- 存放配置文件，如 `config.yaml`，用于管理环境变量、数据库连接等配置信息。

---

## 5. docs/

- 存放项目相关文档，如 `cobra_commands.md` 记录命令行工具的用法。

---

## 6. internal/

- 项目核心业务代码，按领域分层，常见子目录有：

| 子目录      | 说明                                                         |
|-------------|--------------------------------------------------------------|
| biz/        | 领域业务逻辑层（如 greeter.go、user.go）                     |
| conf/       | 配置相关代码（如 conf.proto、conf.pb.go）                    |
| data/       | 数据访问层，数据库操作等（如 data.go、user.go、greeter.go）  |
| server/     | 服务启动与注册（如 grpc.go、http.go、server.go）             |
| service/    | 服务实现层，具体业务服务（如 greeter.go、user.go）           |

---

## 7. third_party/

- 存放第三方 proto 文件，便于 proto 依赖管理。
- 结构示例：

```
third_party/
  errors/
    errors.proto
  google/
    api/
      annotations.proto
      ...
    protobuf/
      any.proto
      ...
  openapi/
    v3/
      annotations.proto
      openapi.proto
  validate/
    validate.proto
    README.md
```

---

## 8. 其他重要文件

- **Dockerfile**：用于容器化部署。
- **Makefile**：自动化构建、测试、部署等任务。
- **openapi.yaml**：API 文档，便于前后端协作。
- **README.md**：项目简介、安装、使用说明等。

---

# 目录结构可视化

```plaintext
IOT_Hummingbird_back_end/
├── api/
├── bin/
├── cmd/
├── configs/
├── Dockerfile
├── docs/
├── go.mod
├── go.sum
├── internal/
├── LICENSE
├── Makefile
├── openapi.yaml
├── README.md
└── third_party/
```

---

# 总结

本项目采用主流的 Go 项目分层结构，便于团队协作、扩展和维护。各目录职责分明，建议在开发过程中严格遵循此结构，保持代码整洁有序。

如需进一步细化每个目录下的文件作用，可补充说明。 