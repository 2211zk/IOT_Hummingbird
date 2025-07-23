package wl_playform

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type WlResourcesRouter struct{}

// InitWlResourcesRouter 初始化 wlResources 路由信息
func (s *WlResourcesRouter) InitWlResourcesRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	wlResourcesRouter := Router.Group("wlResources").Use(middleware.OperationRecord())
	wlResourcesRouterWithoutRecord := Router.Group("wlResources")
	// wlResourcesRouterWithoutAuth := PublicRouter.Group("wlResources")

	var wlResourcesApi = v1.ApiGroupApp.Wl_playformApiGroup.WlResourcesApi
	{
		wlResourcesRouter.POST("createWlResources", wlResourcesApi.CreateWlResources)                    // 新建wlResources
		wlResourcesRouter.POST("createWithTransaction", wlResourcesApi.CreateWlResourcesWithTransaction) // 新建资源（事务处理）
		wlResourcesRouter.DELETE("deleteWlResources", wlResourcesApi.DeleteWlResources)                  // 删除wlResources
		wlResourcesRouter.DELETE("deleteWlResourcesByIds", wlResourcesApi.DeleteWlResourcesByIds)        // 批量删除wlResources
		wlResourcesRouter.PUT("updateWlResources", wlResourcesApi.UpdateWlResources)                     // 更新wlResources
		wlResourcesRouter.POST("verifyWlResources", wlResourcesApi.VerifyWlResources)                    // 验证wlResources
	}
	{
		wlResourcesRouterWithoutRecord.GET("findWlResources", wlResourcesApi.FindWlResources)       // 根据ID获取wlResources
		wlResourcesRouterWithoutRecord.GET("getWlResourcesList", wlResourcesApi.GetWlResourcesList) // 获取wlResources列表
	}
	// {
	// 	wlResourcesRouterWithoutAuth.GET("getWlResourcesPublic", wlResourcesApi.GetWlResourcesPublic) // 获取wlResources列表
	// }
}
