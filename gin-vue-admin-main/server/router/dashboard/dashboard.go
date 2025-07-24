package dashboard

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
)

type DashboardRouter struct{}

// InitDashboardRouter 初始化仪表盘路由
func (s *DashboardRouter) InitDashboardRouter(Router *gin.RouterGroup) {
	dashboardRouterWithoutRecord := Router.Group("dashboard")
	dashboardApi := v1.ApiGroupApp.DashboardApiGroup.DashboardApi
	{
		dashboardRouterWithoutRecord.GET("getDashboardData", dashboardApi.GetDashboardData) // 获取仪表盘数据
	}
}
