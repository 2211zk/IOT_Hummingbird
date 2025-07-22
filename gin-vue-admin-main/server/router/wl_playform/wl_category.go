package wl_playform

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type WlCategoryRouter struct {}

// InitWlCategoryRouter 初始化 wlCategory表 路由信息
func (s *WlCategoryRouter) InitWlCategoryRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	wlCategoryRouter := Router.Group("wlCategory").Use(middleware.OperationRecord())
	wlCategoryRouterWithoutRecord := Router.Group("wlCategory")
	wlCategoryRouterWithoutAuth := PublicRouter.Group("wlCategory")
	{
		wlCategoryRouter.POST("createWlCategory", wlCategoryApi.CreateWlCategory)   // 新建wlCategory表
		wlCategoryRouter.DELETE("deleteWlCategory", wlCategoryApi.DeleteWlCategory) // 删除wlCategory表
		wlCategoryRouter.DELETE("deleteWlCategoryByIds", wlCategoryApi.DeleteWlCategoryByIds) // 批量删除wlCategory表
		wlCategoryRouter.PUT("updateWlCategory", wlCategoryApi.UpdateWlCategory)    // 更新wlCategory表
	}
	{
		wlCategoryRouterWithoutRecord.GET("findWlCategory", wlCategoryApi.FindWlCategory)        // 根据ID获取wlCategory表
		wlCategoryRouterWithoutRecord.GET("getWlCategoryList", wlCategoryApi.GetWlCategoryList)  // 获取wlCategory表列表
	}
	{
	    wlCategoryRouterWithoutAuth.GET("getWlCategoryPublic", wlCategoryApi.GetWlCategoryPublic)  // wlCategory表开放接口
	}
}
