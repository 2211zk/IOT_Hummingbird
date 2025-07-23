package wl_playform

import api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"

type RouterGroup struct {
	WlProductsRouter
	WlEquipmentRouter
	WlCategoryRouter
	WlCaFunctionRouter
}

var (
	wlProductsApi   = api.ApiGroupApp.Wl_playformApiGroup.WlProductsApi
	wlEquipmentApi  = api.ApiGroupApp.Wl_playformApiGroup.WlEquipmentApi
	wlCategoryApi   = api.ApiGroupApp.Wl_playformApiGroup.WlCategoryApi
	wlCaFunctionApi = api.ApiGroupApp.Wl_playformApiGroup.WlCaFunctionApi
)
