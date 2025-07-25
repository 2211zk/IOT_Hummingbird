package wl_department

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/wl_department/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service/wl_department"
	"github.com/gin-gonic/gin"
)

var WlDepartmentApi = new(WlDepartmentApiStruct)

type WlDepartmentApiStruct struct{}

func (a *WlDepartmentApiStruct) GetWlDepartmentList(c *gin.Context) {
	var req request.WlDepartmentSearch
	_ = c.ShouldBindJSON(&req)
	list, total, err := wl_department.WlDepartmentServiceApp.GetWlDepartmentList(req)
	if err != nil {
		response.FailWithMessage("查询失败", c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     req.Page,
		PageSize: req.PageSize,
	}, "查询成功", c)
}

func (a *WlDepartmentApiStruct) CreateWlDepartment(c *gin.Context) {
	var req request.CreateWlDepartmentReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}
	if err := wl_department.WlDepartmentServiceApp.CreateWlDepartment(req); err != nil {
		response.FailWithMessage("新增失败", c)
		return
	}
	response.OkWithMessage("新增成功", c)
}

func (a *WlDepartmentApiStruct) UpdateWlDepartment(c *gin.Context) {
	var req request.UpdateWlDepartmentReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}
	if err := wl_department.WlDepartmentServiceApp.UpdateWlDepartment(req); err != nil {
		response.FailWithMessage("编辑失败", c)
		return
	}
	response.OkWithMessage("编辑成功", c)
}

func (a *WlDepartmentApiStruct) DeleteWlDepartment(c *gin.Context) {
	var req request.DeleteWlDepartmentReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}
	if err := wl_department.WlDepartmentServiceApp.DeleteWlDepartment(req); err != nil {
		response.FailWithMessage("删除失败", c)
		return
	}
	response.OkWithMessage("删除成功", c)
}
