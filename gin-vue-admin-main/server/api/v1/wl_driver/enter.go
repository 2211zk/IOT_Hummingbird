package wl_driver

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	DriverCardsApi
	WlDriversApi
}

var wlDriversService = service.ServiceGroupApp.Wl_driverServiceGroup.WlDriversService
