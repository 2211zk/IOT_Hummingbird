package wl_playform

import api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"

type RouterGroup struct {
	WlProductsRouter
	WlScenesRouter
	WlEngineRulesRouter
	WlResourcesRouter
}

var (
	wlProductsApi    = api.ApiGroupApp.Wl_playformApiGroup.WlProductsApi
	wlScenesApi      = api.ApiGroupApp.Wl_playformApiGroup.WlScenesApi
	wlEngineRulesApi = api.ApiGroupApp.Wl_playformApiGroup.WlEngineRulesApi
	wlResourcesApi   = api.ApiGroupApp.Wl_playformApiGroup.WlResourcesApi
)
