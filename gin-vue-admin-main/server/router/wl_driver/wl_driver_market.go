package wl_driver

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type WlDriverMarketRouter struct{}

func (s *WlDriverMarketRouter) InitWlDriverMarketRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	wlDriverMarketRouter := Router.Group("wlDriverMarket").Use(middleware.OperationRecord())
	wlDriverMarketRouterWithoutRecord := Router.Group("wlDriverMarket")
	{
		wlDriverMarketRouter.POST("createWlDriverMarket", wlDriverMarketApi.CreateWlDriverMarket)
		wlDriverMarketRouter.DELETE("deleteWlDriverMarket", wlDriverMarketApi.DeleteWlDriverMarket)
		wlDriverMarketRouter.DELETE("deleteWlDriverMarketByIds", wlDriverMarketApi.DeleteWlDriverMarketByIds)
		wlDriverMarketRouter.PUT("updateWlDriverMarket", wlDriverMarketApi.UpdateWlDriverMarket)
	}
	{
		wlDriverMarketRouterWithoutRecord.GET("findWlDriverMarket", wlDriverMarketApi.FindWlDriverMarket)
		wlDriverMarketRouterWithoutRecord.GET("getWlDriverMarketList", wlDriverMarketApi.GetWlDriverMarketList)
	}
}
