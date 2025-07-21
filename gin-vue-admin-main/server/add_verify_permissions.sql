-- 添加wlResources验证按钮权限

-- 1. 查找wlResources菜单ID
SELECT 'wlResources菜单ID:' as info, id, name FROM sys_base_menus WHERE name = 'wlResources';

-- 2. 查找wlResources菜单的所有按钮
SELECT 'wlResources菜单按钮:' as info, id, name, `desc`, sys_base_menu_id FROM sys_base_menu_btns WHERE sys_base_menu_id = (SELECT id FROM sys_base_menus WHERE name = 'wlResources');

-- 3. 检查是否已存在角色888的wlResources按钮权限
SELECT '现有角色888的wlResources按钮权限:' as info, id, authority_id, sys_menu_id FROM sys_authority_btns WHERE authority_id = 888 AND sys_menu_id = (SELECT id FROM sys_base_menus WHERE name = 'wlResources');

-- 4. 为角色888添加wlResources的所有按钮权限
INSERT INTO sys_authority_btns (authority_id, sys_menu_id)
SELECT 888, sys_base_menu_id FROM sys_base_menu_btns 
WHERE sys_base_menu_id = (SELECT id FROM sys_base_menus WHERE name = 'wlResources')
AND NOT EXISTS (
    SELECT 1 FROM sys_authority_btns 
    WHERE authority_id = 888 
    AND sys_menu_id = sys_base_menu_btns.sys_base_menu_id
);

-- 5. 为角色8881添加wlResources的所有按钮权限
INSERT INTO sys_authority_btns (authority_id, sys_menu_id)
SELECT 8881, sys_base_menu_id FROM sys_base_menu_btns 
WHERE sys_base_menu_id = (SELECT id FROM sys_base_menus WHERE name = 'wlResources')
AND NOT EXISTS (
    SELECT 1 FROM sys_authority_btns 
    WHERE authority_id = 8881 
    AND sys_menu_id = sys_base_menu_btns.sys_base_menu_id
);

-- 6. 为角色9528添加wlResources的所有按钮权限
INSERT INTO sys_authority_btns (authority_id, sys_menu_id)
SELECT 9528, sys_base_menu_id FROM sys_base_menu_btns 
WHERE sys_base_menu_id = (SELECT id FROM sys_base_menus WHERE name = 'wlResources')
AND NOT EXISTS (
    SELECT 1 FROM sys_authority_btns 
    WHERE authority_id = 9528 
    AND sys_menu_id = sys_base_menu_btns.sys_base_menu_id
);

-- 7. 验证添加结果
SELECT '添加后的角色888按钮权限:' as info, id, authority_id, sys_menu_id FROM sys_authority_btns WHERE authority_id = 888 AND sys_menu_id = (SELECT id FROM sys_base_menus WHERE name = 'wlResources'); 