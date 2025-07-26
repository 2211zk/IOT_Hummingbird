package wl_department

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/wl_department"
	"github.com/gin-gonic/gin"
)

func (s *RouterGroup) InitWlDepartmentRouter(Router *gin.RouterGroup) {
	depRouter := Router.Group("department")
	{
		// 部门基础CRUD操作
		depRouter.GET("list", wl_department.WlDepartmentApi.GetWlDepartmentList)     // 获取部门列表
		depRouter.POST("list", wl_department.WlDepartmentApi.GetWlDepartmentList)    // 兼容POST方式
		depRouter.GET("tree", wl_department.WlDepartmentApi.GetDepartmentTree)       // 获取部门树
		depRouter.GET(":id", wl_department.WlDepartmentApi.GetDepartmentDetail)      // 获取部门详情
		depRouter.POST("create", wl_department.WlDepartmentApi.CreateWlDepartment)   // 创建部门
		depRouter.PUT("update", wl_department.WlDepartmentApi.UpdateWlDepartment)    // 更新部门
		depRouter.DELETE("delete", wl_department.WlDepartmentApi.DeleteWlDepartment) // 删除部门

		// 设备关联相关接口
		depRouter.GET("devices/available", wl_department.WlDepartmentApi.GetAvailableDevices) // 获取可关联设备
		depRouter.GET("devices", wl_department.WlDepartmentApi.GetDepartmentDevices)          // 获取部门设备

		// 兼容旧接口
		depRouter.POST("getWlDepartmentList", wl_department.WlDepartmentApi.GetWlDepartmentList)
		depRouter.POST("createWlDepartment", wl_department.WlDepartmentApi.CreateWlDepartment)
		depRouter.POST("updateWlDepartment", wl_department.WlDepartmentApi.UpdateWlDepartment)
		depRouter.POST("deleteWlDepartment", wl_department.WlDepartmentApi.DeleteWlDepartment)
	}
}
