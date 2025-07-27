package wl_playform

import api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"

type RouterGroup struct {
	WlProductsRouter
	WlScenesRouter
	WlEngineRulesRouter
	WlResourcesRouter
	WlEquipmentRouter
	WlCategoryRouter
	WlCaFunctionRouter
	WlAlarmRouter
<<<<<<< HEAD
	WlDepartmentRouter
=======
	WlUserRouter
>>>>>>> fe0580938f3ae84e1be270b92a56b14cc5c0357a
}

var (
	wlProductsApi    = api.ApiGroupApp.Wl_playformApiGroup.WlProductsApi
	wlScenesApi      = api.ApiGroupApp.Wl_playformApiGroup.WlScenesApi
	wlEngineRulesApi = api.ApiGroupApp.Wl_playformApiGroup.WlEngineRulesApi
	wlResourcesApi   = api.ApiGroupApp.Wl_playformApiGroup.WlResourcesApi
	wlEquipmentApi   = api.ApiGroupApp.Wl_playformApiGroup.WlEquipmentApi
	wlCategoryApi    = api.ApiGroupApp.Wl_playformApiGroup.WlCategoryApi
	wlCaFunctionApi  = api.ApiGroupApp.Wl_playformApiGroup.WlCaFunctionApi
	wlAlarmApi       = api.ApiGroupApp.Wl_playformApiGroup.WlAlarmApi
<<<<<<< HEAD
	wlDepartmentApi  = api.ApiGroupApp.Wl_playformApiGroup.WlDepartmentApi
=======
	wlUserApi        = api.ApiGroupApp.Wl_playformApiGroup.WlUserApi
>>>>>>> fe0580938f3ae84e1be270b92a56b14cc5c0357a
)
