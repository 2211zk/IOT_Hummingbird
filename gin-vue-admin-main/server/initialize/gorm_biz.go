package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/wl_driver"
	"github.com/flipped-aurora/gin-vue-admin/server/model/wl_playform"
)

func bizModel() error {
	db := global.GVA_DB
<<<<<<< HEAD
	err := db.AutoMigrate(wl_playform.WlProducts{}, wl_playform.WlScenes{}, wl_playform.WlEngineRules{}, wl_playform.WlResources{}, wl_playform.WlEquipment{}, wl_playform.WlCategory{}, wl_playform.WlCaFunction{}, wl_driver.WlDrivers{}, wl_playform.WlAlarm{}, wl_playform.WlDepartment{}, wl_playform.WlDepartment{})
	err = db.AutoMigrate(wl_playform.WlDepartment{}, wl_playform.WlDepartment{})
	err = db.AutoMigrate(wl_playform.WlProducts{}, wl_playform.WlScenes{}, wl_playform.WlEngineRules{}, wl_playform.WlResources{}, wl_playform.WlEquipment{}, wl_playform.WlCategory{}, wl_playform.WlCaFunction{}, wl_driver.WlDrivers{}, wl_playform.WlAlarm{}, wl_playform.WlDepartment{}, wl_playform.WlDepartment{})
	err = db.AutoMigrate(wl_playform.WlProducts{}, wl_playform.WlScenes{}, wl_playform.WlEngineRules{}, wl_playform.WlResources{}, wl_playform.WlEquipment{}, wl_playform.WlCategory{}, wl_playform.WlCaFunction{}, wl_playform.WlAlarm{}, wl_driver.WlDrivers{}, wl_driver.DriverCards{}, wl_driver.DriverCards{}, wl_driver.WlDrivers{}, wl_driver.DriverCards{}, wl_playform.WlDepartment{}, wl_playform.WlDepartment{})
=======
	err := db.AutoMigrate(
		wl_playform.WlProducts{},
		wl_playform.WlScenes{},
		wl_playform.WlEngineRules{},
		wl_playform.WlResources{},
		wl_playform.WlEquipment{},
		wl_playform.WlCategory{},
		wl_playform.WlCaFunction{},
		wl_playform.WlAlarm{},
		wl_driver.WlDrivers{},
		wl_driver.DriverCards{},
	)
	err = db.AutoMigrate(wl_playform.WlProducts{}, wl_playform.WlScenes{}, wl_playform.WlEngineRules{}, wl_playform.WlResources{}, wl_playform.WlEquipment{}, wl_playform.WlCategory{}, wl_playform.WlCaFunction{}, wl_driver.WlDrivers{}, wl_playform.WlAlarm{}, wl_playform.WlUser{})
	err = db.AutoMigrate(wl_playform.WlUser{})
	err = db.AutoMigrate(wl_playform.WlProducts{}, wl_playform.WlScenes{}, wl_playform.WlEngineRules{}, wl_playform.WlResources{}, wl_playform.WlEquipment{}, wl_playform.WlCategory{}, wl_playform.WlCaFunction{}, wl_driver.WlDrivers{}, wl_playform.WlAlarm{}, wl_playform.WlUser{})
	err = db.AutoMigrate(wl_playform.WlProducts{}, wl_playform.WlScenes{}, wl_playform.WlEngineRules{}, wl_playform.WlResources{}, wl_playform.WlEquipment{}, wl_playform.WlCategory{}, wl_playform.WlCaFunction{}, wl_playform.WlAlarm{}, wl_driver.WlDrivers{}, wl_driver.DriverCards{}, wl_driver.DriverCards{}, wl_driver.WlDrivers{}, wl_driver.DriverCards{}, wl_playform.WlUser{})
>>>>>>> fe0580938f3ae84e1be270b92a56b14cc5c0357a
	if err != nil {
		return err
	}
	return nil
}
