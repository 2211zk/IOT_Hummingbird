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
		wl_playformRouter.InitWlEquipmentRouter(privateGroup, publicGroup)
		wl_playformRouter.InitWlCategoryRouter(privateGroup, publicGroup)
		wl_playformRouter.InitWlCaFunctionRouter(privateGroup, publicGroup)
	}
	{
		wl_driverRouter := router.RouterGroupApp.Wl_driver
		wl_driverRouter.DriverCardsRouter.InitDriverCardsRouter(privateGroup, publicGroup)
		wl_driverRouter.InitWlDriversRouter(privateGroup, publicGroup)
	}
}
