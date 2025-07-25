package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type LoginLogRouter struct{}

// InitLoginLogRouter 初始化登录日志路由信息
func (s *LoginLogRouter) InitLoginLogRouter(Router *gin.RouterGroup) {
	loginLogRouter := Router.Group("loginLog").Use(middleware.OperationRecord())
	loginLogRouterWithoutRecord := Router.Group("loginLog")
	loginLogApi := v1.ApiGroupApp.SystemApiGroup.LoginLogApi
	{
		loginLogRouter.POST("createLoginLog", loginLogApi.CreateLoginLog)               // 新建登录日志
		loginLogRouter.DELETE("deleteLoginLog", loginLogApi.DeleteLoginLog)             // 删除登录日志
		loginLogRouter.DELETE("deleteLoginLogByIds", loginLogApi.DeleteLoginLogByIds)   // 批量删除登录日志
		loginLogRouter.PUT("updateLoginLog", loginLogApi.UpdateLoginLog)                // 更新登录日志
		loginLogRouter.POST("exportLoginLog", loginLogApi.ExportLoginLog)               // 导出登录日志
		loginLogRouter.POST("cleanExpiredLogs", loginLogApi.CleanExpiredLogs)           // 清理过期日志
		loginLogRouter.POST("setLogRetentionPolicy", loginLogApi.SetLogRetentionPolicy) // 设置日志保留策略
	}
	{
		loginLogRouterWithoutRecord.GET("findLoginLog", loginLogApi.FindLoginLog)                     // 根据ID获取登录日志
		loginLogRouterWithoutRecord.POST("getLoginLogList", loginLogApi.GetLoginLogList)              // 获取登录日志列表
		loginLogRouterWithoutRecord.GET("getLoginStatistics", loginLogApi.GetLoginStatistics)         // 获取登录统计信息
		loginLogRouterWithoutRecord.GET("getRecentLoginLogs", loginLogApi.GetRecentLoginLogs)         // 获取最近登录日志
		loginLogRouterWithoutRecord.GET("getTopLoginIPs", loginLogApi.GetTopLoginIPs)                 // 获取热门登录IP
		loginLogRouterWithoutRecord.GET("exportLoginStatistics", loginLogApi.ExportLoginStatistics)   // 导出登录统计信息
		loginLogRouterWithoutRecord.GET("getCleanupStatistics", loginLogApi.GetCleanupStatistics)     // 获取清理统计信息
		loginLogRouterWithoutRecord.GET("getLogRetentionPolicy", loginLogApi.GetLogRetentionPolicy)   // 获取日志保留策略
		loginLogRouterWithoutRecord.GET("getFailedLoginAttempts", loginLogApi.GetFailedLoginAttempts) // 获取失败登录尝试次数
	}
}
