package wl_driver

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type WlDriversRouter struct {}

// InitWlDriversRouter 初始化 wlDrivers表 路由信息
func (s *WlDriversRouter) InitWlDriversRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	wlDriversRouter := Router.Group("wlDrivers").Use(middleware.OperationRecord())
	wlDriversRouterWithoutRecord := Router.Group("wlDrivers")
	wlDriversRouterWithoutAuth := PublicRouter.Group("wlDrivers")
	{
		wlDriversRouter.POST("createWlDrivers", wlDriversApi.CreateWlDrivers)   // 新建wlDrivers表
		wlDriversRouter.DELETE("deleteWlDrivers", wlDriversApi.DeleteWlDrivers) // 删除wlDrivers表
		wlDriversRouter.DELETE("deleteWlDriversByIds", wlDriversApi.DeleteWlDriversByIds) // 批量删除wlDrivers表
		wlDriversRouter.PUT("updateWlDrivers", wlDriversApi.UpdateWlDrivers)    // 更新wlDrivers表
	}
	{
		wlDriversRouterWithoutRecord.GET("findWlDrivers", wlDriversApi.FindWlDrivers)        // 根据ID获取wlDrivers表
		wlDriversRouterWithoutRecord.GET("getWlDriversList", wlDriversApi.GetWlDriversList)  // 获取wlDrivers表列表
	}
	{
	    wlDriversRouterWithoutAuth.GET("getWlDriversPublic", wlDriversApi.GetWlDriversPublic)  // wlDrivers表开放接口
	}
}
