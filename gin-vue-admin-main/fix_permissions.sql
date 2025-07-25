-- 权限修复脚本 - 解决"权限不足"问题
-- 执行此脚本前请确保已备份数据库

-- 1. 确保超级管理员角色存在
INSERT IGNORE INTO sys_authorities (authority_id, authority_name, parent_id, default_router, created_at, updated_at) 
VALUES (888, '超级管理员', 0, 'dashboard', NOW(), NOW());

-- 2. 确保普通用户角色存在
INSERT IGNORE INTO sys_authorities (authority_id, authority_name, parent_id, default_router, created_at, updated_at) 
VALUES (8881, '普通用户', 888, 'dashboard', NOW(), NOW());

-- 3. 确保测试角色存在
INSERT IGNORE INTO sys_authorities (authority_id, authority_name, parent_id, default_router, created_at, updated_at) 
VALUES (9528, '测试角色', 0, 'dashboard', NOW(), NOW());

-- 4. 给所有角色分配所有菜单权限
INSERT IGNORE INTO sys_authority_menus (sys_base_menu_id, sys_authority_authority_id)
SELECT id, 888 FROM sys_base_menus WHERE id > 0;

INSERT IGNORE INTO sys_authority_menus (sys_base_menu_id, sys_authority_authority_id)
SELECT id, 8881 FROM sys_base_menus WHERE id > 0;

INSERT IGNORE INTO sys_authority_menus (sys_base_menu_id, sys_authority_authority_id)
SELECT id, 9528 FROM sys_base_menus WHERE id > 0;

-- 5. 清除现有的casbin规则（可选，谨慎操作）
-- DELETE FROM casbin_rule WHERE v0 IN ('888', '8881', '9528');

-- 6. 为所有角色添加完整的API权限
-- Dashboard相关权限
INSERT IGNORE INTO casbin_rule (ptype, v0, v1, v2) VALUES 
('p', '888', '/dashboard/getDashboardData', 'GET'),
('p', '8881', '/dashboard/getDashboardData', 'GET'),
('p', '9528', '/dashboard/getDashboardData', 'GET');

-- 用户管理相关权限
INSERT IGNORE INTO casbin_rule (ptype, v0, v1, v2) VALUES 
('p', '888', '/user/*', 'GET'),
('p', '888', '/user/*', 'POST'),
('p', '888', '/user/*', 'PUT'),
('p', '888', '/user/*', 'DELETE'),
('p', '8881', '/user/getUserInfo', 'GET'),
('p', '9528', '/user/getUserInfo', 'GET');

-- 权限管理相关权限
INSERT IGNORE INTO casbin_rule (ptype, v0, v1, v2) VALUES 
('p', '888', '/authority/*', 'GET'),
('p', '888', '/authority/*', 'POST'),
('p', '888', '/authority/*', 'PUT'),
('p', '888', '/authority/*', 'DELETE');

-- 菜单管理相关权限
INSERT IGNORE INTO casbin_rule (ptype, v0, v1, v2) VALUES 
('p', '888', '/menu/*', 'GET'),
('p', '888', '/menu/*', 'POST'),
('p', '888', '/menu/*', 'PUT'),
('p', '888', '/menu/*', 'DELETE');

-- API管理相关权限
INSERT IGNORE INTO casbin_rule (ptype, v0, v1, v2) VALUES 
('p', '888', '/api/*', 'GET'),
('p', '888', '/api/*', 'POST'),
('p', '888', '/api/*', 'PUT'),
('p', '888', '/api/*', 'DELETE');

-- 系统管理相关权限
INSERT IGNORE INTO casbin_rule (ptype, v0, v1, v2) VALUES 
('p', '888', '/system/*', 'GET'),
('p', '888', '/system/*', 'POST'),
('p', '888', '/system/*', 'PUT'),
('p', '888', '/system/*', 'DELETE');

-- wl_playform相关权限
INSERT IGNORE INTO casbin_rule (ptype, v0, v1, v2) VALUES 
('p', '888', '/wl_playform/*', 'GET'),
('p', '888', '/wl_playform/*', 'POST'),
('p', '888', '/wl_playform/*', 'PUT'),
('p', '888', '/wl_playform/*', 'DELETE'),
('p', '8881', '/wl_playform/*', 'GET'),
('p', '8881', '/wl_playform/*', 'POST'),
('p', '8881', '/wl_playform/*', 'PUT'),
('p', '8881', '/wl_playform/*', 'DELETE'),
('p', '9528', '/wl_playform/*', 'GET'),
('p', '9528', '/wl_playform/*', 'POST'),
('p', '9528', '/wl_playform/*', 'PUT'),
('p', '9528', '/wl_playform/*', 'DELETE');

-- wlProducts相关权限
INSERT IGNORE INTO casbin_rule (ptype, v0, v1, v2) VALUES 
('p', '888', '/wlProducts/*', 'GET'),
('p', '888', '/wlProducts/*', 'POST'),
('p', '888', '/wlProducts/*', 'PUT'),
('p', '888', '/wlProducts/*', 'DELETE'),
('p', '8881', '/wlProducts/*', 'GET'),
('p', '8881', '/wlProducts/*', 'POST'),
('p', '8881', '/wlProducts/*', 'PUT'),
('p', '8881', '/wlProducts/*', 'DELETE'),
('p', '9528', '/wlProducts/*', 'GET'),
('p', '9528', '/wlProducts/*', 'POST'),
('p', '9528', '/wlProducts/*', 'PUT'),
('p', '9528', '/wlProducts/*', 'DELETE');

-- wlEquipment相关权限
INSERT IGNORE INTO casbin_rule (ptype, v0, v1, v2) VALUES 
('p', '888', '/wlEquipment/*', 'GET'),
('p', '888', '/wlEquipment/*', 'POST'),
('p', '888', '/wlEquipment/*', 'PUT'),
('p', '888', '/wlEquipment/*', 'DELETE'),
('p', '8881', '/wlEquipment/*', 'GET'),
('p', '8881', '/wlEquipment/*', 'POST'),
('p', '8881', '/wlEquipment/*', 'PUT'),
('p', '8881', '/wlEquipment/*', 'DELETE'),
('p', '9528', '/wlEquipment/*', 'GET'),
('p', '9528', '/wlEquipment/*', 'POST'),
('p', '9528', '/wlEquipment/*', 'PUT'),
('p', '9528', '/wlEquipment/*', 'DELETE');

-- wlDrivers相关权限
INSERT IGNORE INTO casbin_rule (ptype, v0, v1, v2) VALUES 
('p', '888', '/wlDrivers/*', 'GET'),
('p', '888', '/wlDrivers/*', 'POST'),
('p', '888', '/wlDrivers/*', 'PUT'),
('p', '888', '/wlDrivers/*', 'DELETE'),
('p', '8881', '/wlDrivers/*', 'GET'),
('p', '8881', '/wlDrivers/*', 'POST'),
('p', '8881', '/wlDrivers/*', 'PUT'),
('p', '8881', '/wlDrivers/*', 'DELETE'),
('p', '9528', '/wlDrivers/*', 'GET'),
('p', '9528', '/wlDrivers/*', 'POST'),
('p', '9528', '/wlDrivers/*', 'PUT'),
('p', '9528', '/wlDrivers/*', 'DELETE');

-- wlResources相关权限
INSERT IGNORE INTO casbin_rule (ptype, v0, v1, v2) VALUES 
('p', '888', '/wlResources/*', 'GET'),
('p', '888', '/wlResources/*', 'POST'),
('p', '888', '/wlResources/*', 'PUT'),
('p', '888', '/wlResources/*', 'DELETE'),
('p', '8881', '/wlResources/*', 'GET'),
('p', '8881', '/wlResources/*', 'POST'),
('p', '8881', '/wlResources/*', 'PUT'),
('p', '8881', '/wlResources/*', 'DELETE'),
('p', '9528', '/wlResources/*', 'GET'),
('p', '9528', '/wlResources/*', 'POST'),
('p', '9528', '/wlResources/*', 'PUT'),
('p', '9528', '/wlResources/*', 'DELETE');

-- wlScenes相关权限
INSERT IGNORE INTO casbin_rule (ptype, v0, v1, v2) VALUES 
('p', '888', '/wlScenes/*', 'GET'),
('p', '888', '/wlScenes/*', 'POST'),
('p', '888', '/wlScenes/*', 'PUT'),
('p', '888', '/wlScenes/*', 'DELETE'),
('p', '8881', '/wlScenes/*', 'GET'),
('p', '8881', '/wlScenes/*', 'POST'),
('p', '8881', '/wlScenes/*', 'PUT'),
('p', '8881', '/wlScenes/*', 'DELETE'),
('p', '9528', '/wlScenes/*', 'GET'),
('p', '9528', '/wlScenes/*', 'POST'),
('p', '9528', '/wlScenes/*', 'PUT'),
('p', '9528', '/wlScenes/*', 'DELETE');

-- wlEngineRules相关权限
INSERT IGNORE INTO casbin_rule (ptype, v0, v1, v2) VALUES 
('p', '888', '/wlEngineRules/*', 'GET'),
('p', '888', '/wlEngineRules/*', 'POST'),
('p', '888', '/wlEngineRules/*', 'PUT'),
('p', '888', '/wlEngineRules/*', 'DELETE'),
('p', '8881', '/wlEngineRules/*', 'GET'),
('p', '8881', '/wlEngineRules/*', 'POST'),
('p', '8881', '/wlEngineRules/*', 'PUT'),
('p', '8881', '/wlEngineRules/*', 'DELETE'),
('p', '9528', '/wlEngineRules/*', 'GET'),
('p', '9528', '/wlEngineRules/*', 'POST'),
('p', '9528', '/wlEngineRules/*', 'PUT'),
('p', '9528', '/wlEngineRules/*', 'DELETE');

-- 文件上传相关权限
INSERT IGNORE INTO casbin_rule (ptype, v0, v1, v2) VALUES 
('p', '888', '/fileUploadAndDownload/*', 'GET'),
('p', '888', '/fileUploadAndDownload/*', 'POST'),
('p', '888', '/fileUploadAndDownload/*', 'PUT'),
('p', '888', '/fileUploadAndDownload/*', 'DELETE'),
('p', '8881', '/fileUploadAndDownload/*', 'GET'),
('p', '8881', '/fileUploadAndDownload/*', 'POST'),
('p', '8881', '/fileUploadAndDownload/*', 'PUT'),
('p', '8881', '/fileUploadAndDownload/*', 'DELETE'),
('p', '9528', '/fileUploadAndDownload/*', 'GET'),
('p', '9528', '/fileUploadAndDownload/*', 'POST'),
('p', '9528', '/fileUploadAndDownload/*', 'PUT'),
('p', '9528', '/fileUploadAndDownload/*', 'DELETE');

-- JWT相关权限
INSERT IGNORE INTO casbin_rule (ptype, v0, v1, v2) VALUES 
('p', '888', '/jwt/*', 'GET'),
('p', '888', '/jwt/*', 'POST'),
('p', '888', '/jwt/*', 'PUT'),
('p', '888', '/jwt/*', 'DELETE'),
('p', '8881', '/jwt/*', 'GET'),
('p', '8881', '/jwt/*', 'POST'),
('p', '8881', '/jwt/*', 'PUT'),
('p', '8881', '/jwt/*', 'DELETE'),
('p', '9528', '/jwt/*', 'GET'),
('p', '9528', '/jwt/*', 'POST'),
('p', '9528', '/jwt/*', 'PUT'),
('p', '9528', '/jwt/*', 'DELETE');

-- 验证码相关权限
INSERT IGNORE INTO casbin_rule (ptype, v0, v1, v2) VALUES 
('p', '888', '/captcha/*', 'GET'),
('p', '888', '/captcha/*', 'POST'),
('p', '888', '/captcha/*', 'PUT'),
('p', '888', '/captcha/*', 'DELETE'),
('p', '8881', '/captcha/*', 'GET'),
('p', '8881', '/captcha/*', 'POST'),
('p', '8881', '/captcha/*', 'PUT'),
('p', '8881', '/captcha/*', 'DELETE'),
('p', '9528', '/captcha/*', 'GET'),
('p', '9528', '/captcha/*', 'POST'),
('p', '9528', '/captcha/*', 'PUT'),
('p', '9528', '/captcha/*', 'DELETE');

-- 7. 确保当前用户有正确的权限ID
-- 更新所有用户的权限ID为888（超级管理员）
UPDATE sys_users SET authority_id = 888 WHERE authority_id NOT IN (888, 8881, 9528);

-- 8. 检查权限设置结果
SELECT '权限修复完成！' as message;

SELECT '角色888菜单权限数量:' as info, COUNT(*) as count FROM sys_authority_menus WHERE sys_authority_authority_id = 888;
SELECT '角色888 API权限数量:' as info, COUNT(*) as count FROM casbin_rule WHERE v0 = '888';

SELECT '当前用户权限检查:' as info;
SELECT u.username, u.authority_id, a.authority_name 
FROM sys_users u 
LEFT JOIN sys_authorities a ON u.authority_id = a.authority_id 
WHERE u.deleted_at IS NULL; 