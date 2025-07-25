package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/wl_driver"
	"github.com/flipped-aurora/gin-vue-admin/server/model/wl_playform"
)

func bizModel() error {
	db := global.GVA_DB
<<<<<<< HEAD
=======
<<<<<<< HEAD
=======
<<<<<<< HEAD
	err := db.AutoMigrate(wl_playform.WlProducts{}, wl_playform.WlScenes{}, wl_playform.WlEngineRules{}, wl_playform.WlResources{}, wl_playform.WlEquipment{}, wl_playform.WlCategory{}, wl_playform.WlCaFunction{}, wl_driver.WlDrivers{}, wl_playform.WlAlarm{})
=======
>>>>>>> bd532373303d0055a09cdc9e982af00d97395b5d
>>>>>>> 094b5be16c83cc366e42240e08eecd9895990fc0
	err := db.AutoMigrate(
		wl_playform.WlProducts{},
		wl_playform.WlScenes{},
		wl_playform.WlEngineRules{},
		wl_playform.WlResources{},
		wl_playform.WlEquipment{},
		wl_playform.WlCategory{},
		wl_playform.WlCaFunction{},
<<<<<<< HEAD
		wl_playform.WlAlarm{},
		wl_driver.WlDrivers{},
		wl_driver.DriverCards{},
	)
=======
		wl_driver.DriverCards{},
		wl_driver.WlDrivers{},
		wl_driver.DriverCards{},
	)
<<<<<<< HEAD
=======
>>>>>>> 538e42a4863302292bcdecf856a6196bb9c093b0
>>>>>>> bd532373303d0055a09cdc9e982af00d97395b5d
>>>>>>> 094b5be16c83cc366e42240e08eecd9895990fc0
	if err != nil {
		return err
	}
	return nil
}
