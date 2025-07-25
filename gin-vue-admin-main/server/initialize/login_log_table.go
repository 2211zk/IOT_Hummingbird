package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"go.uber.org/zap"
)

// InitLoginLogTable 初始化登录日志表
func InitLoginLogTable() {
	db := global.GVA_DB

	// 检查表是否存在
	if !db.Migrator().HasTable(&system.SysLoginLog{}) {
		global.GVA_LOG.Warn("wy_login_log table does not exist, please create it manually")
		return
	}

	// 验证表结构
	if err := db.AutoMigrate(&system.SysLoginLog{}); err != nil {
		global.GVA_LOG.Error("login log table structure validation failed", zap.Error(err))
	} else {
		global.GVA_LOG.Info("login log table structure validated successfully")
	}
}
