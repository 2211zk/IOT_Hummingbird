package wl_driver

import api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"

type RouterGroup struct {
	DriverCardsRouter
	WlDriversRouter
}

var (
	driverCardsApi = api.ApiGroupApp.Wl_driverApiGroup.DriverCardsApi
	wlDriversApi   = api.ApiGroupApp.Wl_driverApiGroup.WlDriversApi
)
