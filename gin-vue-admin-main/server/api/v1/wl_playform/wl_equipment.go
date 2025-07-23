package wl_playform

import (
	
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/model/wl_playform"
    wl_playformReq "github.com/flipped-aurora/gin-vue-admin/server/model/wl_playform/request"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
    "github.com/flipped-aurora/gin-vue-admin/server/utils"
)

type WlEquipmentApi struct {}



// CreateWlEquipment 创建wlEquipment表
// @Tags WlEquipment
// @Summary 创建wlEquipment表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body wl_playform.WlEquipment true "创建wlEquipment表"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /wlEquipment/createWlEquipment [post]
func (wlEquipmentApi *WlEquipmentApi) CreateWlEquipment(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var wlEquipment wl_playform.WlEquipment
	err := c.ShouldBindJSON(&wlEquipment)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    wlEquipment.CreatedBy = utils.GetUserID(c)
	err = wlEquipmentService.CreateWlEquipment(ctx,&wlEquipment)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeleteWlEquipment 删除wlEquipment表
// @Tags WlEquipment
// @Summary 删除wlEquipment表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body wl_playform.WlEquipment true "删除wlEquipment表"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /wlEquipment/deleteWlEquipment [delete]
func (wlEquipmentApi *WlEquipmentApi) DeleteWlEquipment(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
    userID := utils.GetUserID(c)
	err := wlEquipmentService.DeleteWlEquipment(ctx,ID,userID)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteWlEquipmentByIds 批量删除wlEquipment表
// @Tags WlEquipment
// @Summary 批量删除wlEquipment表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /wlEquipment/deleteWlEquipmentByIds [delete]
func (wlEquipmentApi *WlEquipmentApi) DeleteWlEquipmentByIds(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
    userID := utils.GetUserID(c)
	err := wlEquipmentService.DeleteWlEquipmentByIds(ctx,IDs,userID)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateWlEquipment 更新wlEquipment表
// @Tags WlEquipment
// @Summary 更新wlEquipment表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body wl_playform.WlEquipment true "更新wlEquipment表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /wlEquipment/updateWlEquipment [put]
func (wlEquipmentApi *WlEquipmentApi) UpdateWlEquipment(c *gin.Context) {
    // 从ctx获取标准context进行业务行为
    ctx := c.Request.Context()

	var wlEquipment wl_playform.WlEquipment
	err := c.ShouldBindJSON(&wlEquipment)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    wlEquipment.UpdatedBy = utils.GetUserID(c)
	err = wlEquipmentService.UpdateWlEquipment(ctx,wlEquipment)
	if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindWlEquipment 用id查询wlEquipment表
// @Tags WlEquipment
// @Summary 用id查询wlEquipment表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询wlEquipment表"
// @Success 200 {object} response.Response{data=wl_playform.WlEquipment,msg=string} "查询成功"
// @Router /wlEquipment/findWlEquipment [get]
func (wlEquipmentApi *WlEquipmentApi) FindWlEquipment(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
	rewlEquipment, err := wlEquipmentService.GetWlEquipment(ctx,ID)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
		return
	}
	response.OkWithData(rewlEquipment, c)
}
// GetWlEquipmentList 分页获取wlEquipment表列表
// @Tags WlEquipment
// @Summary 分页获取wlEquipment表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query wl_playformReq.WlEquipmentSearch true "分页获取wlEquipment表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /wlEquipment/getWlEquipmentList [get]
func (wlEquipmentApi *WlEquipmentApi) GetWlEquipmentList(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var pageInfo wl_playformReq.WlEquipmentSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := wlEquipmentService.GetWlEquipmentInfoList(ctx,pageInfo)
	if err != nil {
	    global.GVA_LOG.Error("获取失败!", zap.Error(err))
        response.FailWithMessage("获取失败:" + err.Error(), c)
        return
    }
    response.OkWithDetailed(response.PageResult{
        List:     list,
        Total:    total,
        Page:     pageInfo.Page,
        PageSize: pageInfo.PageSize,
    }, "获取成功", c)
}

// GetWlEquipmentPublic 不需要鉴权的wlEquipment表接口
// @Tags WlEquipment
// @Summary 不需要鉴权的wlEquipment表接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /wlEquipment/getWlEquipmentPublic [get]
func (wlEquipmentApi *WlEquipmentApi) GetWlEquipmentPublic(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    wlEquipmentService.GetWlEquipmentPublic(ctx)
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的wlEquipment表接口信息",
    }, "获取成功", c)
}
