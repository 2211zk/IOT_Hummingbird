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

type WlScenesApi struct {}



// CreateWlScenes 创建wlScenes表
// @Tags WlScenes
// @Summary 创建wlScenes表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body wl_playform.WlScenes true "创建wlScenes表"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /wlScenes/createWlScenes [post]
func (wlScenesApi *WlScenesApi) CreateWlScenes(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var wlScenes wl_playform.WlScenes
	err := c.ShouldBindJSON(&wlScenes)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    wlScenes.CreatedBy = utils.GetUserID(c)
	err = wlScenesService.CreateWlScenes(ctx,&wlScenes)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeleteWlScenes 删除wlScenes表
// @Tags WlScenes
// @Summary 删除wlScenes表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body wl_playform.WlScenes true "删除wlScenes表"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /wlScenes/deleteWlScenes [delete]
func (wlScenesApi *WlScenesApi) DeleteWlScenes(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
    userID := utils.GetUserID(c)
	err := wlScenesService.DeleteWlScenes(ctx,ID,userID)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteWlScenesByIds 批量删除wlScenes表
// @Tags WlScenes
// @Summary 批量删除wlScenes表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /wlScenes/deleteWlScenesByIds [delete]
func (wlScenesApi *WlScenesApi) DeleteWlScenesByIds(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
    userID := utils.GetUserID(c)
	err := wlScenesService.DeleteWlScenesByIds(ctx,IDs,userID)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateWlScenes 更新wlScenes表
// @Tags WlScenes
// @Summary 更新wlScenes表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body wl_playform.WlScenes true "更新wlScenes表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /wlScenes/updateWlScenes [put]
func (wlScenesApi *WlScenesApi) UpdateWlScenes(c *gin.Context) {
    // 从ctx获取标准context进行业务行为
    ctx := c.Request.Context()

	var wlScenes wl_playform.WlScenes
	err := c.ShouldBindJSON(&wlScenes)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    wlScenes.UpdatedBy = utils.GetUserID(c)
	err = wlScenesService.UpdateWlScenes(ctx,wlScenes)
	if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindWlScenes 用id查询wlScenes表
// @Tags WlScenes
// @Summary 用id查询wlScenes表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询wlScenes表"
// @Success 200 {object} response.Response{data=wl_playform.WlScenes,msg=string} "查询成功"
// @Router /wlScenes/findWlScenes [get]
func (wlScenesApi *WlScenesApi) FindWlScenes(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
	rewlScenes, err := wlScenesService.GetWlScenes(ctx,ID)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
		return
	}
	response.OkWithData(rewlScenes, c)
}
// GetWlScenesList 分页获取wlScenes表列表
// @Tags WlScenes
// @Summary 分页获取wlScenes表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query wl_playformReq.WlScenesSearch true "分页获取wlScenes表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /wlScenes/getWlScenesList [get]
func (wlScenesApi *WlScenesApi) GetWlScenesList(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var pageInfo wl_playformReq.WlScenesSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := wlScenesService.GetWlScenesInfoList(ctx,pageInfo)
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

// GetWlScenesPublic 不需要鉴权的wlScenes表接口
// @Tags WlScenes
// @Summary 不需要鉴权的wlScenes表接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /wlScenes/getWlScenesPublic [get]
func (wlScenesApi *WlScenesApi) GetWlScenesPublic(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    wlScenesService.GetWlScenesPublic(ctx)
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的wlScenes表接口信息",
    }, "获取成功", c)
}
