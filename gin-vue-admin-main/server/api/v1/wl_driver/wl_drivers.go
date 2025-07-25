package wl_driver

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/wl_driver"
	wl_driverReq "github.com/flipped-aurora/gin-vue-admin/server/model/wl_driver/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type WlDriversApi struct{}

// CreateWlDrivers 创建wlDrivers表
// @Tags WlDrivers
// @Summary 创建wlDrivers表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body wl_driver.WlDrivers true "创建wlDrivers表"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /wlDrivers/createWlDrivers [post]
func (wlDriversApi *WlDriversApi) CreateWlDrivers(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var wlDrivers wl_driver.WlDrivers
	err := c.ShouldBindJSON(&wlDrivers)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	// 新增必填校验
	if err := utils.Verify(wlDrivers, utils.WlDriversVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	userID := int(utils.GetUserID(c))
	wlDrivers.CreatedBy = &userID
	err = wlDriversService.CreateWlDrivers(ctx, &wlDrivers)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteWlDrivers 删除wlDrivers表
// @Tags WlDrivers
// @Summary 删除wlDrivers表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body wl_driver.WlDrivers true "删除wlDrivers表"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /wlDrivers/deleteWlDrivers [delete]
func (wlDriversApi *WlDriversApi) DeleteWlDrivers(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ID := c.Query("ID")
	userID := utils.GetUserID(c)
	err := wlDriversService.DeleteWlDrivers(ctx, ID, userID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteWlDriversByIds 批量删除wlDrivers表
// @Tags WlDrivers
// @Summary 批量删除wlDrivers表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /wlDrivers/deleteWlDriversByIds [delete]
func (wlDriversApi *WlDriversApi) DeleteWlDriversByIds(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
	userID := utils.GetUserID(c)
	err := wlDriversService.DeleteWlDriversByIds(ctx, IDs, userID)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateWlDrivers 更新wlDrivers表
// @Tags WlDrivers
// @Summary 更新wlDrivers表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body wl_driver.WlDrivers true "更新wlDrivers表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /wlDrivers/updateWlDrivers [put]
func (wlDriversApi *WlDriversApi) UpdateWlDrivers(c *gin.Context) {
	// 从ctx获取标准context进行业务行为
	ctx := c.Request.Context()

	var wlDrivers wl_driver.WlDrivers
	err := c.ShouldBindJSON(&wlDrivers)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	userID := int(utils.GetUserID(c))
	wlDrivers.UpdatedBy = &userID
	err = wlDriversService.UpdateWlDrivers(ctx, wlDrivers)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindWlDrivers 用id查询wlDrivers表
// @Tags WlDrivers
// @Summary 用id查询wlDrivers表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询wlDrivers表"
// @Success 200 {object} response.Response{data=wl_driver.WlDrivers,msg=string} "查询成功"
// @Router /wlDrivers/findWlDrivers [get]
func (wlDriversApi *WlDriversApi) FindWlDrivers(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ID := c.Query("ID")
	rewlDrivers, err := wlDriversService.GetWlDrivers(ctx, ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(rewlDrivers, c)
}

// GetWlDriversList 分页获取wlDrivers表列表
// @Tags WlDrivers
// @Summary 分页获取wlDrivers表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query wl_driverReq.WlDriversSearch true "分页获取wlDrivers表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /wlDrivers/getWlDriversList [get]
func (wlDriversApi *WlDriversApi) GetWlDriversList(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var pageInfo wl_driverReq.WlDriversSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := wlDriversService.GetWlDriversInfoList(ctx, pageInfo)
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

// GetWlDriversPublic 不需要鉴权的wlDrivers表接口
// @Tags WlDrivers
// @Summary 不需要鉴权的wlDrivers表接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /wlDrivers/getWlDriversPublic [get]
func (wlDriversApi *WlDriversApi) GetWlDriversPublic(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	wlDriversService.GetWlDriversPublic(ctx)
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的wlDrivers表接口信息",
	}, "获取成功", c)
}

// // 获取驱动图片（可根据类型自定义图片路径）
// const getDriverImg = (item) => {
//   if (item.driverType === '官方') {
//     return require('@/assets/driver-official.png')
//   } else if (item.driverType === '自定义') {
//     return require('@/assets/driver-custom.png')
//   }
//   return require('@/assets/driver-default.png')
// }
