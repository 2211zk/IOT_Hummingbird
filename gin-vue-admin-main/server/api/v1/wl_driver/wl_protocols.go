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

type WlProtocolsApi struct{}

// CreateWlProtocols 创建wlProtocols表
func (api *WlProtocolsApi) CreateWlProtocols(c *gin.Context) {
	ctx := c.Request.Context()
	var data wl_driver.WlProtocols
	if err := c.ShouldBindJSON(&data); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if data.CreatedBy == nil {
		uid := utils.GetUserID(c)
		uid64 := uint64(uid)
		data.CreatedBy = &uid64
	}
	if err := wlProtocolsService.CreateWlProtocols(ctx, &data); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteWlProtocols 删除wlProtocols表
func (api *WlProtocolsApi) DeleteWlProtocols(c *gin.Context) {
	ctx := c.Request.Context()
	ID := c.Query("ID")
	userID := utils.GetUserID(c)
	if err := wlProtocolsService.DeleteWlProtocols(ctx, ID, userID); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteWlProtocolsByIds 批量删除wlProtocols表
func (api *WlProtocolsApi) DeleteWlProtocolsByIds(c *gin.Context) {
	ctx := c.Request.Context()
	IDs := c.QueryArray("IDs[]")
	userID := utils.GetUserID(c)
	if err := wlProtocolsService.DeleteWlProtocolsByIds(ctx, IDs, userID); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateWlProtocols 更新wlProtocols表
func (api *WlProtocolsApi) UpdateWlProtocols(c *gin.Context) {
	ctx := c.Request.Context()
	var data wl_driver.WlProtocols
	if err := c.ShouldBindJSON(&data); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	uid := utils.GetUserID(c)
	uid64 := uint64(uid)
	data.UpdatedBy = &uid64
	if err := wlProtocolsService.UpdateWlProtocols(ctx, data); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindWlProtocols 用id查询wlProtocols表
func (api *WlProtocolsApi) FindWlProtocols(c *gin.Context) {
	ctx := c.Request.Context()
	ID := c.Query("ID")
	result, err := wlProtocolsService.GetWlProtocols(ctx, ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(result, c)
}

// GetWlProtocolsList 分页获取wlProtocols表列表
func (api *WlProtocolsApi) GetWlProtocolsList(c *gin.Context) {
	ctx := c.Request.Context()
	var pageInfo wl_driverReq.WlProtocolsSearch
	if err := c.ShouldBindQuery(&pageInfo); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := wlProtocolsService.GetWlProtocolsInfoList(ctx, pageInfo)
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
