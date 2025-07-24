package wl_playform

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type WlScenesRouter struct {}

// InitWlScenesRouter 初始化 wlScenes表 路由信息
func (s *WlScenesRouter) InitWlScenesRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	wlScenesRouter := Router.Group("wlScenes").Use(middleware.OperationRecord())
	wlScenesRouterWithoutRecord := Router.Group("wlScenes")
	wlScenesRouterWithoutAuth := PublicRouter.Group("wlScenes")
	{
		wlScenesRouter.POST("createWlScenes", wlScenesApi.CreateWlScenes)   // 新建wlScenes表
		wlScenesRouter.DELETE("deleteWlScenes", wlScenesApi.DeleteWlScenes) // 删除wlScenes表
		wlScenesRouter.DELETE("deleteWlScenesByIds", wlScenesApi.DeleteWlScenesByIds) // 批量删除wlScenes表
		wlScenesRouter.PUT("updateWlScenes", wlScenesApi.UpdateWlScenes)    // 更新wlScenes表
	}
	{
		wlScenesRouterWithoutRecord.GET("findWlScenes", wlScenesApi.FindWlScenes)        // 根据ID获取wlScenes表
		wlScenesRouterWithoutRecord.GET("getWlScenesList", wlScenesApi.GetWlScenesList)  // 获取wlScenes表列表
	}
	{
	    wlScenesRouterWithoutAuth.GET("getWlScenesPublic", wlScenesApi.GetWlScenesPublic)  // wlScenes表开放接口
	}
}
