package wl_playform

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type WlCaFunctionRouter struct {}

// InitWlCaFunctionRouter 初始化 wlCaFunction表 路由信息
func (s *WlCaFunctionRouter) InitWlCaFunctionRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	wlCaFunctionRouter := Router.Group("wlCaFunction").Use(middleware.OperationRecord())
	wlCaFunctionRouterWithoutRecord := Router.Group("wlCaFunction")
	wlCaFunctionRouterWithoutAuth := PublicRouter.Group("wlCaFunction")
	{
		wlCaFunctionRouter.POST("createWlCaFunction", wlCaFunctionApi.CreateWlCaFunction)   // 新建wlCaFunction表
		wlCaFunctionRouter.DELETE("deleteWlCaFunction", wlCaFunctionApi.DeleteWlCaFunction) // 删除wlCaFunction表
		wlCaFunctionRouter.DELETE("deleteWlCaFunctionByIds", wlCaFunctionApi.DeleteWlCaFunctionByIds) // 批量删除wlCaFunction表
		wlCaFunctionRouter.PUT("updateWlCaFunction", wlCaFunctionApi.UpdateWlCaFunction)    // 更新wlCaFunction表
	}
	{
		wlCaFunctionRouterWithoutRecord.GET("findWlCaFunction", wlCaFunctionApi.FindWlCaFunction)        // 根据ID获取wlCaFunction表
		wlCaFunctionRouterWithoutRecord.GET("getWlCaFunctionList", wlCaFunctionApi.GetWlCaFunctionList)  // 获取wlCaFunction表列表
	}
	{
	    wlCaFunctionRouterWithoutAuth.GET("getWlCaFunctionPublic", wlCaFunctionApi.GetWlCaFunctionPublic)  // wlCaFunction表开放接口
	}
}
