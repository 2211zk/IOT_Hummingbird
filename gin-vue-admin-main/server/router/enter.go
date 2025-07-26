package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/router/dashboard"
	"github.com/flipped-aurora/gin-vue-admin/server/router/example"
	"github.com/flipped-aurora/gin-vue-admin/server/router/system"
	"github.com/flipped-aurora/gin-vue-admin/server/router/wl_driver"
	"github.com/flipped-aurora/gin-vue-admin/server/router/wl_playform"
)

var RouterGroupApp = new(RouterGroup)

type RouterGroup struct {
	System      system.RouterGroup
	Example     example.RouterGroup
	Wl_playform wl_playform.RouterGroup
	Dashboard   dashboard.RouterGroup
	Wl_driver   wl_driver.RouterGroup
}

// Wl_user     wl_user.RouterGroup // 已删除，避免未定义报错
