package wl_driver

import api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"

type RouterGroup struct {
	WlDriversRouter
	WlProtocolsRouter
	WlDriverMarketRouter
	WlSystemMoniorRouter
}

var wlDriversApi = api.ApiGroupApp.Wl_driverApiGroup.WlDriversApi
var wlProtocolsApi = api.ApiGroupApp.Wl_driverApiGroup.WlProtocolsApi
var wlDriverMarketApi = api.ApiGroupApp.Wl_driverApiGroup.WlDriverMarketApi
var wlSystemMoniorApi = api.ApiGroupApp.Wl_driverApiGroup.WlSystemMoniorApi
