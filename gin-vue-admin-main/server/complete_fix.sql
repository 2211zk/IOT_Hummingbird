-- 完整修复登录日志功能的SQL脚本
-- 请在MySQL数据库中执行此脚本

-- 1. 创建登录日志表（如果不存在）
CREATE TABLE IF NOT EXISTS `wy_login_log` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `user_id` bigint unsigned DEFAULT NULL COMMENT '用户ID',
  `username` varchar(191) DEFAULT NULL COMMENT '用户名',
  `ip_address` varchar(191) DEFAULT NULL COMMENT 'IP地址',
  `location` varchar(191) DEFAULT NULL COMMENT '登录地点',
  `user_agent` varchar(500) DEFAULT NULL COMMENT '用户代理',
  `device_type` varchar(191) DEFAULT NULL COMMENT '设备类型',
  `browser` varchar(191) DEFAULT NULL COMMENT '浏览器',
  `os` varchar(191) DEFAULT NULL COMMENT '操作系统',
  `login_time` datetime(3) DEFAULT NULL COMMENT '登录时间',
  `status` varchar(191) DEFAULT NULL COMMENT '登录状态',
  `message` varchar(500) DEFAULT NULL COMMENT '登录消息',
  PRIMARY KEY (`id`),
  KEY `idx_wy_login_log_deleted_at` (`deleted_at`),
  KEY `idx_wy_login_log_user_id` (`user_id`),
  KEY `idx_wy_login_log_username` (`username`),
  KEY `idx_wy_login_log_login_time` (`login_time`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- 2. 添加登录日志相关的API记录
INSERT IGNORE INTO `sys_apis` (
    `created_at`, 
    `updated_at`, 
    `deleted_at`, 
    `path`, 
    `description`, 
    `api_group`, 
    `method`
) VALUES 
(NOW(), NOW(), NULL, '/loginLog/getLoginLogList', '获取登录日志列表', '登录日志', 'POST'),
(NOW(), NOW(), NULL, '/loginLog/findLoginLog', '获取登录日志详情', '登录日志', 'GET'),
(NOW(), NOW(), NULL, '/loginLog/exportLoginLog', '导出登录日志', '登录日志', 'POST'),
(NOW(), NOW(), NULL, '/loginLog/getLoginStatistics', '获取登录统计', '登录日志', 'GET'),
(NOW(), NOW(), NULL, '/loginLog/getRecentLoginLogs', '获取最近登录日志', '登录日志', 'GET'),
(NOW(), NOW(), NULL, '/loginLog/getTopLoginIPs', '获取热门登录IP', '登录日志', 'GET'),
(NOW(), NOW(), NULL, '/loginLog/cleanExpiredLogs', '清理过期日志', '登录日志', 'POST'),
(NOW(), NOW(), NULL, '/loginLog/getCleanupStatistics', '获取清理统计', '登录日志', 'GET'),
(NOW(), NOW(), NULL, '/loginLog/getLogRetentionPolicy', '获取日志保留策略', '登录日志', 'GET'),
(NOW(), NOW(), NULL, '/loginLog/setLogRetentionPolicy', '设置日志保留策略', '登录日志', 'POST'),
(NOW(), NOW(), NULL, '/loginLog/getFailedLoginAttempts', '获取失败登录尝试', '登录日志', 'GET'),
(NOW(), NOW(), NULL, '/loginLog/exportLoginStatistics', '导出登录统计', '登录日志', 'GET');

-- 3. 查找正确的父菜单ID（超级管理员菜单）
SET @parent_menu_id = (SELECT id FROM sys_base_menus WHERE name = 'superAdmin' LIMIT 1);

-- 如果找不到superAdmin菜单，使用默认值
SET @parent_menu_id = IFNULL(@parent_menu_id, '3');

-- 4. 添加登录日志菜单
INSERT IGNORE INTO `sys_base_menus` (
    `created_at`, 
    `updated_at`, 
    `deleted_at`, 
    `parent_id`, 
    `path`, 
    `name`, 
    `hidden`, 
    `component`, 
    `sort`, 
    `active_name`, 
    `keep_alive`, 
    `default_menu`, 
    `title`, 
    `icon`, 
    `close_tab`,
    `transition`
) VALUES (
    NOW(), 
    NOW(), 
    NULL, 
    @parent_menu_id,
    'loginLog', 
    'loginLog', 
    0, 
    'view/superAdmin/loginLog/loginLog.vue', 
    7,
    '', 
    0, 
    0, 
    '登录日志', 
    'list', 
    0,
    ''
);

-- 5. 为超级管理员角色(888)添加API权限
INSERT IGNORE INTO `sys_authority_apis` (`sys_authority_authority_id`, `sys_api_id`)
SELECT 888, `id` FROM `sys_apis` WHERE `api_group` = '登录日志';

-- 6. 为管理员角色(9528)添加API权限（如果存在该角色）
INSERT IGNORE INTO `sys_authority_apis` (`sys_authority_authority_id`, `sys_api_id`)
SELECT 9528, `id` FROM `sys_apis` WHERE `api_group` = '登录日志' 
AND EXISTS (SELECT 1 FROM `sys_authorities` WHERE `authority_id` = 9528);

-- 7. 为超级管理员角色添加菜单权限
INSERT IGNORE INTO `sys_authority_menus` (`sys_authority_authority_id`, `sys_base_menu_id`)
SELECT 888, `id` FROM `sys_base_menus` WHERE `name` = 'loginLog';

-- 8. 为管理员角色添加菜单权限（如果存在该角色）
INSERT IGNORE INTO `sys_authority_menus` (`sys_authority_authority_id`, `sys_base_menu_id`)
SELECT 9528, `id` FROM `sys_base_menus` WHERE `name` = 'loginLog' 
AND EXISTS (SELECT 1 FROM `sys_authorities` WHERE `authority_id` = 9528);

-- 验证结果
SELECT '=== 验证结果 ===' as info;

SELECT '1. 登录日志表是否存在:' as check_item;
SELECT 
    CASE 
        WHEN COUNT(*) > 0 THEN '✓ wy_login_log 表已存在' 
        ELSE '✗ wy_login_log 表不存在' 
    END as result
FROM information_schema.tables 
WHERE table_schema = DATABASE() AND table_name = 'wy_login_log';

SELECT '2. 登录日志API列表:' as check_item;
SELECT id, path, description, method FROM `sys_apis` WHERE `api_group` = '登录日志';

SELECT '3. 登录日志菜单:' as check_item;
SELECT id, name, title, path, parent_id FROM `sys_base_menus` WHERE `name` = 'loginLog';

SELECT '4. 超级管理员API权限:' as check_item;
SELECT 
    a.path, 
    a.description, 
    a.method,
    CASE WHEN aa.sys_api_id IS NOT NULL THEN '✓ 已授权' ELSE '✗ 未授权' END as status
FROM `sys_apis` a
LEFT JOIN `sys_authority_apis` aa ON a.id = aa.sys_api_id AND aa.sys_authority_authority_id = 888
WHERE a.api_group = '登录日志'
ORDER BY a.path;

SELECT '5. 超级管理员菜单权限:' as check_item;
SELECT 
    m.name,
    m.title,
    CASE WHEN am.sys_base_menu_id IS NOT NULL THEN '✓ 已授权' ELSE '✗ 未授权' END as status
FROM `sys_base_menus` m
LEFT JOIN `sys_authority_menus` am ON m.id = am.sys_base_menu_id AND am.sys_authority_authority_id = 888
WHERE m.name = 'loginLog';