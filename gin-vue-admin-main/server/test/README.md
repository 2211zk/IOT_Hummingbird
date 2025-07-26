# 部门管理模块测试文档

## 概述

本目录包含部门管理模块的完整测试套件，包括单元测试、集成测试和性能测试。

## 测试结构

```
test/
├── README.md                           # 测试文档
├── run_tests.go                        # 测试运行器
├── unit/                               # 单元测试
│   ├── wl_department_tree_test.go      # 树形结构测试
│   └── wl_department_device_test.go    # 设备关联测试
├── integration/                        # 集成测试
│   └── wl_department_integration_test.go
└── coverage/                           # 覆盖率报告输出目录
    ├── coverage.out
    └── coverage.html
```

## 测试套件说明

### 1. Service层单元测试
- **文件**: `service/wl_department/wl_department_test.go`
- **功能**: 测试Service层的业务逻辑
- **覆盖内容**:
  - 部门CRUD操作
  - 业务规则验证
  - 数据验证
  - 错误处理

### 2. API层单元测试
- **文件**: `api/v1/wl_department/wl_department_test.go`
- **功能**: 测试API接口层
- **覆盖内容**:
  - HTTP接口测试
  - 请求参数验证
  - 响应格式验证
  - 错误响应测试

### 3. 树形结构单元测试
- **文件**: `test/unit/wl_department_tree_test.go`
- **功能**: 测试部门树形结构相关功能
- **覆盖内容**:
  - 树形结构构建
  - 循环引用检测
  - 层级关系验证
  - 性能测试

### 4. 设备关联单元测试
- **文件**: `test/unit/wl_department_device_test.go`
- **功能**: 测试设备关联功能
- **覆盖内容**:
  - 设备关联操作
  - 设备查询功能
  - 关联验证
  - 并发操作测试

### 5. 集成测试
- **文件**: `test/integration/wl_department_integration_test.go`
- **功能**: 测试完整的业务流程
- **覆盖内容**:
  - 端到端业务流程
  - 多模块协作
  - 数据一致性
  - 错误场景处理

## 运行测试

### 使用测试运行器（推荐）

```bash
# 进入server目录
cd gin-vue-admin-main/server

# 运行所有测试
go run test/run_tests.go all

# 运行测试并生成覆盖率报告
go run test/run_tests.go coverage

# 运行性能测试
go run test/run_tests.go bench

# 运行特定测试套件
go run test/run_tests.go test service
go run test/run_tests.go test api
go run test/run_tests.go test tree
go run test/run_tests.go test device
go run test/run_tests.go test integration
```

### 直接运行测试

```bash
# 运行Service层测试
cd service/wl_department
go test -v

# 运行API层测试
cd api/v1/wl_department
go test -v

# 运行单元测试
cd test/unit
go test -v

# 运行集成测试
cd test/integration
go test -v

# 运行所有测试
go test -v ./...

# 运行测试并生成覆盖率
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out -o coverage.html
```

## 测试覆盖范围

### 功能测试覆盖

- ✅ 部门创建、更新、删除
- ✅ 部门列表查询（平铺和树形）
- ✅ 部门树构建和验证
- ✅ 循环引用检测
- ✅ 设备关联管理
- ✅ 状态管理和继承
- ✅ 排序功能
- ✅ 权限验证
- ✅ 错误处理

### 边界条件测试

- ✅ 空数据处理
- ✅ 无效参数处理
- ✅ 深层嵌套结构
- ✅ 大量数据处理
- ✅ 并发操作
- ✅ 数据库约束

### 性能测试

- ✅ 树形结构构建性能
- ✅ 循环引用检测性能
- ✅ 大量数据查询性能
- ✅ 并发操作性能

## 测试数据

### 模拟数据结构

```go
// 部门层级结构示例
// 1 (公司)
//   ├── 2 (技术部)
//   │   ├── 4 (前端组)
//   │   │   ├── 7 (React组)
//   │   │   └── 8 (Vue组)
//   │   └── 5 (后端组)
//   │       └── 9 (Go组)
//   └── 3 (市场部)
//       └── 6 (销售组)
```

### 设备数据

```go
// 设备关联示例
devices := []WlDevice{
    {ID: 1, DeviceName: "温度传感器1", ProductName: "温度传感器产品A"},
    {ID: 2, DeviceName: "湿度传感器1", ProductName: "湿度传感器产品B"},
    {ID: 3, DeviceName: "压力传感器1", ProductName: "压力传感器产品C"},
}
```

## 测试最佳实践

### 1. 测试命名规范

```go
// 测试函数命名：Test + 功能名称
func TestCreateWlDepartment(t *testing.T) {}

// 测试用例命名：描述性的中文名称
suite.Run("创建顶级部门成功", func() {})
```

### 2. 测试结构

```go
// 使用表驱动测试
tests := []struct {
    name        string
    input       InputType
    expected    ExpectedType
    expectError bool
}{
    {
        name:        "正常情况",
        input:       validInput,
        expected:    expectedOutput,
        expectError: false,
    },
}
```

### 3. 断言使用

```go
// 使用testify断言库
suite.Equal(expected, actual, "错误消息")
suite.True(condition, "条件应该为真")
suite.NotEmpty(slice, "切片不应该为空")
```

### 4. 模拟和存根

```go
// 使用模拟函数替代实际的数据库操作
func simulateCreateDepartment(req CreateRequest) CreateResult {
    // 模拟业务逻辑
    return CreateResult{Success: true}
}
```

## 持续集成

### GitHub Actions配置示例

```yaml
name: 部门管理模块测试

on: [push, pull_request]

jobs:
  test:
    runs-on: ubuntu-latest
    steps:
    - uses: actions/checkout@v2
    - name: Set up Go
      uses: actions/setup-go@v2
      with:
        go-version: 1.19
    - name: Run tests
      run: |
        cd gin-vue-admin-main/server
        go run test/run_tests.go all
    - name: Generate coverage
      run: |
        cd gin-vue-admin-main/server
        go run test/run_tests.go coverage
```

## 故障排除

### 常见问题

1. **导入路径错误**
   ```bash
   # 确保在正确的目录下运行测试
   cd gin-vue-admin-main/server
   ```

2. **依赖包缺失**
   ```bash
   go mod tidy
   go mod download
   ```

3. **测试超时**
   ```bash
   go test -timeout 30s
   ```

### 调试技巧

1. **详细输出**
   ```bash
   go test -v -run TestSpecificFunction
   ```

2. **只运行失败的测试**
   ```bash
   go test -v -run TestFailed
   ```

3. **生成测试报告**
   ```bash
   go test -json > test_results.json
   ```

## 贡献指南

### 添加新测试

1. 确定测试类型（单元测试/集成测试）
2. 选择合适的测试文件或创建新文件
3. 遵循现有的测试结构和命名规范
4. 添加必要的文档说明
5. 更新测试运行器配置

### 测试代码审查清单

- [ ] 测试覆盖了所有主要功能路径
- [ ] 包含边界条件和错误情况测试
- [ ] 测试名称清晰描述测试目的
- [ ] 使用适当的断言和错误消息
- [ ] 测试数据合理且有代表性
- [ ] 性能测试包含基准测试
- [ ] 文档更新完整

## 参考资料

- [Go Testing Package](https://golang.org/pkg/testing/)
- [Testify Documentation](https://github.com/stretchr/testify)
- [Go Test Coverage](https://golang.org/doc/tutorial/add-a-test)
- [Testing Best Practices](https://golang.org/doc/tutorial/add-a-test)