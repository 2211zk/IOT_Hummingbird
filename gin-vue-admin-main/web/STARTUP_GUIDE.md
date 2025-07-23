# 前端项目启动指南

## 端口配置

- **前端端口**: 8080
- **后端端口**: 8888

## 问题描述

前端项目启动时遇到端口冲突错误：
```
error when starting dev server: Error: Port 8080 is already in use
```

## 解决方案

### 方案一：关闭占用8080端口的进程（推荐）

#### Windows系统：
```bash
# 查找占用8080端口的进程
netstat -ano | findstr :8080

# 根据PID关闭进程（替换<PID>为实际的进程ID）
taskkill /PID <PID> /F
```

#### Linux/Mac系统：
```bash
# 查找占用8080端口的进程
lsof -i :8080

# 根据PID关闭进程（替换<PID>为实际的进程ID）
kill -9 <PID>
```

### 方案二：使用自动端口选择

如果不想关闭占用8080端口的进程，可以让Vite自动选择可用端口：

```bash
# 使用自动端口选择
npm run serve -- --port 0
```

### 方案三：临时使用其他端口

如果8080端口被重要进程占用，可以临时使用其他端口：

```bash
# 使用8081端口
npm run serve -- --port 8081

# 使用8082端口
npm run serve -- --port 8082
```

## 启动步骤

1. **进入前端项目目录**：
   ```bash
   cd gin-vue-admin-main/web
   ```

2. **安装依赖**（如果还没有安装）：
   ```bash
   npm install
   ```

3. **启动开发服务器**：
   ```bash
   npm run dev
   # 或者
   npm run serve
   ```

4. **访问应用**：
   - 默认地址：http://localhost:8080
   - 如果使用自动端口选择，查看控制台输出的实际URL

## 配置说明

### vite.config.js 配置

```javascript
server: {
    port: 8080, // 前端端口
    strictPort: false, // 允许自动选择可用端口
    host: true,
    open: true,
    proxy: {
        // 代理到后端8888端口
        [process.env.VITE_BASE_API]: {
            target: `${process.env.VITE_BASE_PATH}:${process.env.VITE_SERVER_PORT}/`,
            changeOrigin: true,
            rewrite: (path) =>
                path.replace(new RegExp('^' + process.env.VITE_BASE_API), '')
        }
    }
}
```

## 常见问题

### Q: 为什么会出现端口冲突？
A: 端口冲突通常是因为：
- 之前启动的开发服务器没有正确关闭
- 其他应用程序正在使用该端口
- 多个开发项目同时运行

### Q: 如何检查端口是否被占用？
A: 使用以下命令检查：

**Windows:**
```bash
netstat -ano | findstr :8080
```

**Linux/Mac:**
```bash
lsof -i :8080
```

### Q: 如何强制关闭占用端口的进程？
A: 使用以下命令：

**Windows:**
```bash
# 查找进程
netstat -ano | findstr :8080
# 关闭进程（替换<PID>）
taskkill /PID <PID> /F
```

**Linux/Mac:**
```bash
# 查找进程
lsof -i :8080
# 关闭进程（替换<PID>）
kill -9 <PID>
```

### Q: 如何确保前后端端口配置正确？
A: 检查以下配置：

1. **前端端口**: vite.config.js 中的 `port: 8080`
2. **后端端口**: 后端配置文件中的端口设置
3. **代理配置**: 确保前端代理正确指向后端8888端口

## 推荐做法

1. **优先使用方案一**：关闭占用8080端口的进程，保持原始端口配置
2. **如果无法关闭占用进程**：使用 `npm run serve -- --port 0` 让Vite自动选择端口
3. **临时解决方案**：使用其他端口，但记得更新相关配置

## 验证启动成功

启动成功后，你应该看到类似以下的输出：

```
  VITE v6.2.3  ready in 1234 ms

  ➜  Local:   http://localhost:8080/
  ➜  Network: http://192.168.1.100:8080/
  ➜  press h to show help
```

然后可以在浏览器中访问 http://localhost:8080 来查看应用。

## 完整启动流程

1. **启动后端服务器**（端口8888）
2. **启动前端服务器**（端口8080）
3. **访问前端应用**: http://localhost:8080
4. **验证前后端通信**: 检查API请求是否正常代理到后端 