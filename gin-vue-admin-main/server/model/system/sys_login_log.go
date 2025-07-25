package system

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// SysLoginLog 登录日志结构体
type SysLoginLog struct {
	global.GVA_MODEL
	AccessNumber    *int      `json:"accessNumber" form:"accessNumber" gorm:"column:access_number;comment:访问编号"`                          // 访问编号
	UserName        string    `json:"userName" form:"userName" gorm:"column:user_name;size:11;comment:用户名称"`                              // 用户名称
	LoginAddress    string    `json:"loginAddress" form:"loginAddress" gorm:"column:login_address;size:100;comment:登录地址"`                 // 登录地址(IP)
	LoginLocation   string    `json:"loginLocation" form:"loginLocation" gorm:"column:login_location;size:100;comment:登录地点"`              // 登录地点
	Browser         string    `json:"browser" form:"browser" gorm:"column:browser;size:100;comment:浏览器"`                                  // 浏览器
	OperatingSystem string    `json:"operatingSystem" form:"operatingSystem" gorm:"column:operating_system;size:100;comment:操作系统"`        // 操作系统
	LoginStatus     string    `json:"loginStatus" form:"loginStatus" gorm:"column:login_status;size:10;comment:登录状态"`                     // 登录状态
	OperationalInfo string    `json:"operationalInfo" form:"operationalInfo" gorm:"column:operational_information;size:100;comment:操作信息"` // 操作信息
	LoginTime       time.Time `json:"loginTime" form:"loginTime" gorm:"column:login_time;comment:登录时间"`                                   // 登录时间
}

// TableName 设置表名
func (SysLoginLog) TableName() string {
	return "wy_login_log"
}
