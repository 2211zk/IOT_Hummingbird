package wl_playform

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type WlProductsRouter struct {}

// InitWlProductsRouter 初始化 wlProducts表 路由信息
func (s *WlProductsRouter) InitWlProductsRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	wlProductsRouter := Router.Group("wlProducts").Use(middleware.OperationRecord())
	wlProductsRouterWithoutRecord := Router.Group("wlProducts")
	wlProductsRouterWithoutAuth := PublicRouter.Group("wlProducts")
	{
		wlProductsRouter.POST("createWlProducts", wlProductsApi.CreateWlProducts)   // 新建wlProducts表
		wlProductsRouter.DELETE("deleteWlProducts", wlProductsApi.DeleteWlProducts) // 删除wlProducts表
		wlProductsRouter.DELETE("deleteWlProductsByIds", wlProductsApi.DeleteWlProductsByIds) // 批量删除wlProducts表
		wlProductsRouter.PUT("updateWlProducts", wlProductsApi.UpdateWlProducts)    // 更新wlProducts表
	}
	{
		wlProductsRouterWithoutRecord.GET("findWlProducts", wlProductsApi.FindWlProducts)        // 根据ID获取wlProducts表
		wlProductsRouterWithoutRecord.GET("getWlProductsList", wlProductsApi.GetWlProductsList)  // 获取wlProducts表列表
	}
	{
	    wlProductsRouterWithoutAuth.GET("getWlProductsPublic", wlProductsApi.GetWlProductsPublic)  // wlProducts表开放接口
	}
}
