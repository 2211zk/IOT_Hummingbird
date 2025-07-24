-- 为当前用户添加wl_playform相关页面的权限
-- 这个脚本需要根据实际的菜单ID来执行

-- 首先查看现有的菜单
SELECT id, name, path, component FROM sys_base_menus WHERE name LIKE '%wl%' OR path LIKE '%wl%';

-- 查看当前用户的权限
SELECT authority_id, authority_name FROM sys_authorities;

-- 为角色888（超级管理员）添加wl_playform相关菜单权限
-- 注意：需要根据实际的菜单ID来执行

-- 示例：为角色888添加wl_playform菜单权限
-- INSERT INTO sys_authority_menus (sys_authority_authority_id, sys_base_menu_id) 
-- SELECT 888, id FROM sys_base_menus WHERE name IN ('wl_playform', 'wlProducts', 'wlEquipment', 'wlEngineRules');

-- 为角色9528（测试角色）添加wl_playform相关菜单权限
-- INSERT INTO sys_authority_menus (sys_authority_authority_id, sys_base_menu_id) 
-- SELECT 9528, id FROM sys_base_menus WHERE name IN ('wl_playform', 'wlProducts', 'wlEquipment', 'wlEngineRules');

-- 查看当前用户的菜单权限
SELECT 
    sa.authority_name,
    sbm.name as menu_name,
    sbm.path as menu_path
FROM sys_authority_menus sam
JOIN sys_authorities sa ON sam.sys_authority_authority_id = sa.authority_id
JOIN sys_base_menus sbm ON sam.sys_base_menu_id = sbm.id
WHERE sa.authority_id IN (888, 9528, 8881)
ORDER BY sa.authority_id, sbm.sort; 