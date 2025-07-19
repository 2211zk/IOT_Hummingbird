package wl_playform

import api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"

type RouterGroup struct{ WlProductsRouter }

var wlProductsApi = api.ApiGroupApp.Wl_playformApiGroup.WlProductsApi
