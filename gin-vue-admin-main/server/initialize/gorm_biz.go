package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/wl_driver"
	"github.com/flipped-aurora/gin-vue-admin/server/model/wl_playform"
)

func bizModel() error {
	db := global.GVA_DB

	err := db.AutoMigrate(
		wl_playform.WlProducts{},
		wl_playform.WlScenes{},
		wl_playform.WlEngineRules{},
		wl_playform.WlResources{},
		wl_playform.WlEquipment{},
		wl_playform.WlCategory{},
		wl_playform.WlCaFunction{},
		wl_driver.WlDrivers{},
	)

	err = db.AutoMigrate(wl_playform.WlProducts{}, wl_playform.WlEquipment{}, wl_playform.WlCategory{}, wl_playform.WlCaFunction{}, wl_driver.DriverCards{}, wl_driver.WlDrivers{})

	if err != nil {
		return err
	}
	return nil
}
