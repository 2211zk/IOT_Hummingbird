package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/router/dashboard"
	"github.com/flipped-aurora/gin-vue-admin/server/router/example"
	"github.com/flipped-aurora/gin-vue-admin/server/router/system"
	"github.com/flipped-aurora/gin-vue-admin/server/router/wl_department"
	"github.com/flipped-aurora/gin-vue-admin/server/router/wl_driver"
	"github.com/flipped-aurora/gin-vue-admin/server/router/wl_playform"
)

var RouterGroupApp = new(RouterGroup)

type RouterGroup struct {
<<<<<<< HEAD
	System        system.RouterGroup
	Example       example.RouterGroup
	Wl_playform   wl_playform.RouterGroup
	Wl_department wl_department.RouterGroup
	Dashboard     dashboard.RouterGroup
	Wl_driver     wl_driver.RouterGroup
=======
	System      system.RouterGroup
	Example     example.RouterGroup
	Wl_playform wl_playform.RouterGroup
	Dashboard   dashboard.RouterGroup
	Wl_driver   wl_driver.RouterGroup
>>>>>>> fe0580938f3ae84e1be270b92a56b14cc5c0357a
}

// Wl_user     wl_user.RouterGroup // 已删除，避免未定义报错
