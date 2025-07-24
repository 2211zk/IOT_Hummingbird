package wl_driver

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type DriverCardsRouter struct{}

func (s *DriverCardsRouter) InitDriverCardsRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	driverCardsRouter := Router.Group("driverCards").Use(middleware.OperationRecord())
	driverCardsRouterWithoutRecord := Router.Group("driverCards")
	driverCardsRouterWithoutAuth := PublicRouter.Group("driverCards")
	{
		driverCardsRouter.POST("createDriverCards", driverCardsApi.CreateDriverCards)
		driverCardsRouter.DELETE("deleteDriverCards", driverCardsApi.DeleteDriverCards)
		driverCardsRouter.DELETE("deleteDriverCardsByIds", driverCardsApi.DeleteDriverCardsByIds)
		driverCardsRouter.PUT("updateDriverCards", driverCardsApi.UpdateDriverCards)
	}
	{
		driverCardsRouterWithoutRecord.GET("findDriverCards", driverCardsApi.FindDriverCards)
		driverCardsRouterWithoutRecord.GET("getDriverCardsList", driverCardsApi.GetDriverCardsList)
	}
	{
		driverCardsRouterWithoutAuth.GET("getDriverCardsPublic", driverCardsApi.GetDriverCardsPublic)
	}
}
