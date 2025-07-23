package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"go.uber.org/zap"
)

// WanlianResourceInit 初始化Wanlian_resource数据源
func WanlianResourceInit() {
	if err := WanlianMongo.Initialization(); err != nil {
		global.GVA_LOG.Error("Wanlian_resource MongoDB初始化失败!", zap.Error(err))
		panic("Wanlian_resource MongoDB初始化失败!")
	}
	global.GVA_LOG.Info("Wanlian_resource MongoDB初始化成功!")
}
