package wl_playform

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	WlProductsApi
	WlEquipmentApi
	WlCategoryApi
	WlCaFunctionApi
}

var (
	wlProductsService   = service.ServiceGroupApp.Wl_playformServiceGroup.WlProductsService
	wlEquipmentService  = service.ServiceGroupApp.Wl_playformServiceGroup.WlEquipmentService
	wlCategoryService   = service.ServiceGroupApp.Wl_playformServiceGroup.WlCategoryService
	wlCaFunctionService = service.ServiceGroupApp.Wl_playformServiceGroup.WlCaFunctionService
)
