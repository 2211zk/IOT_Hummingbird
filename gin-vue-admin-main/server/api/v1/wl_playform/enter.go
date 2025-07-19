package wl_playform

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct{ WlProductsApi }

var wlProductsService = service.ServiceGroupApp.Wl_playformServiceGroup.WlProductsService
