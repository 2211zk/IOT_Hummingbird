# 🔧 权限修复执行指南

## 📋 问题描述
用户在删除登录日志时提示"权限不足"，需要为角色888添加相应的API权限。

## 🚀 立即执行步骤

### 方法1：使用数据库管理工具（推荐）

1. **打开你的数据库管理工具**（如phpMyAdmin、Navicat、DBeaver等）

2. **连接到数据库**：
   - 主机：`14.103.143.229`
   - 端口：`3306`
   - 用户名：`root`
   - 密码：`82AF916059F13E331E633AC0B8A8191B`
   - 数据库：`wl_playform`

3. **执行SQL脚本**：
   - 打开文件：`gin-vue-admin-main/server/MANUAL_PERMISSION_FIX.sql`
   - 复制所有内容
   - 粘贴到数据库管理工具的SQL执行窗口
   - 点击执行

### 方法2：使用命令行（如果有MySQL客户端）

```bash
# 进入服务器目录
cd gin-vue-admin-main/server

# 执行权限修复脚本
mysql -h 14.103.143.229 -P 3306 -u root -p82AF916059F13E331E633AC0B8A8191B wl_playform < MANUAL_PERMISSION_FIX.sql
```

### 方法3：逐步手动执行

如果上述方法都不可用，请按以下顺序逐步执行SQL语句：

#### 步骤1：添加API权限
```sql
INSERT IGNORE INTO `sys_apis` (`path`, `description`, `api_group`, `method`, `created_at`, `updated_at`) VALUES
('/loginLog/deleteLoginLog', '删除登录日志', '登录日志', 'DELETE', NOW(), NOW()),
('/loginLog/deleteLoginLogByIds', '批量删除登录日志', '登录日志', 'DELETE', NOW(), NOW());
```

#### 步骤2：为角色888添加权限
```sql
INSERT INTO `sys_authority_apis` (`sys_authority_authority_id`, `sys_api_id`)
SELECT 888, `id` FROM `sys_apis` WHERE `api_group` = '登录日志';
```

#### 步骤3：添加Casbin规则
```sql
INSERT INTO `casbin_rule` (`ptype`, `v0`, `v1`, `v2`) VALUES
('p', '888', '/loginLog/deleteLoginLog', 'DELETE'),
('p', '888', '/loginLog/deleteLoginLogByIds', 'DELETE');
```

## ✅ 验证修复结果

执行以下SQL查询来验证权限是否添加成功：

```sql
-- 检查API权限
SELECT a.path, a.method, a.description
FROM sys_apis a
JOIN sys_authority_apis aa ON a.id = aa.sys_api_id
WHERE aa.sys_authority_authority_id = 888 AND a.api_group = '登录日志'
ORDER BY a.path;

-- 检查Casbin规则
SELECT v1 as api_path, v2 as method
FROM casbin_rule 
WHERE ptype = 'p' AND v0 = '888' AND v1 LIKE '%loginLog%'
ORDER BY v1;
```

## 🔄 完成修复后的步骤

1. **重启后端服务**：
   ```bash
   cd gin-vue-admin-main/server
   go run main.go
   ```

2. **清除浏览器缓存**：
   - 按 `Ctrl+Shift+Delete`
   - 清除缓存和Cookie

3. **重新登录系统**：
   - 退出当前登录
   - 重新登录

4. **测试删除功能**：
   - 进入登录日志页面
   - 选择记录并点击删除
   - 确认功能正常工作

## 🔍 故障排除

### 如果仍然提示权限不足：

1. **检查用户角色**：
   ```sql
   SELECT id, username, authority_id FROM sys_users WHERE username = 'your_username';
   ```

2. **检查角色是否为888**：
   - 如果不是888，请将上述SQL中的888替换为实际的角色ID

3. **检查API是否存在**：
   ```sql
   SELECT * FROM sys_apis WHERE path = '/loginLog/deleteLoginLogByIds';
   ```

4. **重新加载权限**：
   - 重启后端服务
   - 清除前端缓存
   - 重新登录

### 如果SQL执行出错：

1. **检查表是否存在**：
   ```sql
   SHOW TABLES LIKE 'sys_%';
   ```

2. **检查字段是否正确**：
   ```sql
   DESCRIBE sys_apis;
   DESCRIBE sys_authority_apis;
   ```

## 📞 需要帮助？

如果按照以上步骤操作后仍有问题，请提供：

1. 执行SQL后的结果截图
2. 浏览器控制台的错误信息
3. 后端服务的日志输出
4. 用户的角色ID信息

---

**重要提示**：执行SQL脚本前请备份数据库！