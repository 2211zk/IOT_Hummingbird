package wl_driver

import (
	"strconv"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/wl_driver"
	wl_driverReq "github.com/flipped-aurora/gin-vue-admin/server/model/wl_driver/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type DriverCardsApi struct{}

func (driverCardsApi *DriverCardsApi) CreateDriverCards(c *gin.Context) {
	ctx := c.Request.Context()
	var driverCards wl_driver.DriverCards
	err := c.ShouldBindJSON(&driverCards)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = service.ServiceGroupApp.Wl_driverServiceGroup.DriverCardsService.CreateDriverCards(ctx, &driverCards)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

func (driverCardsApi *DriverCardsApi) DeleteDriverCards(c *gin.Context) {
	ctx := c.Request.Context()
	var driverCards wl_driver.DriverCards
	err := c.ShouldBindJSON(&driverCards)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	userID := uint(0) // 可根据实际获取用户ID
	err = service.ServiceGroupApp.Wl_driverServiceGroup.DriverCardsService.DeleteDriverCards(ctx, driverCards.ID, userID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

func (driverCardsApi *DriverCardsApi) DeleteDriverCardsByIds(c *gin.Context) {
	ctx := c.Request.Context()
	var idsReq request.IdsReq
	err := c.ShouldBindJSON(&idsReq)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	userID := uint(0) // 可根据实际获取用户ID
	// 类型转换 []int -> []uint
	ids := make([]uint, len(idsReq.Ids))
	for i, v := range idsReq.Ids {
		ids[i] = uint(v)
	}
	err = service.ServiceGroupApp.Wl_driverServiceGroup.DriverCardsService.DeleteDriverCardsByIds(ctx, ids, userID)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

func (driverCardsApi *DriverCardsApi) UpdateDriverCards(c *gin.Context) {
	ctx := c.Request.Context()
	var driverCards wl_driver.DriverCards
	err := c.ShouldBindJSON(&driverCards)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = service.ServiceGroupApp.Wl_driverServiceGroup.DriverCardsService.UpdateDriverCards(ctx, driverCards)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

func (driverCardsApi *DriverCardsApi) FindDriverCards(c *gin.Context) {
	ctx := c.Request.Context()
	id, err := parseUintQuery(c, "ID")
	if err != nil {
		response.FailWithMessage("ID参数错误", c)
		return
	}
	driverCards, err := service.ServiceGroupApp.Wl_driverServiceGroup.DriverCardsService.GetDriverCards(ctx, id)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(driverCards, c)
}

func (driverCardsApi *DriverCardsApi) GetDriverCardsList(c *gin.Context) {
	ctx := c.Request.Context()
	var pageInfo wl_driverReq.DriverCardsSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := service.ServiceGroupApp.Wl_driverServiceGroup.DriverCardsService.GetDriverCardsInfoList(ctx, pageInfo)
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

func (driverCardsApi *DriverCardsApi) GetDriverCardsPublic(c *gin.Context) {
	ctx := c.Request.Context()
	service.ServiceGroupApp.Wl_driverServiceGroup.DriverCardsService.GetDriverCardsPublic(ctx)
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的driverCards表接口信息",
	}, "获取成功", c)
}

// parseUintQuery 辅助函数
func parseUintQuery(c *gin.Context, key string) (uint, error) {
	idStr := c.Query(key)
	if idStr == "" {
		return 0, nil
	}
	var id uint64
	var err error
	id, err = strconv.ParseUint(idStr, 10, 64)
	return uint(id), err
}
