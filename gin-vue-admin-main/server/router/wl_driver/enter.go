package wl_driver

import api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"

type RouterGroup struct{ WlDriversRouter }

var wlDriversApi = api.ApiGroupApp.Wl_driverApiGroup.WlDriversApi
