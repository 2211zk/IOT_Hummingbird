package wl_driver

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/wl_driver"
	wl_driverReq "github.com/flipped-aurora/gin-vue-admin/server/model/wl_driver/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type WlSystemMoniorApi struct{}

// CreateWlSystemMonior 创建wlSystemMonior表
func (api *WlSystemMoniorApi) CreateWlSystemMonior(c *gin.Context) {
	ctx := c.Request.Context()
	var data wl_driver.WlSystemMonior
	if err := c.ShouldBindJSON(&data); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := wlSystemMoniorService.CreateWlSystemMonior(ctx, &data); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteWlSystemMonior 删除wlSystemMonior表
func (api *WlSystemMoniorApi) DeleteWlSystemMonior(c *gin.Context) {
	ctx := c.Request.Context()
	ID := c.Query("ID")
	if err := wlSystemMoniorService.DeleteWlSystemMonior(ctx, ID); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteWlSystemMoniorByIds 批量删除wlSystemMonior表
func (api *WlSystemMoniorApi) DeleteWlSystemMoniorByIds(c *gin.Context) {
	ctx := c.Request.Context()
	IDs := c.QueryArray("IDs[]")
	if err := wlSystemMoniorService.DeleteWlSystemMoniorByIds(ctx, IDs); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateWlSystemMonior 更新wlSystemMonior表
func (api *WlSystemMoniorApi) UpdateWlSystemMonior(c *gin.Context) {
	ctx := c.Request.Context()
	var data wl_driver.WlSystemMonior
	if err := c.ShouldBindJSON(&data); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := wlSystemMoniorService.UpdateWlSystemMonior(ctx, data); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindWlSystemMonior 用id查询wlSystemMonior表
func (api *WlSystemMoniorApi) FindWlSystemMonior(c *gin.Context) {
	ctx := c.Request.Context()
	ID := c.Query("ID")
	result, err := wlSystemMoniorService.GetWlSystemMonior(ctx, ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(result, c)
}

// GetWlSystemMoniorList 分页获取wlSystemMonior表列表
func (api *WlSystemMoniorApi) GetWlSystemMoniorList(c *gin.Context) {
	ctx := c.Request.Context()
	var pageInfo wl_driverReq.WlSystemMoniorSearch
	if err := c.ShouldBindQuery(&pageInfo); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := wlSystemMoniorService.GetWlSystemMoniorInfoList(ctx, pageInfo)
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
