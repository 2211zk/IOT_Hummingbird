package dashboard

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/gin-gonic/gin"
)

type DashboardApi struct{}

// GetDashboardData 获取首页仪表盘数据
func (d *DashboardApi) GetDashboardData(c *gin.Context) {
	var result struct {
		PlatformData struct {
			ProductCount        int64 `json:"productCount"`
			PublishedProducts   int64 `json:"publishedProducts"`
			UnpublishedProducts int64 `json:"unpublishedProducts"`
			DeviceCount         int64 `json:"deviceCount"`
			OnlineDevices       int64 `json:"onlineDevices"`
			OfflineDevices      int64 `json:"offlineDevices"`
			DriverCount         int64 `json:"driverCount"`
			RunningDrivers      int64 `json:"runningDrivers"`
			StoppedDrivers      int64 `json:"stoppedDrivers"`
			AlarmCount          int64 `json:"alarmCount"`
		} `json:"platformData"`
		SystemStatus struct {
			CPU struct {
				Usage  float64 `json:"usage"`
				Used   int     `json:"used"`
				Total  int     `json:"total"`
				Status string  `json:"status"`
			} `json:"cpu"`
			Memory struct {
				Usage  float64 `json:"usage"`
				Used   float64 `json:"used"`
				Total  float64 `json:"total"`
				Status string  `json:"status"`
			} `json:"memory"`
			Load struct {
				Usage  float64 `json:"usage"`
				Status string  `json:"status"`
			} `json:"load"`
			Disk struct {
				Usage  float64 `json:"usage"`
				Used   float64 `json:"used"`
				Total  float64 `json:"total"`
				Status string  `json:"status"`
			} `json:"disk"`
		} `json:"systemStatus"`
		AlarmData struct {
			Hint      int64 `json:"hint"`
			Minor     int64 `json:"minor"`
			Important int64 `json:"important"`
			Urgent    int64 `json:"urgent"`
		} `json:"alarmData"`
	}

	// 查询产品数据
	var productCount int64
	global.GVA_DB.Table("wl_products").Count(&productCount)
	// 由于产品表没有status字段，暂时使用总数
	publishedProducts := productCount
	unpublishedProducts := int64(0)

	// 查询设备数据
	var deviceCount int64
	global.GVA_DB.Table("wl_equipment").Count(&deviceCount)
	// 由于设备表没有status字段，暂时使用总数
	onlineDevices := deviceCount
	offlineDevices := int64(0)

	// 查询驱动数据
	var driverCount int64
	global.GVA_DB.Table("wl_drivers").Count(&driverCount)
	// 由于驱动表没有status字段，暂时使用总数
	runningDrivers := driverCount
	stoppedDrivers := int64(0)

	// 查询告警数据 - 由于wl_alarm表可能不存在，暂时使用默认值
	alarmCount := int64(0)
	hintCount := int64(0)
	minorCount := int64(0)
	importantCount := int64(0)
	urgentCount := int64(0)

	// 填充平台数据
	result.PlatformData.ProductCount = productCount
	result.PlatformData.PublishedProducts = publishedProducts
	result.PlatformData.UnpublishedProducts = unpublishedProducts
	result.PlatformData.DeviceCount = deviceCount
	result.PlatformData.OnlineDevices = onlineDevices
	result.PlatformData.OfflineDevices = offlineDevices
	result.PlatformData.DriverCount = driverCount
	result.PlatformData.RunningDrivers = runningDrivers
	result.PlatformData.StoppedDrivers = stoppedDrivers
	result.PlatformData.AlarmCount = alarmCount

	// 模拟系统状态数据（实际项目中应该从系统监控获取）
	result.SystemStatus.CPU.Usage = 22.65
	result.SystemStatus.CPU.Used = 0
	result.SystemStatus.CPU.Total = 2
	result.SystemStatus.CPU.Status = "运行流畅"

	result.SystemStatus.Memory.Usage = 64.98
	result.SystemStatus.Memory.Used = 1.27
	result.SystemStatus.Memory.Total = 1.95
	result.SystemStatus.Memory.Status = "内存使用率较高"

	result.SystemStatus.Load.Usage = 76.36
	result.SystemStatus.Load.Status = "负载较高"

	result.SystemStatus.Disk.Usage = 92.22
	result.SystemStatus.Disk.Used = 45.27
	result.SystemStatus.Disk.Total = 49.09
	result.SystemStatus.Disk.Status = "磁盘空间充足"

	// 填充告警数据
	result.AlarmData.Hint = hintCount
	result.AlarmData.Minor = minorCount
	result.AlarmData.Important = importantCount
	result.AlarmData.Urgent = urgentCount

	response.OkWithData(result, c)
}
