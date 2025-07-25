package wl_user

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/wl_user"
	"github.com/gin-gonic/gin"
)

func InitWlUserRouter(Router *gin.RouterGroup) {
	wlUserRouter := Router.Group("sysUser")
	{
		wlUserRouter.POST("getWlUserList", wl_user.WlUserApi.GetWlUserList)
	}
}
