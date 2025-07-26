package wl_department

import (
	"strconv"

	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/wl_department/request"
	deptResponse "github.com/flipped-aurora/gin-vue-admin/server/model/wl_department/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service/wl_department"
	"github.com/gin-gonic/gin"
)

var WlDepartmentApi = new(WlDepartmentApiStruct)

type WlDepartmentApiStruct struct{}

// GetWlDepartmentList 获取部门列表
// @Tags WlDepartment
// @Summary 获取部门列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.WlDepartmentSearch true "分页获取部门列表"
// @Success 200 {object} response.Response{data=deptResponse.PageResult,msg=string} "获取成功"
// @Router /department/list [get]
func (a *WlDepartmentApiStruct) GetWlDepartmentList(c *gin.Context) {
	var req request.WlDepartmentSearch

	// 支持GET和POST两种方式
	if c.Request.Method == "GET" {
		if err := c.ShouldBindQuery(&req); err != nil {
			response.FailWithMessage("参数错误: "+err.Error(), c)
			return
		}
	} else {
		if err := c.ShouldBindJSON(&req); err != nil {
			response.FailWithMessage("参数错误: "+err.Error(), c)
			return
		}
	}

	// 设置默认分页参数
	if req.Page <= 0 {
		req.Page = 1
	}
	if req.PageSize <= 0 {
		req.PageSize = 10
	}

	list, total, err := wl_department.WlDepartmentServiceApp.GetWlDepartmentList(req)
	if err != nil {
		response.FailWithMessage("查询失败: "+err.Error(), c)
		return
	}

	// 转换为响应格式
	var responseList []deptResponse.DepartmentListResponse
	for _, dept := range list {
		responseList = append(responseList, *deptResponse.ConvertToListResponse(&dept))
	}

	response.OkWithDetailed(response.PageResult{
		List:     responseList,
		Total:    total,
		Page:     req.Page,
		PageSize: req.PageSize,
	}, "查询成功", c)
}

// GetDepartmentTree 获取部门树
// @Tags WlDepartment
// @Summary 获取部门树
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.DepartmentTreeReq false "获取部门树"
// @Success 200 {object} response.Response{data=[]deptResponse.DepartmentTreeNode,msg=string} "获取成功"
// @Router /department/tree [get]
func (a *WlDepartmentApiStruct) GetDepartmentTree(c *gin.Context) {
	var req request.DepartmentTreeReq
	if err := c.ShouldBindQuery(&req); err != nil {
		response.FailWithMessage("参数错误: "+err.Error(), c)
		return
	}

	tree, err := wl_department.WlDepartmentServiceApp.GetDepartmentTree(req)
	if err != nil {
		response.FailWithMessage("查询失败: "+err.Error(), c)
		return
	}

	response.OkWithDetailed(tree, "查询成功", c)
}

// CreateWlDepartment 创建部门
// @Tags WlDepartment
// @Summary 创建部门
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.CreateWlDepartmentReq true "创建部门"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /department/create [post]
func (a *WlDepartmentApiStruct) CreateWlDepartment(c *gin.Context) {
	var req request.CreateWlDepartmentReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage("参数错误: "+err.Error(), c)
		return
	}

	// 参数验证
	if req.Name == "" && req.DepartmentName == "" {
		response.FailWithMessage("部门名称不能为空", c)
		return
	}

	if err := wl_department.WlDepartmentServiceApp.CreateWlDepartment(req); err != nil {
		// 根据错误类型返回不同的错误码
		switch err.Error() {
		case "同级部门名称已存在":
			response.FailWithCodeMessage(7001, err.Error(), c)
		case "上级部门不存在":
			response.FailWithCodeMessage(7004, err.Error(), c)
		default:
			response.FailWithMessage(err.Error(), c)
		}
		return
	}

	response.OkWithMessage("新增成功", c)
}

// UpdateWlDepartment 更新部门
// @Tags WlDepartment
// @Summary 更新部门
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.UpdateWlDepartmentReq true "更新部门"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /department/update [put]
func (a *WlDepartmentApiStruct) UpdateWlDepartment(c *gin.Context) {
	var req request.UpdateWlDepartmentReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage("参数错误: "+err.Error(), c)
		return
	}

	// 参数验证
	if req.ID <= 0 {
		response.FailWithMessage("部门ID不能为空", c)
		return
	}
	if req.Name == "" && req.DepartmentName == "" {
		response.FailWithMessage("部门名称不能为空", c)
		return
	}

	if err := wl_department.WlDepartmentServiceApp.UpdateWlDepartment(req); err != nil {
		// 根据错误类型返回不同的错误码
		switch err.Error() {
		case "部门不存在":
			response.FailWithCodeMessage(7004, err.Error(), c)
		case "不能选择自身作为上级部门":
			response.FailWithCodeMessage(7002, err.Error(), c)
		case "不能选择子部门作为上级部门":
			response.FailWithCodeMessage(7002, err.Error(), c)
		case "同级部门名称已存在":
			response.FailWithCodeMessage(7001, err.Error(), c)
		case "上级部门不存在":
			response.FailWithCodeMessage(7004, err.Error(), c)
		default:
			response.FailWithMessage(err.Error(), c)
		}
		return
	}

	response.OkWithMessage("编辑成功", c)
}

// DeleteWlDepartment 删除部门
// @Tags WlDepartment
// @Summary 删除部门
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.DeleteWlDepartmentReq true "删除部门"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /department/delete [delete]
func (a *WlDepartmentApiStruct) DeleteWlDepartment(c *gin.Context) {
	var req request.DeleteWlDepartmentReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage("参数错误: "+err.Error(), c)
		return
	}

	// 参数验证
	if req.ID <= 0 {
		response.FailWithMessage("部门ID不能为空", c)
		return
	}

	if err := wl_department.WlDepartmentServiceApp.DeleteWlDepartment(req); err != nil {
		// 根据错误类型返回不同的错误码
		switch err.Error() {
		case "该部门下还有子部门，无法删除":
			response.FailWithCodeMessage(7003, err.Error(), c)
		case "部门不存在":
			response.FailWithCodeMessage(7004, err.Error(), c)
		default:
			response.FailWithMessage(err.Error(), c)
		}
		return
	}

	response.OkWithMessage("删除成功", c)
}

// GetDepartmentDetail 获取部门详情
// @Tags WlDepartment
// @Summary 获取部门详情
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param id path int true "部门ID"
// @Success 200 {object} response.Response{data=deptResponse.DepartmentDetailResponse,msg=string} "获取成功"
// @Router /department/{id} [get]
func (a *WlDepartmentApiStruct) GetDepartmentDetail(c *gin.Context) {
	idStr := c.Param("id")
	idInt, err := strconv.Atoi(idStr)
	if err != nil || idInt <= 0 {
		response.FailWithMessage("无效的部门ID", c)
		return
	}

	id := uint(idInt)
	dept, err := wl_department.WlDepartmentServiceApp.GetDepartmentDetail(id)
	if err != nil {
		response.FailWithMessage("查询失败: "+err.Error(), c)
		return
	}

	// 转换为响应格式
	detailResp := deptResponse.ConvertToDetailResponse(&dept)
	response.OkWithDetailed(detailResp, "查询成功", c)
}

// GetAvailableDevices 获取可关联的设备列表
// @Tags WlDepartment
// @Summary 获取可关联的设备列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.AvailableDevicesReq true "获取可关联设备"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /department/devices/available [get]
func (a *WlDepartmentApiStruct) GetAvailableDevices(c *gin.Context) {
	var req request.AvailableDevicesReq
	if err := c.ShouldBindQuery(&req); err != nil {
		response.FailWithMessage("参数错误: "+err.Error(), c)
		return
	}

	// 设置默认分页参数
	if req.Page <= 0 {
		req.Page = 1
	}
	if req.PageSize <= 0 {
		req.PageSize = 10
	}

	list, total, err := wl_department.WlDepartmentServiceApp.GetAvailableDevices(req)
	if err != nil {
		response.FailWithMessage("查询失败: "+err.Error(), c)
		return
	}

	// 转换为响应格式
	var deviceList []deptResponse.DeviceResponse
	for _, device := range list {
		deviceList = append(deviceList, deptResponse.DeviceResponse{
			ID:          device.ID,
			DeviceName:  device.DeviceName,
			ProductName: device.ProductName,
			Status:      device.Status,
		})
	}

	response.OkWithDetailed(response.PageResult{
		List:     deviceList,
		Total:    total,
		Page:     req.Page,
		PageSize: req.PageSize,
	}, "查询成功", c)
}

// GetDepartmentDevices 获取部门已关联的设备
// @Tags WlDepartment
// @Summary 获取部门已关联的设备
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.DepartmentDevicesReq true "获取部门设备"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /department/devices [get]
func (a *WlDepartmentApiStruct) GetDepartmentDevices(c *gin.Context) {
	var req request.DepartmentDevicesReq
	if err := c.ShouldBindQuery(&req); err != nil {
		response.FailWithMessage("参数错误: "+err.Error(), c)
		return
	}

	// 参数验证
	if req.DepartmentID <= 0 {
		response.FailWithMessage("部门ID不能为空", c)
		return
	}

	// 设置默认分页参数
	if req.Page <= 0 {
		req.Page = 1
	}
	if req.PageSize <= 0 {
		req.PageSize = 10
	}

	list, total, err := wl_department.WlDepartmentServiceApp.GetDepartmentDevices(req)
	if err != nil {
		response.FailWithMessage("查询失败: "+err.Error(), c)
		return
	}

	// 转换为响应格式
	var deviceList []deptResponse.DeviceResponse
	for _, device := range list {
		deviceList = append(deviceList, deptResponse.DeviceResponse{
			ID:          device.ID,
			DeviceName:  device.DeviceName,
			ProductName: device.ProductName,
			Status:      device.Status,
		})
	}

	response.OkWithDetailed(response.PageResult{
		List:     deviceList,
		Total:    total,
		Page:     req.Page,
		PageSize: req.PageSize,
	}, "查询成功", c)
}
