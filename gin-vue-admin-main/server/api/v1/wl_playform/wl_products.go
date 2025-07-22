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

type WlProductsApi struct {}



// CreateWlProducts 创建wlProducts表
// @Tags WlProducts
// @Summary 创建wlProducts表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body wl_playform.WlProducts true "创建wlProducts表"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /wlProducts/createWlProducts [post]
func (wlProductsApi *WlProductsApi) CreateWlProducts(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var wlProducts wl_playform.WlProducts
	err := c.ShouldBindJSON(&wlProducts)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    wlProducts.CreatedBy = utils.GetUserID(c)
	err = wlProductsService.CreateWlProducts(ctx,&wlProducts)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeleteWlProducts 删除wlProducts表
// @Tags WlProducts
// @Summary 删除wlProducts表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body wl_playform.WlProducts true "删除wlProducts表"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /wlProducts/deleteWlProducts [delete]
func (wlProductsApi *WlProductsApi) DeleteWlProducts(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
    userID := utils.GetUserID(c)
	err := wlProductsService.DeleteWlProducts(ctx,ID,userID)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteWlProductsByIds 批量删除wlProducts表
// @Tags WlProducts
// @Summary 批量删除wlProducts表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /wlProducts/deleteWlProductsByIds [delete]
func (wlProductsApi *WlProductsApi) DeleteWlProductsByIds(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
    userID := utils.GetUserID(c)
	err := wlProductsService.DeleteWlProductsByIds(ctx,IDs,userID)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateWlProducts 更新wlProducts表
// @Tags WlProducts
// @Summary 更新wlProducts表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body wl_playform.WlProducts true "更新wlProducts表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /wlProducts/updateWlProducts [put]
func (wlProductsApi *WlProductsApi) UpdateWlProducts(c *gin.Context) {
    // 从ctx获取标准context进行业务行为
    ctx := c.Request.Context()

	var wlProducts wl_playform.WlProducts
	err := c.ShouldBindJSON(&wlProducts)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    wlProducts.UpdatedBy = utils.GetUserID(c)
	err = wlProductsService.UpdateWlProducts(ctx,wlProducts)
	if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindWlProducts 用id查询wlProducts表
// @Tags WlProducts
// @Summary 用id查询wlProducts表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询wlProducts表"
// @Success 200 {object} response.Response{data=wl_playform.WlProducts,msg=string} "查询成功"
// @Router /wlProducts/findWlProducts [get]
func (wlProductsApi *WlProductsApi) FindWlProducts(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
	rewlProducts, err := wlProductsService.GetWlProducts(ctx,ID)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
		return
	}
	response.OkWithData(rewlProducts, c)
}
// GetWlProductsList 分页获取wlProducts表列表
// @Tags WlProducts
// @Summary 分页获取wlProducts表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query wl_playformReq.WlProductsSearch true "分页获取wlProducts表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /wlProducts/getWlProductsList [get]
func (wlProductsApi *WlProductsApi) GetWlProductsList(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var pageInfo wl_playformReq.WlProductsSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := wlProductsService.GetWlProductsInfoList(ctx,pageInfo)
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

// GetWlProductsPublic 不需要鉴权的wlProducts表接口
// @Tags WlProducts
// @Summary 不需要鉴权的wlProducts表接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /wlProducts/getWlProductsPublic [get]
func (wlProductsApi *WlProductsApi) GetWlProductsPublic(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    wlProductsService.GetWlProductsPublic(ctx)
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的wlProducts表接口信息",
    }, "获取成功", c)
}
