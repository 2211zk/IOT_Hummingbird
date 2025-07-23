package wl_driver

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct{ WlDriversApi }

var wlDriversService = service.ServiceGroupApp.Wl_driverServiceGroup.WlDriversService
