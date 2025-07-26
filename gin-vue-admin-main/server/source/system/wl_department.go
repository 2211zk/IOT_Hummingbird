package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/wl_department"
	"go.uber.org/zap"
)

type WlDepartmentData struct{}

func (w *WlDepartmentData) InitDB() {
	err := global.GVA_DB.AutoMigrate(&wl_department.WlDepartment{}, &wl_department.WlDevice{}, &wl_department.WlDepartmentDevice{})
	if err != nil {
		global.GVA_LOG.Error("初始化部门表失败!", zap.Error(err))
		return
	}
	global.GVA_LOG.Info("初始化部门表成功!")
}
