package wl_driver

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type WlSystemMoniorRouter struct{}

func (s *WlSystemMoniorRouter) InitWlSystemMoniorRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	wlSystemMoniorRouter := Router.Group("wlSystemMonior").Use(middleware.OperationRecord())
	wlSystemMoniorRouterWithoutRecord := Router.Group("wlSystemMonior")
	{
		wlSystemMoniorRouter.POST("createWlSystemMonior", wlSystemMoniorApi.CreateWlSystemMonior)
		wlSystemMoniorRouter.DELETE("deleteWlSystemMonior", wlSystemMoniorApi.DeleteWlSystemMonior)
		wlSystemMoniorRouter.DELETE("deleteWlSystemMoniorByIds", wlSystemMoniorApi.DeleteWlSystemMoniorByIds)
		wlSystemMoniorRouter.PUT("updateWlSystemMonior", wlSystemMoniorApi.UpdateWlSystemMonior)
	}
	{
		wlSystemMoniorRouterWithoutRecord.GET("findWlSystemMonior", wlSystemMoniorApi.FindWlSystemMonior)
		wlSystemMoniorRouterWithoutRecord.GET("getWlSystemMoniorList", wlSystemMoniorApi.GetWlSystemMoniorList)
	}
}
