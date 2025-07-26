package wl_playform

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	WlProductsApi
	WlScenesApi
	WlEngineRulesApi
	WlResourcesApi
	WlEquipmentApi
	WlCategoryApi
	WlCaFunctionApi
	WlAlarmApi
	WlDepartmentApi
}

var (
	wlProductsService    = service.ServiceGroupApp.Wl_playformServiceGroup.WlProductsService
	wlScenesService      = service.ServiceGroupApp.Wl_playformServiceGroup.WlScenesService
	wlEngineRulesService = service.ServiceGroupApp.Wl_playformServiceGroup.WlEngineRulesService
	wlResourcesService   = service.ServiceGroupApp.Wl_playformServiceGroup.WlResourcesService
	wlEquipmentService   = service.ServiceGroupApp.Wl_playformServiceGroup.WlEquipmentService
	wlCategoryService    = service.ServiceGroupApp.Wl_playformServiceGroup.WlCategoryService
	wlCaFunctionService  = service.ServiceGroupApp.Wl_playformServiceGroup.WlCaFunctionService
	wlAlarmService       = service.ServiceGroupApp.Wl_playformServiceGroup.WlAlarmService
	wlDepartmentService  = service.ServiceGroupApp.Wl_playformServiceGroup.WlDepartmentService
)
