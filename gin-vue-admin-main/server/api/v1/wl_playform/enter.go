package wl_playform

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	WlProductsApi
	WlScenesApi
	WlEngineRulesApi
	WlResourcesApi
}

var (
	wlProductsService    = service.ServiceGroupApp.Wl_playformServiceGroup.WlProductsService
	wlScenesService      = service.ServiceGroupApp.Wl_playformServiceGroup.WlScenesService
	wlEngineRulesService = service.ServiceGroupApp.Wl_playformServiceGroup.WlEngineRulesService
	wlResourcesService   = service.ServiceGroupApp.Wl_playformServiceGroup.WlResourcesService
)
