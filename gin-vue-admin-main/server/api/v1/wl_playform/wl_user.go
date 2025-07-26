package wl_playform

import (
	
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/model/wl_playform"
    wl_playformReq "github.com/flipped-aurora/gin-vue-admin/server/model/wl_playform/request"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type WlUserApi struct {}



// CreateWlUser 创建wlUser表
// @Tags WlUser
// @Summary 创建wlUser表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body wl_playform.WlUser true "创建wlUser表"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /wlUser/createWlUser [post]
func (wlUserApi *WlUserApi) CreateWlUser(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var wlUser wl_playform.WlUser
	err := c.ShouldBindJSON(&wlUser)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = wlUserService.CreateWlUser(ctx,&wlUser)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeleteWlUser 删除wlUser表
// @Tags WlUser
// @Summary 删除wlUser表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body wl_playform.WlUser true "删除wlUser表"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /wlUser/deleteWlUser [delete]
func (wlUserApi *WlUserApi) DeleteWlUser(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	id := c.Query("id")
	err := wlUserService.DeleteWlUser(ctx,id)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteWlUserByIds 批量删除wlUser表
// @Tags WlUser
// @Summary 批量删除wlUser表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /wlUser/deleteWlUserByIds [delete]
func (wlUserApi *WlUserApi) DeleteWlUserByIds(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ids := c.QueryArray("ids[]")
	err := wlUserService.DeleteWlUserByIds(ctx,ids)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateWlUser 更新wlUser表
// @Tags WlUser
// @Summary 更新wlUser表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body wl_playform.WlUser true "更新wlUser表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /wlUser/updateWlUser [put]
func (wlUserApi *WlUserApi) UpdateWlUser(c *gin.Context) {
    // 从ctx获取标准context进行业务行为
    ctx := c.Request.Context()

	var wlUser wl_playform.WlUser
	err := c.ShouldBindJSON(&wlUser)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = wlUserService.UpdateWlUser(ctx,wlUser)
	if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindWlUser 用id查询wlUser表
// @Tags WlUser
// @Summary 用id查询wlUser表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param id query int true "用id查询wlUser表"
// @Success 200 {object} response.Response{data=wl_playform.WlUser,msg=string} "查询成功"
// @Router /wlUser/findWlUser [get]
func (wlUserApi *WlUserApi) FindWlUser(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	id := c.Query("id")
	rewlUser, err := wlUserService.GetWlUser(ctx,id)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
		return
	}
	response.OkWithData(rewlUser, c)
}
// GetWlUserList 分页获取wlUser表列表
// @Tags WlUser
// @Summary 分页获取wlUser表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query wl_playformReq.WlUserSearch true "分页获取wlUser表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /wlUser/getWlUserList [get]
func (wlUserApi *WlUserApi) GetWlUserList(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var pageInfo wl_playformReq.WlUserSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := wlUserService.GetWlUserInfoList(ctx,pageInfo)
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

// GetWlUserPublic 不需要鉴权的wlUser表接口
// @Tags WlUser
// @Summary 不需要鉴权的wlUser表接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /wlUser/getWlUserPublic [get]
func (wlUserApi *WlUserApi) GetWlUserPublic(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    wlUserService.GetWlUserPublic(ctx)
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的wlUser表接口信息",
    }, "获取成功", c)
}
