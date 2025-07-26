# 部门管理菜单设置指南

## 问题描述
系统管理下没有显示部门管理菜单，需要将部门管理功能添加到系统菜单中。

## 解决方案

### 方法1：使用批处理脚本（推荐）

1. 打开命令提示符，进入server目录：
   ```cmd
   cd gin-vue-admin-main\server
   ```

2. 运行批处理脚本：
   ```cmd
   add_department_menu.bat
   ```

3. 按提示输入数据库连接信息：
   - 数据库主机：localhost（默认）
   - 数据库端口：3306（默认）
   - 数据库名称：gva（默认）
   - 数据库用户名：root（默认）
   - 数据库密码：输入你的MySQL密码

### 方法2：使用PowerShell脚本

1. 打开PowerShell，进入server目录：
   ```powershell
   cd gin-vue-admin-main\server
   ```

2. 运行PowerShell脚本：
   ```powershell
   .\execute_department_menu.ps1
   ```

### 方法3：手动执行SQL

1. 连接到MySQL数据库：
   ```cmd
   mysql -u root -p gva
   ```

2. 执行SQL文件：
   ```sql
   source add_department_menu.sql;
   ```

## 脚本功能

该脚本会执行以下操作：

1. **找到现有的系统管理菜单**：
   - 查找名为"superAdmin"的现有系统管理菜单

2. **添加部门管理子菜单**：
   - 菜单名称：部门管理
   - 路径：departmentManage
   - 组件：view/system/department/index.vue
   - 图标：office-building
   - 排序：9（放在现有菜单项之后）

3. **添加按钮权限**：
   - 为部门管理添加增删改查等按钮权限

4. **分配权限**：
   - 将部门管理菜单分配给管理员角色

## 验证结果

执行完成后，脚本会显示添加的菜单信息。你应该能看到：

```
+----+------------------+------------------+----------------------------------+-----------+------------+------+
| id | name            | path             | component                        | parent_id | menu_level | sort |
+----+------------------+------------------+----------------------------------+-----------+------------+------+
| XX | superAdmin      | admin            | view/superAdmin/index.vue        |         0 |          0 |    3 |
| XX | departmentManage| departmentManage | view/system/department/index.vue |        XX |          1 |    9 |
+----+------------------+------------------+----------------------------------+-----------+------------+------+
```

## 完成后的步骤

1. **重启服务器**：
   ```cmd
   cd gin-vue-admin-main\server
   go run main.go
   ```

2. **刷新浏览器**：
   - 清除浏览器缓存
   - 重新登录系统
   - 查看左侧菜单，应该能看到"系统管理"菜单下的"部门管理"选项

## 故障排除

### 如果菜单仍然不显示：

1. **检查用户权限**：
   ```sql
   SELECT * FROM sys_authority_menus WHERE sys_base_menu_id IN (
       SELECT id FROM sys_base_menus WHERE name = 'departmentManage'
   );
   ```

2. **检查菜单状态**：
   ```sql
   SELECT * FROM sys_base_menus WHERE name = 'departmentManage';
   ```

3. **清除前端缓存**：
   - 按F12打开开发者工具
   - 右键刷新按钮，选择"清空缓存并硬性重新加载"

4. **检查前端路由配置**：
   - 确认 `gin-vue-admin-main/web/src/router/index.js` 中包含部门管理路由
   - 确认 `gin-vue-admin-main/web/src/view/system/department/index.vue` 文件存在

## 注意事项

- 执行脚本前请备份数据库
- 确保MySQL服务正在运行
- 确保有足够的数据库权限执行DDL和DML操作
- 如果数据库名称不是"gva"，请相应修改脚本中的数据库名称

## 联系支持

如果遇到问题，请检查：
1. 数据库连接是否正常
2. 用户是否有足够权限
3. 前端组件文件是否存在
4. 服务器是否正常重启