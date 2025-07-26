-- 将部门管理添加到现有的超级管理员菜单下
-- 数据库：wl_playform

-- 1. 获取超级管理员菜单ID
SET @superadmin_id = (SELECT id FROM sys_base_menus WHERE name = 'superAdmin' LIMIT 1);

-- 2. 检查部门管理菜单是否已存在
SELECT @superadmin_id as superadmin_menu_id;

-- 3. 添加部门管理菜单（如果不存在）
INSERT IGNORE INTO sys_base_menus (
    created_at, 
    updated_at, 
    menu_level, 
    hidden, 
    parent_id, 
    path, 
    name, 
    component, 
    sort, 
    meta
) VALUES (
    NOW(), 
    NOW(), 
    1, 
    0, 
    @superadmin_id, 
    'departmentManage', 
    'departmentManage', 
    'view/system/department/index.vue', 
    10, 
    '{"title":"部门管理","icon":"office-building","keepAlive":true}'
);

-- 4. 获取部门管理菜单ID
SET @dept_menu_id = (SELECT id FROM sys_base_menus WHERE name = 'departmentManage' AND parent_id = @superadmin_id LIMIT 1);

-- 5. 为管理员角色添加权限
INSERT IGNORE INTO sys_authority_menus (sys_base_menu_id, sys_authority_authority_id) 
VALUES (@dept_menu_id, 1);

-- 6. 添加按钮权限
INSERT IGNORE INTO sys_base_menu_btns (created_at, updated_at, name, `desc`, sys_base_menu_id) VALUES
(NOW(), NOW(), 'add', '新增', @dept_menu_id),
(NOW(), NOW(), 'edit', '编辑', @dept_menu_id),
(NOW(), NOW(), 'delete', '删除', @dept_menu_id),
(NOW(), NOW(), 'info', '查看', @dept_menu_id);

-- 7. 验证结果
SELECT 
    m.id,
    m.name,
    m.path,
    m.parent_id,
    m.sort,
    m.meta,
    CASE WHEN m.parent_id = 0 THEN '父菜单' ELSE '子菜单' END as menu_type
FROM sys_base_menus m 
WHERE m.name = 'superAdmin' OR m.parent_id = @superadmin_id
ORDER BY m.parent_id, m.sort;