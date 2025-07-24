package wl_driver

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type WlProtocolsRouter struct{}

func (s *WlProtocolsRouter) InitWlProtocolsRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	wlProtocolsRouter := Router.Group("wlProtocols").Use(middleware.OperationRecord())
	wlProtocolsRouterWithoutRecord := Router.Group("wlProtocols")
	{
		wlProtocolsRouter.POST("createWlProtocols", wlProtocolsApi.CreateWlProtocols)
		wlProtocolsRouter.DELETE("deleteWlProtocols", wlProtocolsApi.DeleteWlProtocols)
		wlProtocolsRouter.DELETE("deleteWlProtocolsByIds", wlProtocolsApi.DeleteWlProtocolsByIds)
		wlProtocolsRouter.PUT("updateWlProtocols", wlProtocolsApi.UpdateWlProtocols)
	}
	{
		wlProtocolsRouterWithoutRecord.GET("findWlProtocols", wlProtocolsApi.FindWlProtocols)
		wlProtocolsRouterWithoutRecord.GET("getWlProtocolsList", wlProtocolsApi.GetWlProtocolsList)
	}
}
