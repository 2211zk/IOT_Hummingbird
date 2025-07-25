package request

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

// LoginLogSearchReq 登录日志搜索请求结构体
type LoginLogSearchReq struct {
	request.PageInfo
	UserName      string    `json:"userName" form:"userName"`           // 用户名称
	LoginAddress  string    `json:"loginAddress" form:"loginAddress"`   // 登录地址
	LoginLocation string    `json:"loginLocation" form:"loginLocation"` // 登录地点
	LoginStatus   string    `json:"loginStatus" form:"loginStatus"`     // 登录状态
	StartTime     time.Time `json:"startTime" form:"startTime"`         // 开始时间
	EndTime       time.Time `json:"endTime" form:"endTime"`             // 结束时间
}

// CreateLoginLogReq 创建登录日志请求结构体
type CreateLoginLogReq struct {
	AccessNumber    *int   `json:"accessNumber" form:"accessNumber"`       // 访问编号
	UserName        string `json:"userName" form:"userName"`               // 用户名称
	LoginAddress    string `json:"loginAddress" form:"loginAddress"`       // 登录地址
	LoginLocation   string `json:"loginLocation" form:"loginLocation"`     // 登录地点
	Browser         string `json:"browser" form:"browser"`                 // 浏览器
	OperatingSystem string `json:"operatingSystem" form:"operatingSystem"` // 操作系统
	LoginStatus     string `json:"loginStatus" form:"loginStatus"`         // 登录状态
	OperationalInfo string `json:"operationalInfo" form:"operationalInfo"` // 操作信息
}

// DeleteLoginLogReq 删除登录日志请求结构体
type DeleteLoginLogReq struct {
	IDs []uint `json:"ids" form:"ids"` // 要删除的ID列表
}

// CleanExpiredLogsReq 清理过期日志请求结构体
type CleanExpiredLogsReq struct {
	Days int `json:"days" form:"days"` // 保留天数
}
