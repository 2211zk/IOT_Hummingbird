package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/wl_department"
	"go.uber.org/zap"
)

// InitWlDepartmentTables 初始化部门相关表
func InitWlDepartmentTables() {
	// 自动迁移表结构
	err := global.GVA_DB.AutoMigrate(
		&wl_department.WlDepartment{},
		&wl_department.WlDevice{},
		&wl_department.WlDepartmentDevice{},
	)
	if err != nil {
		global.GVA_LOG.Error("初始化部门表失败!", zap.Error(err))
		return
	}
	global.GVA_LOG.Info("初始化部门表成功!")

	// 检查是否有数据，如果没有则插入测试数据
	var count int64
	global.GVA_DB.Model(&wl_department.WlDepartment{}).Count(&count)
	if count == 0 {
		// 插入测试数据
		testDepartments := []wl_department.WlDepartment{
			{
				Name:           "技术部",
				DepartmentName: "技术部",
				Leader:         "张三",
				Phone:          "13800138001",
				Email:          "tech@company.com",
				Status:         "启用",
				Sort:           1,
			},
			{
				Name:           "市场部",
				DepartmentName: "市场部",
				Leader:         "李四",
				Phone:          "13800138002",
				Email:          "market@company.com",
				Status:         "启用",
				Sort:           2,
			},
			{
				Name:           "人事部",
				DepartmentName: "人事部",
				Leader:         "王五",
				Phone:          "13800138003",
				Email:          "hr@company.com",
				Status:         "启用",
				Sort:           3,
			},
		}

		for _, dept := range testDepartments {
			if err := global.GVA_DB.Create(&dept).Error; err != nil {
				global.GVA_LOG.Error("插入测试部门数据失败!", zap.Error(err))
			}
		}

		// 插入测试设备数据
		testDevices := []wl_department.WlDevice{
			{
				DeviceName:  "服务器001",
				ProductName: "Dell PowerEdge",
				Status:      "启用",
			},
			{
				DeviceName:  "服务器002",
				ProductName: "HP ProLiant",
				Status:      "启用",
			},
			{
				DeviceName:  "网络设备001",
				ProductName: "Cisco Switch",
				Status:      "启用",
			},
		}

		for _, device := range testDevices {
			if err := global.GVA_DB.Create(&device).Error; err != nil {
				global.GVA_LOG.Error("插入测试设备数据失败!", zap.Error(err))
			}
		}

		global.GVA_LOG.Info("插入测试数据成功!")
	}
}
