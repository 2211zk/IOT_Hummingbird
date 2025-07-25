package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"go.uber.org/zap"
)

// FixPermissions 修复权限问题
func FixPermissions() {
	global.GVA_LOG.Info("开始修复权限问题...")
	
	// 1. 确保超级管理员角色存在
	var count int64
	global.GVA_DB.Model(&struct{}{}).Table("sys_authorities").Where("authority_id = ?", 888).Count(&count)
	if count == 0 {
		global.GVA_DB.Exec(`INSERT IGNORE INTO sys_authorities (authority_id, authority_name, parent_id, default_router, created_at, updated_at) 
			VALUES (888, '超级管理员', 0, 'dashboard', NOW(), NOW())`)
		global.GVA_LOG.Info("创建超级管理员角色")
	}

	// 2. 确保普通用户角色存在
	global.GVA_DB.Model(&struct{}{}).Table("sys_authorities").Where("authority_id = ?", 8881).Count(&count)
	if count == 0 {
		global.GVA_DB.Exec(`INSERT IGNORE INTO sys_authorities (authority_id, authority_name, parent_id, default_router, created_at, updated_at) 
			VALUES (8881, '普通用户', 888, 'dashboard', NOW(), NOW())`)
		global.GVA_LOG.Info("创建普通用户角色")
	}

	// 3. 确保测试角色存在
	global.GVA_DB.Model(&struct{}{}).Table("sys_authorities").Where("authority_id = ?", 9528).Count(&count)
	if count == 0 {
		global.GVA_DB.Exec(`INSERT IGNORE INTO sys_authorities (authority_id, authority_name, parent_id, default_router, created_at, updated_at) 
			VALUES (9528, '测试角色', 0, 'dashboard', NOW(), NOW())`)
		global.GVA_LOG.Info("创建测试角色")
	}

	// 4. 给所有角色分配所有菜单权限
	global.GVA_DB.Exec(`INSERT IGNORE INTO sys_authority_menus (sys_base_menu_id, sys_authority_authority_id)
		SELECT id, 888 FROM sys_base_menus WHERE id > 0`)
	global.GVA_DB.Exec(`INSERT IGNORE INTO sys_authority_menus (sys_base_menu_id, sys_authority_authority_id)
		SELECT id, 8881 FROM sys_base_menus WHERE id > 0`)
	global.GVA_DB.Exec(`INSERT IGNORE INTO sys_authority_menus (sys_base_menu_id, sys_authority_authority_id)
		SELECT id, 9528 FROM sys_base_menus WHERE id > 0`)

	// 5. 添加关键的API权限
	criticalPermissions := []string{
		"('p', '888', '/dashboard/getDashboardData', 'GET')",
		"('p', '8881', '/dashboard/getDashboardData', 'GET')",
		"('p', '9528', '/dashboard/getDashboardData', 'GET')",
		"('p', '888', '/user/getUserInfo', 'GET')",
		"('p', '8881', '/user/getUserInfo', 'GET')",
		"('p', '9528', '/user/getUserInfo', 'GET')",
		"('p', '888', '/wlProducts/*', 'GET')",
		"('p', '888', '/wlProducts/*', 'POST')",
		"('p', '888', '/wlProducts/*', 'PUT')",
		"('p', '888', '/wlProducts/*', 'DELETE')",
		"('p', '8881', '/wlProducts/*', 'GET')",
		"('p', '8881', '/wlProducts/*', 'POST')",
		"('p', '8881', '/wlProducts/*', 'PUT')",
		"('p', '8881', '/wlProducts/*', 'DELETE')",
		"('p', '9528', '/wlProducts/*', 'GET')",
		"('p', '9528', '/wlProducts/*', 'POST')",
		"('p', '9528', '/wlProducts/*', 'PUT')",
		"('p', '9528', '/wlProducts/*', 'DELETE')",
		"('p', '888', '/wlEquipment/*', 'GET')",
		"('p', '888', '/wlEquipment/*', 'POST')",
		"('p', '888', '/wlEquipment/*', 'PUT')",
		"('p', '888', '/wlEquipment/*', 'DELETE')",
		"('p', '8881', '/wlEquipment/*', 'GET')",
		"('p', '8881', '/wlEquipment/*', 'POST')",
		"('p', '8881', '/wlEquipment/*', 'PUT')",
		"('p', '8881', '/wlEquipment/*', 'DELETE')",
		"('p', '9528', '/wlEquipment/*', 'GET')",
		"('p', '9528', '/wlEquipment/*', 'POST')",
		"('p', '9528', '/wlEquipment/*', 'PUT')",
		"('p', '9528', '/wlEquipment/*', 'DELETE')",
		"('p', '888', '/wlDrivers/*', 'GET')",
		"('p', '888', '/wlDrivers/*', 'POST')",
		"('p', '888', '/wlDrivers/*', 'PUT')",
		"('p', '888', '/wlDrivers/*', 'DELETE')",
		"('p', '8881', '/wlDrivers/*', 'GET')",
		"('p', '8881', '/wlDrivers/*', 'POST')",
		"('p', '8881', '/wlDrivers/*', 'PUT')",
		"('p', '8881', '/wlDrivers/*', 'DELETE')",
		"('p', '9528', '/wlDrivers/*', 'GET')",
		"('p', '9528', '/wlDrivers/*', 'POST')",
		"('p', '9528', '/wlDrivers/*', 'PUT')",
		"('p', '9528', '/wlDrivers/*', 'DELETE')",
		"('p', '888', '/wlResources/*', 'GET')",
		"('p', '888', '/wlResources/*', 'POST')",
		"('p', '888', '/wlResources/*', 'PUT')",
		"('p', '888', '/wlResources/*', 'DELETE')",
		"('p', '8881', '/wlResources/*', 'GET')",
		"('p', '8881', '/wlResources/*', 'POST')",
		"('p', '8881', '/wlResources/*', 'PUT')",
		"('p', '8881', '/wlResources/*', 'DELETE')",
		"('p', '9528', '/wlResources/*', 'GET')",
		"('p', '9528', '/wlResources/*', 'POST')",
		"('p', '9528', '/wlResources/*', 'PUT')",
		"('p', '9528', '/wlResources/*', 'DELETE')",
		"('p', '888', '/wlScenes/*', 'GET')",
		"('p', '888', '/wlScenes/*', 'POST')",
		"('p', '888', '/wlScenes/*', 'PUT')",
		"('p', '888', '/wlScenes/*', 'DELETE')",
		"('p', '8881', '/wlScenes/*', 'GET')",
		"('p', '8881', '/wlScenes/*', 'POST')",
		"('p', '8881', '/wlScenes/*', 'PUT')",
		"('p', '8881', '/wlScenes/*', 'DELETE')",
		"('p', '9528', '/wlScenes/*', 'GET')",
		"('p', '9528', '/wlScenes/*', 'POST')",
		"('p', '9528', '/wlScenes/*', 'PUT')",
		"('p', '9528', '/wlScenes/*', 'DELETE')",
		"('p', '888', '/wlEngineRules/*', 'GET')",
		"('p', '888', '/wlEngineRules/*', 'POST')",
		"('p', '888', '/wlEngineRules/*', 'PUT')",
		"('p', '888', '/wlEngineRules/*', 'DELETE')",
		"('p', '8881', '/wlEngineRules/*', 'GET')",
		"('p', '8881', '/wlEngineRules/*', 'POST')",
		"('p', '8881', '/wlEngineRules/*', 'PUT')",
		"('p', '8881', '/wlEngineRules/*', 'DELETE')",
		"('p', '9528', '/wlEngineRules/*', 'GET')",
		"('p', '9528', '/wlEngineRules/*', 'POST')",
		"('p', '9528', '/wlEngineRules/*', 'PUT')",
		"('p', '9528', '/wlEngineRules/*', 'DELETE')",
		"('p', '888', '/fileUploadAndDownload/*', 'GET')",
		"('p', '888', '/fileUploadAndDownload/*', 'POST')",
		"('p', '888', '/fileUploadAndDownload/*', 'PUT')",
		"('p', '888', '/fileUploadAndDownload/*', 'DELETE')",
		"('p', '8881', '/fileUploadAndDownload/*', 'GET')",
		"('p', '8881', '/fileUploadAndDownload/*', 'POST')",
		"('p', '8881', '/fileUploadAndDownload/*', 'PUT')",
		"('p', '8881', '/fileUploadAndDownload/*', 'DELETE')",
		"('p', '9528', '/fileUploadAndDownload/*', 'GET')",
		"('p', '9528', '/fileUploadAndDownload/*', 'POST')",
		"('p', '9528', '/fileUploadAndDownload/*', 'PUT')",
		"('p', '9528', '/fileUploadAndDownload/*', 'DELETE')",
		"('p', '888', '/jwt/*', 'GET')",
		"('p', '888', '/jwt/*', 'POST')",
		"('p', '888', '/jwt/*', 'PUT')",
		"('p', '888', '/jwt/*', 'DELETE')",
		"('p', '8881', '/jwt/*', 'GET')",
		"('p', '8881', '/jwt/*', 'POST')",
		"('p', '8881', '/jwt/*', 'PUT')",
		"('p', '8881', '/jwt/*', 'DELETE')",
		"('p', '9528', '/jwt/*', 'GET')",
		"('p', '9528', '/jwt/*', 'POST')",
		"('p', '9528', '/jwt/*', 'PUT')",
		"('p', '9528', '/jwt/*', 'DELETE')",
		"('p', '888', '/captcha/*', 'GET')",
		"('p', '888', '/captcha/*', 'POST')",
		"('p', '888', '/captcha/*', 'PUT')",
		"('p', '888', '/captcha/*', 'DELETE')",
		"('p', '8881', '/captcha/*', 'GET')",
		"('p', '8881', '/captcha/*', 'POST')",
		"('p', '8881', '/captcha/*', 'PUT')",
		"('p', '8881', '/captcha/*', 'DELETE')",
		"('p', '9528', '/captcha/*', 'GET')",
		"('p', '9528', '/captcha/*', 'POST')",
		"('p', '9528', '/captcha/*', 'PUT')",
		"('p', '9528', '/captcha/*', 'DELETE')",
	}

	for _, permission := range criticalPermissions {
		global.GVA_DB.Exec("INSERT IGNORE INTO casbin_rule (ptype, v0, v1, v2) VALUES " + permission)
	}

	// 6. 确保当前用户有正确的权限ID
	global.GVA_DB.Exec("UPDATE sys_users SET authority_id = 888 WHERE authority_id NOT IN (888, 8881, 9528)")

	// 7. 刷新Casbin权限
	e := utils.GetCasbin()
	if e != nil {
		err := e.LoadPolicy()
		if err != nil {
			global.GVA_LOG.Error("刷新Casbin权限失败", zap.Error(err))
		} else {
			global.GVA_LOG.Info("刷新Casbin权限成功")
		}
	}

	global.GVA_LOG.Info("权限修复完成")
} 