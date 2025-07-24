package service

import (
	"github.com/flipped-aurora/gin-vue-admin/server/service/example"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
	"github.com/flipped-aurora/gin-vue-admin/server/service/wl_driver"
	"github.com/flipped-aurora/gin-vue-admin/server/service/wl_playform"
)

var ServiceGroupApp = new(ServiceGroup)

type ServiceGroup struct {
	SystemServiceGroup      system.ServiceGroup
	ExampleServiceGroup     example.ServiceGroup
	Wl_playformServiceGroup wl_playform.ServiceGroup
	Wl_driverServiceGroup   wl_driver.ServiceGroup
}
