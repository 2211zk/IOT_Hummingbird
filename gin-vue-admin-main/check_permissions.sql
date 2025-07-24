-- 检查wl_playform父菜单
SELECT id, name, path, parent_id FROM sys_base_menus WHERE name = 'wl_playform';

-- 检查角色888是否有wl_playform父菜单权限
SELECT COUNT(*) as has_wl_playform_permission 
FROM sys_authority_menus sam 
JOIN sys_base_menus sbm ON sam.sys_base_menu_id = sbm.id 
WHERE sam.sys_authority_authority_id = 888 AND sbm.name = 'wl_playform';

-- 如果没有权限，添加wl_playform父菜单权限
INSERT INTO sys_authority_menus (sys_authority_authority_id, sys_base_menu_id) 
SELECT 888, id FROM sys_base_menus WHERE name = 'wl_playform'
ON DUPLICATE KEY UPDATE sys_authority_authority_id = sys_authority_authority_id;

-- 为角色9528也添加wl_playform父菜单权限
INSERT INTO sys_authority_menus (sys_authority_authority_id, sys_base_menu_id) 
SELECT 9528, id FROM sys_base_menus WHERE name = 'wl_playform'
ON DUPLICATE KEY UPDATE sys_authority_authority_id = sys_authority_authority_id;

-- 验证权限
SELECT 
    sa.authority_name,
    sbm.name as menu_name,
    sbm.path as menu_path
FROM sys_authority_menus sam
JOIN sys_authorities sa ON sam.sys_authority_authority_id = sa.authority_id
JOIN sys_base_menus sbm ON sam.sys_base_menu_id = sbm.id
WHERE sa.authority_id IN (888, 9528) AND sbm.name LIKE '%wl%'
ORDER BY sa.authority_id, sbm.sort; 