package wl_department

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/wl_department"
	"github.com/gin-gonic/gin"
)

func InitWlDepartmentRouter(Router *gin.RouterGroup) {
	depRouter := Router.Group("department")
	{
		depRouter.POST("getWlDepartmentList", wl_department.WlDepartmentApi.GetWlDepartmentList)
		depRouter.POST("createWlDepartment", wl_department.WlDepartmentApi.CreateWlDepartment)
		depRouter.POST("updateWlDepartment", wl_department.WlDepartmentApi.UpdateWlDepartment)
		depRouter.POST("deleteWlDepartment", wl_department.WlDepartmentApi.DeleteWlDepartment)
	}
}
