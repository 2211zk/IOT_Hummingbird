package wl_driver

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/wl_driver"
	wl_driverReq "github.com/flipped-aurora/gin-vue-admin/server/model/wl_driver/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type WlDriverMarketApi struct{}

// CreateWlDriverMarket 创建wlDriverMarket表
func (api *WlDriverMarketApi) CreateWlDriverMarket(c *gin.Context) {
	ctx := c.Request.Context()
	var data wl_driver.WlDriverMarket
	if err := c.ShouldBindJSON(&data); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := wlDriverMarketService.CreateWlDriverMarket(ctx, &data); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteWlDriverMarket 删除wlDriverMarket表
func (api *WlDriverMarketApi) DeleteWlDriverMarket(c *gin.Context) {
	ctx := c.Request.Context()
	marketId := c.Query("market_id")
	if err := wlDriverMarketService.DeleteWlDriverMarket(ctx, marketId); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteWlDriverMarketByIds 批量删除wlDriverMarket表
func (api *WlDriverMarketApi) DeleteWlDriverMarketByIds(c *gin.Context) {
	ctx := c.Request.Context()
	marketIds := c.QueryArray("market_ids[]")
	if err := wlDriverMarketService.DeleteWlDriverMarketByIds(ctx, marketIds); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateWlDriverMarket 更新wlDriverMarket表
func (api *WlDriverMarketApi) UpdateWlDriverMarket(c *gin.Context) {
	ctx := c.Request.Context()
	var data wl_driver.WlDriverMarket
	if err := c.ShouldBindJSON(&data); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := wlDriverMarketService.UpdateWlDriverMarket(ctx, data); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindWlDriverMarket 用market_id查询wlDriverMarket表
func (api *WlDriverMarketApi) FindWlDriverMarket(c *gin.Context) {
	ctx := c.Request.Context()
	marketId := c.Query("market_id")
	result, err := wlDriverMarketService.GetWlDriverMarket(ctx, marketId)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(result, c)
}

// GetWlDriverMarketList 分页获取wlDriverMarket表列表
func (api *WlDriverMarketApi) GetWlDriverMarketList(c *gin.Context) {
	ctx := c.Request.Context()
	var pageInfo wl_driverReq.WlDriverMarketSearch
	if err := c.ShouldBindQuery(&pageInfo); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := wlDriverMarketService.GetWlDriverMarketInfoList(ctx, pageInfo)
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
