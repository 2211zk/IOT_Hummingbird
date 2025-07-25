package wl_playform

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type WlAlarmRouter struct {}

// InitWlAlarmRouter 初始化 wlAlarm表 路由信息
func (s *WlAlarmRouter) InitWlAlarmRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	wlAlarmRouter := Router.Group("wlAlarm").Use(middleware.OperationRecord())
	wlAlarmRouterWithoutRecord := Router.Group("wlAlarm")
	wlAlarmRouterWithoutAuth := PublicRouter.Group("wlAlarm")
	{
		wlAlarmRouter.POST("createWlAlarm", wlAlarmApi.CreateWlAlarm)   // 新建wlAlarm表
		wlAlarmRouter.DELETE("deleteWlAlarm", wlAlarmApi.DeleteWlAlarm) // 删除wlAlarm表
		wlAlarmRouter.DELETE("deleteWlAlarmByIds", wlAlarmApi.DeleteWlAlarmByIds) // 批量删除wlAlarm表
		wlAlarmRouter.PUT("updateWlAlarm", wlAlarmApi.UpdateWlAlarm)    // 更新wlAlarm表
	}
	{
		wlAlarmRouterWithoutRecord.GET("findWlAlarm", wlAlarmApi.FindWlAlarm)        // 根据ID获取wlAlarm表
		wlAlarmRouterWithoutRecord.GET("getWlAlarmList", wlAlarmApi.GetWlAlarmList)  // 获取wlAlarm表列表
	}
	{
	    wlAlarmRouterWithoutAuth.GET("getWlAlarmPublic", wlAlarmApi.GetWlAlarmPublic)  // wlAlarm表开放接口
	}
}
