package wl_playform

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type WlEngineRulesRouter struct {}

// InitWlEngineRulesRouter 初始化 wlEngineRules表 路由信息
func (s *WlEngineRulesRouter) InitWlEngineRulesRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	wlEngineRulesRouter := Router.Group("wlEngineRules").Use(middleware.OperationRecord())
	wlEngineRulesRouterWithoutRecord := Router.Group("wlEngineRules")
	wlEngineRulesRouterWithoutAuth := PublicRouter.Group("wlEngineRules")
	{
		wlEngineRulesRouter.POST("createWlEngineRules", wlEngineRulesApi.CreateWlEngineRules)   // 新建wlEngineRules表
		wlEngineRulesRouter.DELETE("deleteWlEngineRules", wlEngineRulesApi.DeleteWlEngineRules) // 删除wlEngineRules表
		wlEngineRulesRouter.DELETE("deleteWlEngineRulesByIds", wlEngineRulesApi.DeleteWlEngineRulesByIds) // 批量删除wlEngineRules表
		wlEngineRulesRouter.PUT("updateWlEngineRules", wlEngineRulesApi.UpdateWlEngineRules)    // 更新wlEngineRules表
	}
	{
		wlEngineRulesRouterWithoutRecord.GET("findWlEngineRules", wlEngineRulesApi.FindWlEngineRules)        // 根据ID获取wlEngineRules表
		wlEngineRulesRouterWithoutRecord.GET("getWlEngineRulesList", wlEngineRulesApi.GetWlEngineRulesList)  // 获取wlEngineRules表列表
	}
	{
	    wlEngineRulesRouterWithoutAuth.GET("getWlEngineRulesPublic", wlEngineRulesApi.GetWlEngineRulesPublic)  // wlEngineRules表开放接口
	}
}
