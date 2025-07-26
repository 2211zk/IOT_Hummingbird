package wl_playform

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type WlDepartmentRouter struct{}

// InitWlDepartmentRouter 初始化 wlDepartment表 路由信息
func (s *WlDepartmentRouter) InitWlDepartmentRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	wlDepartmentRouter := Router.Group("wlDepartment").Use(middleware.OperationRecord())
	wlDepartmentRouterWithoutRecord := Router.Group("wlDepartment")
	wlDepartmentRouterWithoutAuth := PublicRouter.Group("wlDepartment")
	{
		wlDepartmentRouter.POST("createWlDepartment", wlDepartmentApi.CreateWlDepartment)                 // 新建wlDepartment表
		wlDepartmentRouter.DELETE("deleteWlDepartment", wlDepartmentApi.DeleteWlDepartment)               // 删除wlDepartment表
		wlDepartmentRouter.DELETE("deleteWlDepartmentByIds", wlDepartmentApi.DeleteWlDepartmentByIds)     // 批量删除wlDepartment表
		wlDepartmentRouter.PUT("updateWlDepartment", wlDepartmentApi.UpdateWlDepartment)                  // 更新wlDepartment表
		wlDepartmentRouter.POST("assignDevicesToDepartment", wlDepartmentApi.AssignDevicesToDepartment)   // 分配设备到部门
		wlDepartmentRouter.POST("getDevicesByDepartment", wlDepartmentApi.GetDevicesByDepartment)         // 查询部门下所有设备
		wlDepartmentRouter.POST("removeDeviceFromDepartment", wlDepartmentApi.RemoveDeviceFromDepartment) // 移除部门下某个设备
	}
	{
		wlDepartmentRouterWithoutRecord.GET("findWlDepartment", wlDepartmentApi.FindWlDepartment)       // 根据ID获取wlDepartment表
		wlDepartmentRouterWithoutRecord.GET("getWlDepartmentList", wlDepartmentApi.GetWlDepartmentList) // 获取wlDepartment表列表
	}
	{
		wlDepartmentRouterWithoutAuth.GET("getWlDepartmentPublic", wlDepartmentApi.GetWlDepartmentPublic) // wlDepartment表开放接口
	}
}
