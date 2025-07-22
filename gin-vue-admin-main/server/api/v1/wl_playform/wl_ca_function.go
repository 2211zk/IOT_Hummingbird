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

type WlCaFunctionApi struct {}



// CreateWlCaFunction 创建wlCaFunction表
// @Tags WlCaFunction
// @Summary 创建wlCaFunction表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body wl_playform.WlCaFunction true "创建wlCaFunction表"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /wlCaFunction/createWlCaFunction [post]
func (wlCaFunctionApi *WlCaFunctionApi) CreateWlCaFunction(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var wlCaFunction wl_playform.WlCaFunction
	err := c.ShouldBindJSON(&wlCaFunction)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    wlCaFunction.CreatedBy = utils.GetUserID(c)
	err = wlCaFunctionService.CreateWlCaFunction(ctx,&wlCaFunction)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeleteWlCaFunction 删除wlCaFunction表
// @Tags WlCaFunction
// @Summary 删除wlCaFunction表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body wl_playform.WlCaFunction true "删除wlCaFunction表"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /wlCaFunction/deleteWlCaFunction [delete]
func (wlCaFunctionApi *WlCaFunctionApi) DeleteWlCaFunction(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
    userID := utils.GetUserID(c)
	err := wlCaFunctionService.DeleteWlCaFunction(ctx,ID,userID)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteWlCaFunctionByIds 批量删除wlCaFunction表
// @Tags WlCaFunction
// @Summary 批量删除wlCaFunction表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /wlCaFunction/deleteWlCaFunctionByIds [delete]
func (wlCaFunctionApi *WlCaFunctionApi) DeleteWlCaFunctionByIds(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
    userID := utils.GetUserID(c)
	err := wlCaFunctionService.DeleteWlCaFunctionByIds(ctx,IDs,userID)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateWlCaFunction 更新wlCaFunction表
// @Tags WlCaFunction
// @Summary 更新wlCaFunction表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body wl_playform.WlCaFunction true "更新wlCaFunction表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /wlCaFunction/updateWlCaFunction [put]
func (wlCaFunctionApi *WlCaFunctionApi) UpdateWlCaFunction(c *gin.Context) {
    // 从ctx获取标准context进行业务行为
    ctx := c.Request.Context()

	var wlCaFunction wl_playform.WlCaFunction
	err := c.ShouldBindJSON(&wlCaFunction)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    wlCaFunction.UpdatedBy = utils.GetUserID(c)
	err = wlCaFunctionService.UpdateWlCaFunction(ctx,wlCaFunction)
	if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindWlCaFunction 用id查询wlCaFunction表
// @Tags WlCaFunction
// @Summary 用id查询wlCaFunction表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询wlCaFunction表"
// @Success 200 {object} response.Response{data=wl_playform.WlCaFunction,msg=string} "查询成功"
// @Router /wlCaFunction/findWlCaFunction [get]
func (wlCaFunctionApi *WlCaFunctionApi) FindWlCaFunction(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
	rewlCaFunction, err := wlCaFunctionService.GetWlCaFunction(ctx,ID)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
		return
	}
	response.OkWithData(rewlCaFunction, c)
}
// GetWlCaFunctionList 分页获取wlCaFunction表列表
// @Tags WlCaFunction
// @Summary 分页获取wlCaFunction表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query wl_playformReq.WlCaFunctionSearch true "分页获取wlCaFunction表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /wlCaFunction/getWlCaFunctionList [get]
func (wlCaFunctionApi *WlCaFunctionApi) GetWlCaFunctionList(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var pageInfo wl_playformReq.WlCaFunctionSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := wlCaFunctionService.GetWlCaFunctionInfoList(ctx,pageInfo)
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

// GetWlCaFunctionPublic 不需要鉴权的wlCaFunction表接口
// @Tags WlCaFunction
// @Summary 不需要鉴权的wlCaFunction表接口
// @Accept application/json
// @Produce application/json
// @Param data query wl_playformReq.WlCaFunctionSearch true "分页获取wlCaFunction表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /wlCaFunction/getWlCaFunctionPublic [get]
func (wlCaFunctionApi *WlCaFunctionApi) GetWlCaFunctionPublic(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

    // 此接口不需要鉴权，返回真实的功能定义数据
    var pageInfo wl_playformReq.WlCaFunctionSearch
    err := c.ShouldBindQuery(&pageInfo)
    if err != nil {
        response.FailWithMessage(err.Error(), c)
        return
    }
    
    list, total, err := wlCaFunctionService.GetWlCaFunctionInfoList(ctx, pageInfo)
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
