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

type WlEngineRulesApi struct {}



// CreateWlEngineRules 创建wlEngineRules表
// @Tags WlEngineRules
// @Summary 创建wlEngineRules表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body wl_playform.WlEngineRules true "创建wlEngineRules表"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /wlEngineRules/createWlEngineRules [post]
func (wlEngineRulesApi *WlEngineRulesApi) CreateWlEngineRules(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var wlEngineRules wl_playform.WlEngineRules
	err := c.ShouldBindJSON(&wlEngineRules)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    wlEngineRules.CreatedBy = utils.GetUserID(c)
	err = wlEngineRulesService.CreateWlEngineRules(ctx,&wlEngineRules)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeleteWlEngineRules 删除wlEngineRules表
// @Tags WlEngineRules
// @Summary 删除wlEngineRules表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body wl_playform.WlEngineRules true "删除wlEngineRules表"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /wlEngineRules/deleteWlEngineRules [delete]
func (wlEngineRulesApi *WlEngineRulesApi) DeleteWlEngineRules(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
    userID := utils.GetUserID(c)
	err := wlEngineRulesService.DeleteWlEngineRules(ctx,ID,userID)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteWlEngineRulesByIds 批量删除wlEngineRules表
// @Tags WlEngineRules
// @Summary 批量删除wlEngineRules表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /wlEngineRules/deleteWlEngineRulesByIds [delete]
func (wlEngineRulesApi *WlEngineRulesApi) DeleteWlEngineRulesByIds(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
    userID := utils.GetUserID(c)
	err := wlEngineRulesService.DeleteWlEngineRulesByIds(ctx,IDs,userID)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateWlEngineRules 更新wlEngineRules表
// @Tags WlEngineRules
// @Summary 更新wlEngineRules表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body wl_playform.WlEngineRules true "更新wlEngineRules表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /wlEngineRules/updateWlEngineRules [put]
func (wlEngineRulesApi *WlEngineRulesApi) UpdateWlEngineRules(c *gin.Context) {
    // 从ctx获取标准context进行业务行为
    ctx := c.Request.Context()

	var wlEngineRules wl_playform.WlEngineRules
	err := c.ShouldBindJSON(&wlEngineRules)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    wlEngineRules.UpdatedBy = utils.GetUserID(c)
	err = wlEngineRulesService.UpdateWlEngineRules(ctx,wlEngineRules)
	if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindWlEngineRules 用id查询wlEngineRules表
// @Tags WlEngineRules
// @Summary 用id查询wlEngineRules表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询wlEngineRules表"
// @Success 200 {object} response.Response{data=wl_playform.WlEngineRules,msg=string} "查询成功"
// @Router /wlEngineRules/findWlEngineRules [get]
func (wlEngineRulesApi *WlEngineRulesApi) FindWlEngineRules(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
	rewlEngineRules, err := wlEngineRulesService.GetWlEngineRules(ctx,ID)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
		return
	}
	response.OkWithData(rewlEngineRules, c)
}
// GetWlEngineRulesList 分页获取wlEngineRules表列表
// @Tags WlEngineRules
// @Summary 分页获取wlEngineRules表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query wl_playformReq.WlEngineRulesSearch true "分页获取wlEngineRules表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /wlEngineRules/getWlEngineRulesList [get]
func (wlEngineRulesApi *WlEngineRulesApi) GetWlEngineRulesList(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var pageInfo wl_playformReq.WlEngineRulesSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := wlEngineRulesService.GetWlEngineRulesInfoList(ctx,pageInfo)
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

// GetWlEngineRulesPublic 不需要鉴权的wlEngineRules表接口
// @Tags WlEngineRules
// @Summary 不需要鉴权的wlEngineRules表接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /wlEngineRules/getWlEngineRulesPublic [get]
func (wlEngineRulesApi *WlEngineRulesApi) GetWlEngineRulesPublic(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    wlEngineRulesService.GetWlEngineRulesPublic(ctx)
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的wlEngineRules表接口信息",
    }, "获取成功", c)
}
