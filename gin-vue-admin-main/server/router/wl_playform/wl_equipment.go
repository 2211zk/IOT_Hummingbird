package wl_playform

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type WlEquipmentRouter struct {}

// InitWlEquipmentRouter 初始化 wlEquipment表 路由信息
func (s *WlEquipmentRouter) InitWlEquipmentRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	wlEquipmentRouter := Router.Group("wlEquipment").Use(middleware.OperationRecord())
	wlEquipmentRouterWithoutRecord := Router.Group("wlEquipment")
	wlEquipmentRouterWithoutAuth := PublicRouter.Group("wlEquipment")
	{
		wlEquipmentRouter.POST("createWlEquipment", wlEquipmentApi.CreateWlEquipment)   // 新建wlEquipment表
		wlEquipmentRouter.DELETE("deleteWlEquipment", wlEquipmentApi.DeleteWlEquipment) // 删除wlEquipment表
		wlEquipmentRouter.DELETE("deleteWlEquipmentByIds", wlEquipmentApi.DeleteWlEquipmentByIds) // 批量删除wlEquipment表
		wlEquipmentRouter.PUT("updateWlEquipment", wlEquipmentApi.UpdateWlEquipment)    // 更新wlEquipment表
	}
	{
		wlEquipmentRouterWithoutRecord.GET("findWlEquipment", wlEquipmentApi.FindWlEquipment)        // 根据ID获取wlEquipment表
		wlEquipmentRouterWithoutRecord.GET("getWlEquipmentList", wlEquipmentApi.GetWlEquipmentList)  // 获取wlEquipment表列表
	}
	{
	    wlEquipmentRouterWithoutAuth.GET("getWlEquipmentPublic", wlEquipmentApi.GetWlEquipmentPublic)  // wlEquipment表开放接口
	}
}
