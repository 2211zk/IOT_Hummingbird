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

type WlAlarmApi struct{}

// CreateWlAlarm 创建wlAlarm表
// @Tags WlAlarm
// @Summary 创建wlAlarm表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body wl_playform.WlAlarm true "创建wlAlarm表"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /wlAlarm/createWlAlarm [post]
func (wlAlarmApi *WlAlarmApi) CreateWlAlarm(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var wlAlarm wl_playform.WlAlarm
	err := c.ShouldBindJSON(&wlAlarm)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	uid := int(utils.GetUserID(c))
	wlAlarm.CreatedBy = &uid
	err = wlAlarmService.CreateWlAlarm(ctx, &wlAlarm)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteWlAlarm 删除wlAlarm表
// @Tags WlAlarm
// @Summary 删除wlAlarm表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body wl_playform.WlAlarm true "删除wlAlarm表"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /wlAlarm/deleteWlAlarm [delete]
func (wlAlarmApi *WlAlarmApi) DeleteWlAlarm(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ID := c.Query("ID")
	userID := utils.GetUserID(c)
	err := wlAlarmService.DeleteWlAlarm(ctx, ID, userID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteWlAlarmByIds 批量删除wlAlarm表
// @Tags WlAlarm
// @Summary 批量删除wlAlarm表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /wlAlarm/deleteWlAlarmByIds [delete]
func (wlAlarmApi *WlAlarmApi) DeleteWlAlarmByIds(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
	userID := utils.GetUserID(c)
	err := wlAlarmService.DeleteWlAlarmByIds(ctx, IDs, userID)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateWlAlarm 更新wlAlarm表
// @Tags WlAlarm
// @Summary 更新wlAlarm表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body wl_playform.WlAlarm true "更新wlAlarm表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /wlAlarm/updateWlAlarm [put]
func (wlAlarmApi *WlAlarmApi) UpdateWlAlarm(c *gin.Context) {
	// 从ctx获取标准context进行业务行为
	ctx := c.Request.Context()

	var wlAlarm wl_playform.WlAlarm
	err := c.ShouldBindJSON(&wlAlarm)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	uid2 := int(utils.GetUserID(c))
	wlAlarm.UpdatedBy = &uid2
	err = wlAlarmService.UpdateWlAlarm(ctx, wlAlarm)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindWlAlarm 用id查询wlAlarm表
// @Tags WlAlarm
// @Summary 用id查询wlAlarm表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询wlAlarm表"
// @Success 200 {object} response.Response{data=wl_playform.WlAlarm,msg=string} "查询成功"
// @Router /wlAlarm/findWlAlarm [get]
func (wlAlarmApi *WlAlarmApi) FindWlAlarm(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ID := c.Query("ID")
	rewlAlarm, err := wlAlarmService.GetWlAlarm(ctx, ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(rewlAlarm, c)
}

// GetWlAlarmList 分页获取wlAlarm表列表
// @Tags WlAlarm
// @Summary 分页获取wlAlarm表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query wl_playformReq.WlAlarmSearch true "分页获取wlAlarm表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /wlAlarm/getWlAlarmList [get]
func (wlAlarmApi *WlAlarmApi) GetWlAlarmList(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var pageInfo wl_playformReq.WlAlarmSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := wlAlarmService.GetWlAlarmInfoList(ctx, pageInfo)
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

// GetWlAlarmPublic 不需要鉴权的wlAlarm表接口
// @Tags WlAlarm
// @Summary 不需要鉴权的wlAlarm表接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /wlAlarm/getWlAlarmPublic [get]
func (wlAlarmApi *WlAlarmApi) GetWlAlarmPublic(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	wlAlarmService.GetWlAlarmPublic(ctx)
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的wlAlarm表接口信息",
	}, "获取成功", c)
}
