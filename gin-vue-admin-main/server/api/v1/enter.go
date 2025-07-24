package v1

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/example"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/system"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/wl_driver"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/wl_playform"
)

var ApiGroupApp = new(ApiGroup)

type ApiGroup struct {
	SystemApiGroup      system.ApiGroup
	ExampleApiGroup     example.ApiGroup
	Wl_playformApiGroup wl_playform.ApiGroup
	Wl_driverApiGroup   wl_driver.ApiGroup
}
