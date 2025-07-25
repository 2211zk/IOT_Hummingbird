-- 给角色888授予所有权限
-- 1. 获取所有菜单ID并给888权限
INSERT IGNORE INTO sys_authority_menus (sys_base_menu_id, sys_authority_authority_id)
SELECT id, 888 FROM sys_base_menus WHERE id > 0;

-- 2. 获取所有API权限并给888权限
-- 先获取所有现有的API权限
INSERT IGNORE INTO casbin_rule (ptype, v0, v1, v2)
SELECT DISTINCT 'p', '888', v1, v2 FROM casbin_rule WHERE ptype = 'p' AND v0 != '888';

-- 3. 确保dashboard相关权限
INSERT IGNORE INTO casbin_rule (ptype, v0, v1, v2) VALUES 
('p', '888', '/dashboard/getDashboardData', 'GET'),
('p', '888', '/dashboard/*', 'GET'),
('p', '888', '/dashboard/*', 'POST'),
('p', '888', '/dashboard/*', 'PUT'),
('p', '888', '/dashboard/*', 'DELETE');

-- 4. 确保wl_playform相关权限
INSERT IGNORE INTO casbin_rule (ptype, v0, v1, v2) VALUES 
('p', '888', '/wl_playform/*', 'GET'),
('p', '888', '/wl_playform/*', 'POST'),
('p', '888', '/wl_playform/*', 'PUT'),
('p', '888', '/wl_playform/*', 'DELETE'),
('p', '888', '/wl_playform/wlProducts/*', 'GET'),
('p', '888', '/wl_playform/wlProducts/*', 'POST'),
('p', '888', '/wl_playform/wlProducts/*', 'PUT'),
('p', '888', '/wl_playform/wlProducts/*', 'DELETE'),
('p', '888', '/wl_playform/wlEquipment/*', 'GET'),
('p', '888', '/wl_playform/wlEquipment/*', 'POST'),
('p', '888', '/wl_playform/wlEquipment/*', 'PUT'),
('p', '888', '/wl_playform/wlEquipment/*', 'DELETE'),
('p', '888', '/wl_playform/wlResources/*', 'GET'),
('p', '888', '/wl_playform/wlResources/*', 'POST'),
('p', '888', '/wl_playform/wlResources/*', 'PUT'),
('p', '888', '/wl_playform/wlResources/*', 'DELETE'),
('p', '888', '/wl_playform/wlEngineRules/*', 'GET'),
('p', '888', '/wl_playform/wlEngineRules/*', 'POST'),
('p', '888', '/wl_playform/wlEngineRules/*', 'PUT'),
('p', '888', '/wl_playform/wlEngineRules/*', 'DELETE');

-- 5. 确保系统管理相关权限
INSERT IGNORE INTO casbin_rule (ptype, v0, v1, v2) VALUES 
('p', '888', '/system/*', 'GET'),
('p', '888', '/system/*', 'POST'),
('p', '888', '/system/*', 'PUT'),
('p', '888', '/system/*', 'DELETE'),
('p', '888', '/user/*', 'GET'),
('p', '888', '/user/*', 'POST'),
('p', '888', '/user/*', 'PUT'),
('p', '888', '/user/*', 'DELETE'),
('p', '888', '/authority/*', 'GET'),
('p', '888', '/authority/*', 'POST'),
('p', '888', '/authority/*', 'PUT'),
('p', '888', '/authority/*', 'DELETE'),
('p', '888', '/menu/*', 'GET'),
('p', '888', '/menu/*', 'POST'),
('p', '888', '/menu/*', 'PUT'),
('p', '888', '/menu/*', 'DELETE'),
('p', '888', '/api/*', 'GET'),
('p', '888', '/api/*', 'POST'),
('p', '888', '/api/*', 'PUT'),
('p', '888', '/api/*', 'DELETE');

-- 6. 确保设备管理相关权限
INSERT IGNORE INTO casbin_rule (ptype, v0, v1, v2) VALUES 
('p', '888', '/wl_driver/*', 'GET'),
('p', '888', '/wl_driver/*', 'POST'),
('p', '888', '/wl_driver/*', 'PUT'),
('p', '888', '/wl_driver/*', 'DELETE'),
('p', '888', '/driver/*', 'GET'),
('p', '888', '/driver/*', 'POST'),
('p', '888', '/driver/*', 'PUT'),
('p', '888', '/driver/*', 'DELETE');

-- 7. 确保告警中心相关权限
INSERT IGNORE INTO casbin_rule (ptype, v0, v1, v2) VALUES 
('p', '888', '/alarm/*', 'GET'),
('p', '888', '/alarm/*', 'POST'),
('p', '888', '/alarm/*', 'PUT'),
('p', '888', '/alarm/*', 'DELETE');

-- 8. 确保系统工具相关权限
INSERT IGNORE INTO casbin_rule (ptype, v0, v1, v2) VALUES 
('p', '888', '/sysTools/*', 'GET'),
('p', '888', '/sysTools/*', 'POST'),
('p', '888', '/sysTools/*', 'PUT'),
('p', '888', '/sysTools/*', 'DELETE'),
('p', '888', '/state/*', 'GET'),
('p', '888', '/state/*', 'POST'),
('p', '888', '/state/*', 'PUT'),
('p', '888', '/state/*', 'DELETE');

-- 9. 确保文件上传相关权限
INSERT IGNORE INTO casbin_rule (ptype, v0, v1, v2) VALUES 
('p', '888', '/fileUploadAndDownload/*', 'GET'),
('p', '888', '/fileUploadAndDownload/*', 'POST'),
('p', '888', '/fileUploadAndDownload/*', 'PUT'),
('p', '888', '/fileUploadAndDownload/*', 'DELETE'),
('p', '888', '/upload/*', 'GET'),
('p', '888', '/upload/*', 'POST'),
('p', '888', '/upload/*', 'PUT'),
('p', '888', '/upload/*', 'DELETE');

-- 10. 确保其他通用权限
INSERT IGNORE INTO casbin_rule (ptype, v0, v1, v2) VALUES 
('p', '888', '/jwt/*', 'GET'),
('p', '888', '/jwt/*', 'POST'),
('p', '888', '/jwt/*', 'PUT'),
('p', '888', '/jwt/*', 'DELETE'),
('p', '888', '/captcha/*', 'GET'),
('p', '888', '/captcha/*', 'POST'),
('p', '888', '/captcha/*', 'PUT'),
('p', '888', '/captcha/*', 'DELETE'),
('p', '888', '/email/*', 'GET'),
('p', '888', '/email/*', 'POST'),
('p', '888', '/email/*', 'PUT'),
('p', '888', '/email/*', 'DELETE');

-- 11. 检查权限设置结果
SELECT '角色888菜单权限数量:' as info, COUNT(*) as count FROM sys_authority_menus WHERE sys_authority_authority_id = 888;

SELECT '角色888 API权限数量:' as info, COUNT(*) as count FROM casbin_rule WHERE v0 = '888';

SELECT '所有菜单权限检查:' as info;
SELECT m.menu_name, am.sys_authority_authority_id 
FROM sys_authority_menus am 
JOIN sys_base_menus m ON am.sys_base_menu_id = m.id 
WHERE am.sys_authority_authority_id = 888 
ORDER BY m.id; 