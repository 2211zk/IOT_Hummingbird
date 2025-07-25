package wl_playform

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type WlUserRouter struct {}

// InitWlUserRouter 初始化 wlUser表 路由信息
func (s *WlUserRouter) InitWlUserRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	wlUserRouter := Router.Group("wlUser").Use(middleware.OperationRecord())
	wlUserRouterWithoutRecord := Router.Group("wlUser")
	wlUserRouterWithoutAuth := PublicRouter.Group("wlUser")
	{
		wlUserRouter.POST("createWlUser", wlUserApi.CreateWlUser)   // 新建wlUser表
		wlUserRouter.DELETE("deleteWlUser", wlUserApi.DeleteWlUser) // 删除wlUser表
		wlUserRouter.DELETE("deleteWlUserByIds", wlUserApi.DeleteWlUserByIds) // 批量删除wlUser表
		wlUserRouter.PUT("updateWlUser", wlUserApi.UpdateWlUser)    // 更新wlUser表
	}
	{
		wlUserRouterWithoutRecord.GET("findWlUser", wlUserApi.FindWlUser)        // 根据ID获取wlUser表
		wlUserRouterWithoutRecord.GET("getWlUserList", wlUserApi.GetWlUserList)  // 获取wlUser表列表
	}
	{
	    wlUserRouterWithoutAuth.GET("getWlUserPublic", wlUserApi.GetWlUserPublic)  // wlUser表开放接口
	}
}
