package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/wl_playform"
)

func bizModel() error {
	db := global.GVA_DB
	err := db.AutoMigrate(wl_playform.WlProducts{})
	if err != nil {
		return err
	}
	return nil
}
