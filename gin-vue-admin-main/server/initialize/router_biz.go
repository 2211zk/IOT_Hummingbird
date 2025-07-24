package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/router"
	"github.com/gin-gonic/gin"
)

func holder(routers ...*gin.RouterGroup) {
	_ = routers
	_ = router.RouterGroupApp
}
func initBizRouter(routers ...*gin.RouterGroup) {
	privateGroup := routers[0]
	publicGroup := routers[1]
	holder(publicGroup, privateGroup)
	{
		wl_playformRouter := router.RouterGroupApp.Wl_playform
		wl_playformRouter.InitWlProductsRouter(privateGroup, publicGroup)
		wl_playformRouter.InitWlScenesRouter(privateGroup, publicGroup)
		wl_playformRouter.InitWlEngineRulesRouter(privateGroup, publicGroup) // 占位方法，保证文件可以正确加载，避免go空变量检测报错，请勿删除。
		wl_playformRouter.InitWlResourcesRouter(privateGroup, publicGroup)
		wl_playformRouter.InitWlEquipmentRouter(privateGroup, publicGroup)
		wl_playformRouter.InitWlCategoryRouter(privateGroup, publicGroup)
		wl_playformRouter.InitWlCaFunctionRouter(privateGroup, publicGroup)
	}
	{
		dashboardRouter := router.RouterGroupApp.Dashboard
		dashboardRouter.InitDashboardRouter(privateGroup)
	}
	{
		wl_driverRouter := router.RouterGroupApp.Wl_driver
		wl_driverRouter.WlDriversRouter.InitWlDriversRouter(privateGroup, publicGroup)
		wl_driverRouter.WlProtocolsRouter.InitWlProtocolsRouter(privateGroup, publicGroup)
		wl_driverRouter.WlDriverMarketRouter.InitWlDriverMarketRouter(privateGroup, publicGroup)
		wl_driverRouter.WlSystemMoniorRouter.InitWlSystemMoniorRouter(privateGroup, publicGroup)
	}
}
