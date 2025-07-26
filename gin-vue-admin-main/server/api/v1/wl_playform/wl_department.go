package wl_playform

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/wl_playform"
	wl_playformReq "github.com/flipped-aurora/gin-vue-admin/server/model/wl_playform/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type WlDepartmentApi struct{}

// CreateWlDepartment 创建wlDepartment表
// @Tags WlDepartment
// @Summary 创建wlDepartment表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body wl_playform.WlDepartment true "创建wlDepartment表"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /wlDepartment/createWlDepartment [post]
func (wlDepartmentApi *WlDepartmentApi) CreateWlDepartment(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var wlDepartment wl_playform.WlDepartment
	err := c.ShouldBindJSON(&wlDepartment)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	userID := utils.GetUserID(c)
	userIDInt := int(userID)
	wlDepartment.CreatedBy = &userIDInt
	err = wlDepartmentService.CreateWlDepartment(ctx, &wlDepartment)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteWlDepartment 删除wlDepartment表
// @Tags WlDepartment
// @Summary 删除wlDepartment表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body wl_playform.WlDepartment true "删除wlDepartment表"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /wlDepartment/deleteWlDepartment [delete]
func (wlDepartmentApi *WlDepartmentApi) DeleteWlDepartment(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ID := c.Query("ID")
	userID := utils.GetUserID(c)
	err := wlDepartmentService.DeleteWlDepartment(ctx, ID, userID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteWlDepartmentByIds 批量删除wlDepartment表
// @Tags WlDepartment
// @Summary 批量删除wlDepartment表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /wlDepartment/deleteWlDepartmentByIds [delete]
func (wlDepartmentApi *WlDepartmentApi) DeleteWlDepartmentByIds(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
	userID := utils.GetUserID(c)
	err := wlDepartmentService.DeleteWlDepartmentByIds(ctx, IDs, userID)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateWlDepartment 更新wlDepartment表
// @Tags WlDepartment
// @Summary 更新wlDepartment表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body wl_playform.WlDepartment true "更新wlDepartment表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /wlDepartment/updateWlDepartment [put]
func (wlDepartmentApi *WlDepartmentApi) UpdateWlDepartment(c *gin.Context) {
	// 从ctx获取标准context进行业务行为
	ctx := c.Request.Context()

	var wlDepartment wl_playform.WlDepartment
	err := c.ShouldBindJSON(&wlDepartment)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	// 修复类型不匹配
	userID := utils.GetUserID(c)
	userIDInt := int(userID)
	wlDepartment.UpdatedBy = &userIDInt
	err = wlDepartmentService.UpdateWlDepartment(ctx, wlDepartment)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindWlDepartment 用id查询wlDepartment表
// @Tags WlDepartment
// @Summary 用id查询wlDepartment表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询wlDepartment表"
// @Success 200 {object} response.Response{data=wl_playform.WlDepartment,msg=string} "查询成功"
// @Router /wlDepartment/findWlDepartment [get]
func (wlDepartmentApi *WlDepartmentApi) FindWlDepartment(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ID := c.Query("ID")
	rewlDepartment, err := wlDepartmentService.GetWlDepartment(ctx, ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(rewlDepartment, c)
}

// GetWlDepartmentList 分页获取wlDepartment表列表
// @Tags WlDepartment
// @Summary 分页获取wlDepartment表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query wl_playformReq.WlDepartmentSearch true "分页获取wlDepartment表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /wlDepartment/getWlDepartmentList [get]
func (wlDepartmentApi *WlDepartmentApi) GetWlDepartmentList(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var pageInfo wl_playformReq.WlDepartmentSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := wlDepartmentService.GetWlDepartmentInfoList(ctx, pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败:"+err.Error(), c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}

// GetWlDepartmentPublic 不需要鉴权的wlDepartment表接口
// @Tags WlDepartment
// @Summary 不需要鉴权的wlDepartment表接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /wlDepartment/getWlDepartmentPublic [get]
func (wlDepartmentApi *WlDepartmentApi) GetWlDepartmentPublic(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	wlDepartmentService.GetWlDepartmentPublic(ctx)
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的wlDepartment表接口信息",
	}, "获取成功", c)
}

// 分配设备到部门
func (wlDepartmentApi *WlDepartmentApi) AssignDevicesToDepartment(c *gin.Context) {
	var req wl_playformReq.AssignDevicesRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage("参数错误: "+err.Error(), c)
		return
	}
	err := wlDepartmentService.AssignDevicesToDepartment(c.Request.Context(), req.DepartmentID, req.DeviceIDs)
	if err != nil {
		response.FailWithMessage("分配设备失败: "+err.Error(), c)
		return
	}
	response.OkWithMessage("分配设备成功", c)
}

// 查询部门下所有设备
func (wlDepartmentApi *WlDepartmentApi) GetDevicesByDepartment(c *gin.Context) {
	var req wl_playformReq.DepartmentDevicesRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage("参数错误: "+err.Error(), c)
		return
	}
	devices, err := wlDepartmentService.GetDevicesByDepartment(c.Request.Context(), req.DepartmentID)
	if err != nil {
		response.FailWithMessage("查询设备失败: "+err.Error(), c)
		return
	}
	response.OkWithData(devices, c)
}

// 移除部门下某个设备
func (wlDepartmentApi *WlDepartmentApi) RemoveDeviceFromDepartment(c *gin.Context) {
	var req struct {
		DepartmentID int `json:"departmentId" binding:"required"`
		DeviceID     int `json:"deviceId" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage("参数错误: "+err.Error(), c)
		return
	}
	err := wlDepartmentService.RemoveDeviceFromDepartment(c.Request.Context(), req.DepartmentID, req.DeviceID)
	if err != nil {
		response.FailWithMessage("移除设备失败: "+err.Error(), c)
		return
	}
	response.OkWithMessage("移除设备成功", c)
}
