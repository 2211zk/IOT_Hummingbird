# 登录日志管理功能使用说明

## 功能概述

登录日志管理功能是gin-vue-admin系统的一个重要安全监控模块，用于记录、查看和管理所有用户的登录活动。

## 主要功能

### 1. 自动日志记录
- 自动记录所有用户的登录尝试（成功和失败）
- 记录登录IP地址、地理位置、浏览器、操作系统等详细信息
- 支持登录失败原因记录

### 2. 日志查询和筛选
- 支持按用户名、IP地址、登录地点、登录状态筛选
- 支持时间范围查询
- 分页显示，支持自定义每页显示数量

### 3. 详细信息查看
- 点击任意日志记录可查看详细信息
- 包含完整的登录环境信息

### 4. 数据导出
- 支持Excel格式导出
- 可导出当前筛选条件下的所有记录
- 支持导出登录统计信息

### 5. 统计分析
- 显示登录成功率、总登录次数等统计信息
- 热门登录IP分析
- 最近登录记录展示

### 6. 日志管理
- 支持清理过期日志
- 可配置日志保留策略
- 清理前自动备份功能

## 使用方法

### 访问登录日志
1. 登录gin-vue-admin管理后台
2. 在左侧菜单中找到"系统管理"
3. 点击"登录日志"进入日志管理页面

### 查询日志
1. 在搜索框中输入筛选条件：
   - 用户名：输入要查询的用户名
   - 登录IP：输入IP地址
   - 登录地点：输入地理位置关键词
   - 登录状态：选择"成功"或"失败"
   - 时间范围：选择开始和结束时间
2. 点击"查询"按钮执行搜索
3. 点击"重置"按钮清空搜索条件

### 查看详情
- 点击表格中的任意一行即可查看该登录记录的详细信息

### 导出数据
1. 设置好筛选条件（可选）
2. 点击"导出Excel"按钮
3. 系统会自动下载Excel文件

### 查看统计信息
1. 点击"统计信息"按钮
2. 查看登录统计数据、热门IP和最近登录记录

### 清理日志
1. 点击"清理日志"按钮
2. 设置要保留的天数（默认90天）
3. 确认清理操作
4. 系统会自动备份要删除的数据

## API接口

### 主要接口列表
- `POST /loginLog/getLoginLogList` - 获取登录日志列表
- `GET /loginLog/findLoginLog` - 获取登录日志详情
- `POST /loginLog/exportLoginLog` - 导出登录日志
- `GET /loginLog/getLoginStatistics` - 获取登录统计
- `POST /loginLog/cleanExpiredLogs` - 清理过期日志

### 权限要求
- 需要管理员权限才能访问登录日志功能
- 相关API已自动配置到超级管理员角色

## 数据库表结构

### 登录日志表 (wy_login_log)
```sql
CREATE TABLE `wy_login_log` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `access_number` int DEFAULT NULL COMMENT '访问编号',
  `user_name` char(11) DEFAULT NULL COMMENT '用户名称',
  `login_address` varchar(100) DEFAULT NULL COMMENT '登录地址',
  `login_location` varchar(100) DEFAULT NULL COMMENT '登录地点',
  `browser` varchar(100) DEFAULT NULL COMMENT '浏览器',
  `operating_system` varchar(100) DEFAULT NULL COMMENT '操作系统',
  `login_status` varchar(10) DEFAULT NULL COMMENT '登录状态',
  `operational_information` varchar(100) DEFAULT NULL COMMENT '操作信息',
  `login_time` datetime(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6) COMMENT '登录时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='登录日志表';
```

## 安全考虑

1. **数据脱敏**：IP地址等敏感信息已进行适当处理
2. **权限控制**：只有管理员可以查看登录日志
3. **数据保护**：支持定期清理和备份功能
4. **审计跟踪**：所有管理操作都有审计记录

## 故障排除

### 常见问题

1. **菜单不显示**
   - 检查用户是否有管理员权限
   - 确认菜单已正确添加到数据库

2. **数据不显示**
   - 检查数据库表是否存在
   - 确认后端服务是否正常运行

3. **导出功能异常**
   - 检查服务器磁盘空间
   - 确认Excel导出依赖是否正常

### 日志查看
- 后端日志：查看gin-vue-admin的日志文件
- 前端错误：打开浏览器开发者工具查看控制台

## 技术实现

### 后端技术栈
- Go + Gin框架
- GORM数据库ORM
- Excel导出使用excelize库

### 前端技术栈
- Vue 3 + Composition API
- Element Plus UI组件库
- Axios HTTP客户端

### 主要文件结构
```
server/
├── api/v1/system/sys_login_log.go          # API控制器
├── service/system/sys_login_log.go         # 业务逻辑层
├── model/system/sys_login_log.go           # 数据模型
├── router/system/sys_login_log.go          # 路由配置
├── middleware/login_log.go                 # 登录日志中间件
└── utils/excel_export.go                   # Excel导出工具

web/
├── src/api/loginLog.js                     # 前端API调用
├── src/view/superAdmin/loginLog/loginLog.vue # 主页面组件
└── src/pathInfo.json                       # 路由路径配置
```

## 更新日志

### v1.0.0 (2025-01-25)
- 初始版本发布
- 支持基本的登录日志记录和查看功能
- 支持数据导出和统计分析
- 支持日志清理和管理功能