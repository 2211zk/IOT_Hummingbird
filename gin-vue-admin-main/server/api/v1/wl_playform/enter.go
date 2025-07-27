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
<<<<<<< HEAD
	WlDepartmentApi
=======
	WlUserApi
>>>>>>> fe0580938f3ae84e1be270b92a56b14cc5c0357a
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
<<<<<<< HEAD
	wlDepartmentService  = service.ServiceGroupApp.Wl_playformServiceGroup.WlDepartmentService
=======
	wlUserService        = service.ServiceGroupApp.Wl_playformServiceGroup.WlUserService
>>>>>>> fe0580938f3ae84e1be270b92a56b14cc5c0357a
)
