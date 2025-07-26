package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"go.uber.org/zap"
)

// InitLoginLogMenu 初始化登录日志菜单
func InitLoginLogMenu() {
	// 检查菜单是否已存在
	var count int64
	global.GVA_DB.Model(&system.SysBaseMenu{}).Where("name = ?", "loginLog").Count(&count)
	if count > 0 {
		global.GVA_LOG.Info("登录日志菜单已存在，跳过初始化")
		return
	}

	// 创建登录日志菜单
	loginLogMenu := system.SysBaseMenu{
		ParentId:  3, // 假设3是超级管理员菜单的ID
		Path:      "loginLog",
		Name:      "loginLog",
		Hidden:    false,
		Component: "view/superAdmin/loginLog/loginLog.vue",
		Sort:      7,
		Meta: system.Meta{
			Title:       "登录日志",
			Icon:        "list",
			KeepAlive:   false,
			DefaultMenu: false,
			CloseTab:    false,
		},
	}

	// 插入菜单
	err := global.GVA_DB.Create(&loginLogMenu).Error
	if err != nil {
		global.GVA_LOG.Error("创建登录日志菜单失败", zap.Error(err))
		return
	}

	global.GVA_LOG.Info("登录日志菜单创建成功", zap.Uint("menuId", loginLogMenu.ID))

	// 定义API列表
	apis := []system.SysApi{
		{Path: "/loginLog/getLoginLogList", Description: "获取登录日志列表", ApiGroup: "登录日志", Method: "POST"},
		{Path: "/loginLog/findLoginLog", Description: "获取登录日志详情", ApiGroup: "登录日志", Method: "GET"},
		{Path: "/loginLog/deleteLoginLog", Description: "删除登录日志", ApiGroup: "登录日志", Method: "DELETE"},
		{Path: "/loginLog/deleteLoginLogByIds", Description: "批量删除登录日志", ApiGroup: "登录日志", Method: "DELETE"},
		{Path: "/loginLog/updateLoginLog", Description: "更新登录日志", ApiGroup: "登录日志", Method: "PUT"},
		{Path: "/loginLog/createLoginLog", Description: "创建登录日志", ApiGroup: "登录日志", Method: "POST"},
		{Path: "/loginLog/exportLoginLog", Description: "导出登录日志", ApiGroup: "登录日志", Method: "POST"},
		{Path: "/loginLog/cleanExpiredLogs", Description: "清理过期日志", ApiGroup: "登录日志", Method: "POST"},
		{Path: "/loginLog/getLoginStatistics", Description: "获取登录统计", ApiGroup: "登录日志", Method: "GET"},
		{Path: "/loginLog/getRecentLoginLogs", Description: "获取最近登录日志", ApiGroup: "登录日志", Method: "GET"},
		{Path: "/loginLog/getTopLoginIPs", Description: "获取热门登录IP", ApiGroup: "登录日志", Method: "GET"},
		{Path: "/loginLog/exportLoginStatistics", Description: "导出登录统计", ApiGroup: "登录日志", Method: "GET"},
		{Path: "/loginLog/getCleanupStatistics", Description: "获取清理统计", ApiGroup: "登录日志", Method: "GET"},
		{Path: "/loginLog/getLogRetentionPolicy", Description: "获取日志保留策略", ApiGroup: "登录日志", Method: "GET"},
		{Path: "/loginLog/setLogRetentionPolicy", Description: "设置日志保留策略", ApiGroup: "登录日志", Method: "POST"},
		{Path: "/loginLog/getFailedLoginAttempts", Description: "获取失败登录尝试", ApiGroup: "登录日志", Method: "GET"},
	}

	// 为所有管理员角色分配菜单和API权限
	var authorities []system.SysAuthority
	global.GVA_DB.Find(&authorities)
	for _, authority := range authorities {
		// 分配菜单权限
		err = global.GVA_DB.Model(&authority).Association("SysBaseMenus").Append(&loginLogMenu)
		if err != nil {
			global.GVA_LOG.Error("为角色添加登录日志菜单权限失败", zap.Uint("authorityId", authority.AuthorityId), zap.Error(err))
		}
		// 分配API权限
		for _, api := range apis {
			err = global.GVA_DB.Model(&authority).Association("SysApis").Append(&api)
			if err != nil {
				global.GVA_LOG.Error("为角色添加API权限失败", zap.Uint("authorityId", authority.AuthorityId), zap.String("path", api.Path), zap.Error(err))
			}
		}
	}

	// 定义首页相关API
	dashboardApis := []system.SysApi{
		{Path: "/user/getUserInfo", Description: "获取用户信息", ApiGroup: "用户", Method: "GET"},
		{Path: "/user/getMenu", Description: "获取菜单", ApiGroup: "用户", Method: "GET"},
		{Path: "/dashboard/getPanelData", Description: "获取首页面板数据", ApiGroup: "首页", Method: "GET"},
		// 如有其他首页相关API请补充
	}
	for _, authority := range authorities {
		// 分配首页API权限
		for _, api := range dashboardApis {
			err = global.GVA_DB.Model(&authority).Association("SysApis").Append(&api)
			if err != nil {
				global.GVA_LOG.Error("为角色添加首页API权限失败", zap.Uint("authorityId", authority.AuthorityId), zap.String("path", api.Path), zap.Error(err))
			}
		}
	}

	global.GVA_LOG.Info("登录日志菜单和API初始化完成")
}
