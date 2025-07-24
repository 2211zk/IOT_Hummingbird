package wl_driver

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	WlDriversApi
	WlProtocolsApi
	WlDriverMarketApi
	WlSystemMoniorApi
}

var wlDriversService = service.ServiceGroupApp.Wl_driverServiceGroup.WlDriversService
var wlProtocolsService = service.ServiceGroupApp.Wl_driverServiceGroup.WlProtocolsService
var wlProtocolsApi = new(WlProtocolsApi)
var wlSystemMoniorService = service.ServiceGroupApp.Wl_driverServiceGroup.WlSystemMoniorService
var wlSystemMoniorApi = new(WlSystemMoniorApi)
var wlDriverMarketService = service.ServiceGroupApp.Wl_driverServiceGroup.WlDriverMarketService
var wlDriverMarketApi = new(WlDriverMarketApi)
