-- 为所有角色添加wl_playform相关页面的权限
-- 执行这个脚本后需要重启后端服务

-- 1. 首先查看现有的wl_playform相关菜单
SELECT id, name, path, component, parent_id FROM sys_base_menus WHERE name LIKE '%wl%' OR path LIKE '%wl%';

-- 2. 查看现有的角色
SELECT authority_id, authority_name FROM sys_authorities;

-- 3. 为角色888（超级管理员）添加wl_playform相关菜单权限
INSERT INTO sys_authority_menus (sys_authority_authority_id, sys_base_menu_id) 
SELECT 888, id FROM sys_base_menus WHERE name IN ('wl_playform', 'wlProducts', 'wlEquipment', 'wlEngineRules', 'wlResources', 'wlScenes', 'wlCategory', 'wlCaFunction')
ON DUPLICATE KEY UPDATE sys_authority_authority_id = sys_authority_authority_id;

-- 4. 为角色9528（测试角色）添加wl_playform相关菜单权限
INSERT INTO sys_authority_menus (sys_authority_authority_id, sys_base_menu_id) 
SELECT 9528, id FROM sys_base_menus WHERE name IN ('wl_playform', 'wlProducts', 'wlEquipment', 'wlEngineRules', 'wlResources', 'wlScenes', 'wlCategory', 'wlCaFunction')
ON DUPLICATE KEY UPDATE sys_authority_authority_id = sys_authority_authority_id;

-- 5. 为角色8881（普通用户子角色）添加wl_playform相关菜单权限
INSERT INTO sys_authority_menus (sys_authority_authority_id, sys_base_menu_id) 
SELECT 8881, id FROM sys_base_menus WHERE name IN ('wl_playform', 'wlProducts', 'wlEquipment', 'wlEngineRules', 'wlResources', 'wlScenes', 'wlCategory', 'wlCaFunction')
ON DUPLICATE KEY UPDATE sys_authority_authority_id = sys_authority_authority_id;

-- 6. 验证权限是否添加成功
SELECT 
    sa.authority_name,
    sbm.name as menu_name,
    sbm.path as menu_path
FROM sys_authority_menus sam
JOIN sys_authorities sa ON sam.sys_authority_authority_id = sa.authority_id
JOIN sys_base_menus sbm ON sam.sys_base_menu_id = sbm.id
WHERE sa.authority_id IN (888, 9528, 8881) AND sbm.name LIKE '%wl%'
ORDER BY sa.authority_id, sbm.sort; 