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

type WlCategoryApi struct {}



// CreateWlCategory 创建wlCategory表
// @Tags WlCategory
// @Summary 创建wlCategory表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body wl_playform.WlCategory true "创建wlCategory表"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /wlCategory/createWlCategory [post]
func (wlCategoryApi *WlCategoryApi) CreateWlCategory(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var wlCategory wl_playform.WlCategory
	err := c.ShouldBindJSON(&wlCategory)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    wlCategory.CreatedBy = utils.GetUserID(c)
	err = wlCategoryService.CreateWlCategory(ctx,&wlCategory)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeleteWlCategory 删除wlCategory表
// @Tags WlCategory
// @Summary 删除wlCategory表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body wl_playform.WlCategory true "删除wlCategory表"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /wlCategory/deleteWlCategory [delete]
func (wlCategoryApi *WlCategoryApi) DeleteWlCategory(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
    userID := utils.GetUserID(c)
	err := wlCategoryService.DeleteWlCategory(ctx,ID,userID)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteWlCategoryByIds 批量删除wlCategory表
// @Tags WlCategory
// @Summary 批量删除wlCategory表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /wlCategory/deleteWlCategoryByIds [delete]
func (wlCategoryApi *WlCategoryApi) DeleteWlCategoryByIds(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
    userID := utils.GetUserID(c)
	err := wlCategoryService.DeleteWlCategoryByIds(ctx,IDs,userID)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateWlCategory 更新wlCategory表
// @Tags WlCategory
// @Summary 更新wlCategory表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body wl_playform.WlCategory true "更新wlCategory表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /wlCategory/updateWlCategory [put]
func (wlCategoryApi *WlCategoryApi) UpdateWlCategory(c *gin.Context) {
    // 从ctx获取标准context进行业务行为
    ctx := c.Request.Context()

	var wlCategory wl_playform.WlCategory
	err := c.ShouldBindJSON(&wlCategory)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    wlCategory.UpdatedBy = utils.GetUserID(c)
	err = wlCategoryService.UpdateWlCategory(ctx,wlCategory)
	if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindWlCategory 用id查询wlCategory表
// @Tags WlCategory
// @Summary 用id查询wlCategory表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询wlCategory表"
// @Success 200 {object} response.Response{data=wl_playform.WlCategory,msg=string} "查询成功"
// @Router /wlCategory/findWlCategory [get]
func (wlCategoryApi *WlCategoryApi) FindWlCategory(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
	rewlCategory, err := wlCategoryService.GetWlCategory(ctx,ID)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
		return
	}
	response.OkWithData(rewlCategory, c)
}
// GetWlCategoryList 分页获取wlCategory表列表
// @Tags WlCategory
// @Summary 分页获取wlCategory表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query wl_playformReq.WlCategorySearch true "分页获取wlCategory表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /wlCategory/getWlCategoryList [get]
func (wlCategoryApi *WlCategoryApi) GetWlCategoryList(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var pageInfo wl_playformReq.WlCategorySearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := wlCategoryService.GetWlCategoryInfoList(ctx,pageInfo)
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

// GetWlCategoryPublic 不需要鉴权的wlCategory表接口
// @Tags WlCategory
// @Summary 不需要鉴权的wlCategory表接口
// @Accept application/json
// @Produce application/json
// @Param data query wl_playformReq.WlCategorySearch true "分页获取wlCategory表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /wlCategory/getWlCategoryPublic [get]
func (wlCategoryApi *WlCategoryApi) GetWlCategoryPublic(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

    // 此接口不需要鉴权，返回真实的品类数据
    var pageInfo wl_playformReq.WlCategorySearch
    err := c.ShouldBindQuery(&pageInfo)
    if err != nil {
        response.FailWithMessage(err.Error(), c)
        return
    }
    
    list, total, err := wlCategoryService.GetWlCategoryInfoList(ctx, pageInfo)
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
