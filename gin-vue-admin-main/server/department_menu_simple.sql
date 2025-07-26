-- 部门管理菜单添加脚本（简化版）
-- 数据库：wl_playform

-- 获取系统管理菜单ID
SET @system_menu_id = (SELECT id FROM sys_base_menus WHERE name = 'superAdmin' LIMIT 1);

-- 添加部门管理菜单
INSERT IGNORE INTO sys_base_menus (created_at, updated_at, menu_level, hidden, parent_id, path, name, component, sort, meta) 
VALUES (NOW(), NOW(), 1, 0, @system_menu_id, 'departmentManage', 'departmentManage', 'view/system/department/index.vue', 9, '{"title":"部门管理","icon":"office-building","keepAlive":true}');

-- 获取部门管理菜单ID
SET @dept_menu_id = (SELECT id FROM sys_base_menus WHERE name = 'departmentManage' LIMIT 1);

-- 添加按钮权限
INSERT IGNORE INTO sys_base_menu_btns (created_at, updated_at, name, `desc`, sys_base_menu_id) VALUES
(NOW(), NOW(), 'add', '新增', @dept_menu_id),
(NOW(), NOW(), 'edit', '编辑', @dept_menu_id),
(NOW(), NOW(), 'delete', '删除', @dept_menu_id),
(NOW(), NOW(), 'info', '查看', @dept_menu_id),
(NOW(), NOW(), 'batchDelete', '批量删除', @dept_menu_id),
(NOW(), NOW(), 'exportTemplate', '导出模板', @dept_menu_id),
(NOW(), NOW(), 'exportExcel', '导出Excel', @dept_menu_id),
(NOW(), NOW(), 'importExcel', '导入Excel', @dept_menu_id);

-- 添加权限关联
INSERT IGNORE INTO sys_authority_menus (sys_base_menu_id, sys_authority_authority_id) 
VALUES (@dept_menu_id, 1);

-- 验证结果
SELECT m.id, m.name, m.path, m.component, m.parent_id, m.menu_level, m.sort
FROM sys_base_menus m 
WHERE m.name = 'superAdmin' OR m.name = 'departmentManage'
ORDER BY m.parent_id, m.sort;